//go:build integration

package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// Pipe content strings copied verbatim from docs/exercise.md.

const pipeCleanPageViews = `type: TABLE
name: clean_page_views
description: Drops incomplete page view rows and normalizes timestamp to date.
destination: clean_page_views

sources:
  - raw: page_views

pipeline:
  @result:
    SELECT
      toDate(timestamp) AS day,
      user_id,
      page
    FROM raw
    WHERE user_id != ''
      AND page != ''`

const pipeCleanProfiles = `type: TABLE
name: clean_profiles
description: Drops user profiles with missing plan tier.
destination: clean_profiles

sources:
  - raw: user_profiles

pipeline:
  @result:
    SELECT
      user_id,
      plan,
      country,
      signup_date
    FROM raw
    WHERE user_id != ''
      AND plan != ''`

const pipeEnrichSessions = `type: TABLE
name: enrich_sessions
description: Joins clean page views with user profiles to produce enriched sessions.
destination: enriched_sessions

sources:
  - views:    clean_page_views
  - profiles: clean_profiles

pipeline:
  @result:
    SELECT
      views.day,
      views.user_id,
      views.page,
      profiles.plan,
      profiles.country
    FROM views
    INNER JOIN profiles ON views.user_id = profiles.user_id`

const pipeDailyStats = `type: ENDPOINT
name: daily_stats
description: Daily active users segmented by plan and country, with optional filters.

sources:
  - enriched: enriched_sessions

params:
  plan:    { type: string, default: "" }
  country: { type: string, default: "" }

pipeline:
  @filter:
    SELECT day, user_id, plan, country
    FROM enriched
    WHERE ({{ params.plan }} = '' OR plan = {{ params.plan }})
      AND ({{ params.country }} = '' OR country = {{ params.country }})

  @result:
    SELECT
      day,
      plan,
      country,
      uniqExact(user_id) AS dau
    FROM @filter
    GROUP BY day, plan, country
    ORDER BY day DESC, dau DESC`

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
  @filter:
    SELECT day, user_id, plan, country
    FROM enriched
    WHERE ({{ params.plan }} = '' OR plan = {{ params.plan }})
      AND ({{ params.country }} = '' OR country = {{ params.country }})
      AND day >= toDate({{ params.start_day }})

  @result:
    SELECT
      day,
      plan,
      country,
      uniqExact(user_id) AS dau
    FROM @filter
    GROUP BY day, plan, country
    ORDER BY day DESC, dau DESC`

type postablePipe struct {
	Content string `json:"content"`
}

// TestCrossPipeExercise runs the exercise from docs/exercise.md as an integration test.
// Steps 8 and 9 are UI-only and are omitted.
// Steps 6 and 7 require an ENDPOINT execution route that does not exist yet (see TODO below).
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
		// TODO(route): no ENDPOINT execution route exists yet.
		// exercise.md shows GET /v0/pipes/daily_stats?plan=pro but the server only registers:
		//   GET /api/v0/pipes         (list)
		//   GET /api/v0/pipes/{name}/meta  (metadata)
		// Once the execute route is added, implement:
		//   data := readResponse(t, apiDo(t, "GET", "/api/v0/pipes/daily_stats/execute", nil), http.StatusOK)
		//   assertRowsNotEmpty(t, data)
		//   data = readResponse(t, apiDo(t, "GET", "/api/v0/pipes/daily_stats/execute?plan=pro", nil), http.StatusOK)
		//   assertAllPlan(t, data, "pro")
		//   data = readResponse(t, apiDo(t, "GET", "/api/v0/pipes/daily_stats/execute?country=JP", nil), http.StatusOK)
		//   assertRowCount(t, data, 0)
		t.Skip("TODO(route): ENDPOINT execution route not yet implemented")
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

		// TODO(route): once execute route exists, verify:
		//   GET /execute?start_day=2024-01-16 returns only 2024-01-16 rows
		//   GET /execute?plan=pro&start_day=2024-01-16 returns pro rows from 2024-01-16 only
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
