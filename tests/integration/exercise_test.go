//go:build integration

package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// Pipe content uses PRQL (Pipeline Query Language) pipeline syntax.

const pipeCleanPageViews = `type: TABLE
name: clean_page_views
description: Drops incomplete page view rows and normalizes timestamp to date.
destination: clean_page_views

sources:
  - raw: page_views

pipeline:
  @result:
    from raw
    filter user_id != '' AND page != ''
    select { day = toDate(timestamp), user_id, page }`

const pipeCleanProfiles = `type: TABLE
name: clean_profiles
description: Drops user profiles with missing plan tier.
destination: clean_profiles

sources:
  - raw: user_profiles

pipeline:
  @result:
    from raw
    filter user_id != '' AND plan != ''
    select { user_id, plan, country, signup_date }`

const pipeEnrichSessions = `type: TABLE
name: enrich_sessions
description: Joins clean page views with user profiles to produce enriched sessions.
destination: enriched_sessions

sources:
  - views:    clean_page_views
  - profiles: clean_profiles

pipeline:
  @result:
    from views
    join side:inner profiles (views.user_id = profiles.user_id)
    select { views.day, views.user_id, views.page, profiles.plan, profiles.country }`

const pipeDailyStats = `type: ENDPOINT
name: daily_stats
description: Daily active users segmented by plan and country, with optional filters.

sources:
  - enriched: enriched_sessions

params:
  plan:    { type: string, default: "" }
  country: { type: string, default: "" }

pipeline:
  @filtered:
    from enriched
    filter ({{ params.plan }} = '' OR plan = {{ params.plan }}) AND ({{ params.country }} = '' OR country = {{ params.country }})
    select { day, user_id, plan, country }

  @result:
    from filtered
    group { day, plan, country } ( aggregate { dau = uniqExact(user_id) } )
    sort { -day, -dau }`

const pipeDailyStatsWithStartDay = `type: ENDPOINT
name: daily_stats
description: Daily active users segmented by plan and country, with optional filters.

sources:
  - enriched: enriched_sessions

params:
  plan:      { type: string, default: "" }
  country:   { type: string, default: "" }
  start_day: { type: string, default: "2000-01-01" }

pipeline:
  @filtered:
    from enriched
    filter ({{ params.plan }} = '' OR plan = {{ params.plan }}) AND ({{ params.country }} = '' OR country = {{ params.country }}) AND day >= toDate({{ params.start_day }})
    select { day, user_id, plan, country }

  @result:
    from filtered
    group { day, plan, country } ( aggregate { dau = uniqExact(user_id) } )
    sort { -day, -dau }`

type postablePipe struct {
	Content string `json:"content"`
}

// TestCrossPipeExercise runs the exercise from docs/exercise.md as an integration test.
// Steps 8 and 9 are UI-only and are omitted.
func TestCrossPipeExercise(t *testing.T) {

	t.Run("Step1_RegisterSources", func(t *testing.T) {
		pageViews := map[string]any{
			"name":   "page_views",
			"engine": "MergeTree()",
			"fields": []map[string]string{
				{"name": "timestamp", "type": "datetime64"},
				{"name": "user_id", "type": "string"},
				{"name": "page", "type": "string"},
			},
		}
		data := readResponse(t, apiDo(t, "POST", "/api/v0/sources", pageViews), http.StatusCreated)
		var src map[string]any
		mustUnmarshal(t, data, &src)
		if src["name"] != "page_views" {
			t.Errorf("want name page_views, got %v", src["name"])
		}

		userProfiles := map[string]any{
			"name":   "user_profiles",
			"engine": "MergeTree()",
			"fields": []map[string]string{
				{"name": "user_id", "type": "string"},
				{"name": "plan", "type": "string"},
				{"name": "country", "type": "string"},
				{"name": "signup_date", "type": "date"},
			},
		}
		data = readResponse(t, apiDo(t, "POST", "/api/v0/sources", userProfiles), http.StatusCreated)
		mustUnmarshal(t, data, &src)
		if src["name"] != "user_profiles" {
			t.Errorf("want name user_profiles, got %v", src["name"])
		}

		// Verify both are listed.
		data = readResponse(t, apiDo(t, "GET", "/api/v0/sources", nil), http.StatusOK)
		var sources []map[string]any
		mustUnmarshal(t, data, &sources)
		if len(sources) < 2 {
			t.Errorf("want >= 2 sources, got %d", len(sources))
		}
	})

	t.Run("Step1b_SeedData", func(t *testing.T) {
		seedTestData(t)
	})

	t.Run("Step2_CleanPageViews", func(t *testing.T) {
		readResponse(t, apiDo(t, "POST", "/api/v0/pipes", postablePipe{Content: pipeCleanPageViews}), http.StatusCreated)

		data := readResponse(t, apiDo(t, "GET", "/api/v0/pipes/clean_page_views/meta", nil), http.StatusOK)
		assertPipeType(t, data, "TABLE")
	})

	t.Run("Step3_CleanProfiles", func(t *testing.T) {
		readResponse(t, apiDo(t, "POST", "/api/v0/pipes", postablePipe{Content: pipeCleanProfiles}), http.StatusCreated)

		data := readResponse(t, apiDo(t, "GET", "/api/v0/pipes/clean_profiles/meta", nil), http.StatusOK)
		assertPipeType(t, data, "TABLE")
	})

	t.Run("Step4_EnrichSessions", func(t *testing.T) {
		readResponse(t, apiDo(t, "POST", "/api/v0/pipes", postablePipe{Content: pipeEnrichSessions}), http.StatusCreated)

		data := readResponse(t, apiDo(t, "GET", "/api/v0/pipes/enrich_sessions/meta", nil), http.StatusOK)
		assertPipeType(t, data, "TABLE")
		// Verify two sources are declared.
		var pipe map[string]any
		mustUnmarshal(t, data, &pipe)
		sources, _ := pipe["sources"].([]any)
		if len(sources) != 2 {
			t.Errorf("enrich_sessions: want 2 sources, got %d", len(sources))
		}
	})

	t.Run("Step5_DailyStats", func(t *testing.T) {
		readResponse(t, apiDo(t, "POST", "/api/v0/pipes", postablePipe{Content: pipeDailyStats}), http.StatusCreated)

		data := readResponse(t, apiDo(t, "GET", "/api/v0/pipes/daily_stats/meta", nil), http.StatusOK)
		assertPipeType(t, data, "ENDPOINT")

		// All four pipes are listed.
		data = readResponse(t, apiDo(t, "GET", "/api/v0/pipes", nil), http.StatusOK)
		var pipes []map[string]any
		mustUnmarshal(t, data, &pipes)
		if len(pipes) != 4 {
			t.Errorf("want 4 pipes, got %d", len(pipes))
		}
	})

	t.Run("Step6_ExecuteEndpoint", func(t *testing.T) {
		data := readResponse(t, apiDo(t, "GET", "/api/v0/execute/daily_stats", nil), http.StatusOK)
		assertRowsNotEmpty(t, data)

		data = readResponse(t, apiDo(t, "GET", "/api/v0/execute/daily_stats?plan=pro", nil), http.StatusOK)
		assertAllPlan(t, data, "pro")

		// JP has no clean rows (u4 has empty plan, filtered by clean_profiles).
		data = readResponse(t, apiDo(t, "GET", "/api/v0/execute/daily_stats?country=JP", nil), http.StatusOK)
		assertRowCount(t, data, 0)
	})

	t.Run("Step7_UpdateEndpoint", func(t *testing.T) {
		// Depends on Step 6 (execution) to verify the start_day param filters correctly.
		// TODO(route): unblock once the execute route exists (see Step6 TODO).
		readResponse(t, apiDo(t, "PUT", "/api/v0/pipes/daily_stats", postablePipe{Content: pipeDailyStatsWithStartDay}), http.StatusOK)

		// Verify the updated pipe has the new start_day param.
		data := readResponse(t, apiDo(t, "GET", "/api/v0/pipes/daily_stats/meta", nil), http.StatusOK)
		var pipe map[string]any
		mustUnmarshal(t, data, &pipe)
		params, _ := pipe["params"].(map[string]any)
		if _, ok := params["start_day"]; !ok {
			t.Errorf("updated pipe missing start_day param; params: %v", params)
		}

		// Verify start_day filters correctly.
		data = readResponse(t, apiDo(t, "GET", "/api/v0/execute/daily_stats?start_day=2024-01-16", nil), http.StatusOK)
		assertAllDay(t, data, "2024-01-16")

		data = readResponse(t, apiDo(t, "GET", "/api/v0/execute/daily_stats?plan=pro&start_day=2024-01-16", nil), http.StatusOK)
		assertAllPlan(t, data, "pro")
		assertAllDay(t, data, "2024-01-16")
	})

	t.Run("Step10_DeletePipes", func(t *testing.T) {
		// Delete downstream-first to avoid dangling dependencies.
		for _, name := range []string{"daily_stats", "enrich_sessions", "clean_profiles", "clean_page_views"} {
			assertNoContent(t, apiDo(t, "DELETE", fmt.Sprintf("/api/v0/pipes/%s", name), nil))
		}

		data := readResponse(t, apiDo(t, "GET", "/api/v0/pipes", nil), http.StatusOK)
		var pipes []map[string]any
		mustUnmarshal(t, data, &pipes)
		if len(pipes) != 0 {
			t.Errorf("want 0 pipes after deletion, got %d", len(pipes))
		}
	})
}

func assertPipeType(t *testing.T, data json.RawMessage, want string) {
	t.Helper()
	var pipe map[string]any
	mustUnmarshal(t, data, &pipe)
	if pipe["type"] != want {
		t.Errorf("want pipe type %q, got %v", want, pipe["type"])
	}
}

func mustUnmarshal(t *testing.T, data json.RawMessage, v any) {
	t.Helper()
	if err := json.Unmarshal(data, v); err != nil {
		t.Fatalf("unmarshal: %v (raw: %s)", err, data)
	}
}

type executeResult struct {
	Data []map[string]any `json:"data"`
}

func unmarshalExecute(t *testing.T, data json.RawMessage) executeResult {
	t.Helper()
	var r executeResult
	mustUnmarshal(t, data, &r)
	return r
}

func assertRowsNotEmpty(t *testing.T, data json.RawMessage) {
	t.Helper()
	if r := unmarshalExecute(t, data); len(r.Data) == 0 {
		t.Error("expected non-empty execute result")
	}
}

func assertRowCount(t *testing.T, data json.RawMessage, n int) {
	t.Helper()
	if r := unmarshalExecute(t, data); len(r.Data) != n {
		t.Errorf("execute result: want %d rows, got %d", n, len(r.Data))
	}
}

func assertAllPlan(t *testing.T, data json.RawMessage, plan string) {
	t.Helper()
	r := unmarshalExecute(t, data)
	for _, row := range r.Data {
		if row["plan"] != plan {
			t.Errorf("execute result: want plan %q, got %v (row: %v)", plan, row["plan"], row)
		}
	}
}

func assertAllDay(t *testing.T, data json.RawMessage, day string) {
	t.Helper()
	r := unmarshalExecute(t, data)
	for _, row := range r.Data {
		got := fmt.Sprintf("%v", row["day"])
		// ClickHouse Date columns marshal as RFC 3339 ("2024-01-16T00:00:00Z"); take YYYY-MM-DD prefix.
		if len(got) > 10 {
			got = got[:10]
		}
		if got != day {
			t.Errorf("execute result: want day %q, got %v (row: %v)", day, row["day"], row)
		}
	}
}
