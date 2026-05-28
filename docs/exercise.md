# Exercise: Build a Cross-Pipe Analytics Pipeline

This exercise builds a four-pipe analytics pipeline from scratch. Two raw sources feed two independent cleaning pipes, whose outputs are joined into an enriched table, which is then exposed as a parameterized REST endpoint. Each pipe depends on the one before it — nothing is standalone.

**Scenario**: Page view events joined with user profile data → daily active user counts segmented by plan tier and country.

---

## Prerequisite

The server is running and connected to ClickHouse. The following raw tables exist with data:

**`page_views`**

| Column      | Type         |
|-------------|--------------|
| `timestamp` | `DateTime64` |
| `user_id`   | `String`     |
| `page`      | `String`     |

**`user_profiles`**

| Column        | Type     |
|---------------|----------|
| `user_id`     | `String` |
| `plan`        | `String` |
| `country`     | `String` |
| `signup_date` | `Date`   |

---

## Step 1 — Register Both Raw Sources

Sources must be registered before pipes can reference them. Registration provisions the table in ClickHouse and makes the schema available for autocomplete in the pipe editor.

**Register `page_views`:**

```bash
curl -X POST http://localhost:7181/v0/sources \
  -H "Content-Type: application/json" \
  -d '{
    "name": "page_views",
    "engine": "MergeTree()",
    "fields": [
      {"name": "timestamp", "type": "datetime64"},
      {"name": "user_id",   "type": "string"},
      {"name": "page",      "type": "string"}
    ]
  }'
```

**Register `user_profiles`:**

```bash
curl -X POST http://localhost:7181/v0/sources \
  -H "Content-Type: application/json" \
  -d '{
    "name": "user_profiles",
    "engine": "MergeTree()",
    "fields": [
      {"name": "user_id",     "type": "string"},
      {"name": "plan",        "type": "string"},
      {"name": "country",     "type": "string"},
      {"name": "signup_date", "type": "date"}
    ]
  }'
```

Verify both are registered:

```bash
curl http://localhost:7181/v0/sources
# → page_views and user_profiles both appear
```

---

## Step 2 — Pipe 1: Clean Page Views

Create a `TABLE` pipe that filters out bad rows from `page_views` — missing `user_id` or `page` — and casts `timestamp` to a date. Its `destination` becomes a source automatically; no separate registration is needed.

**Pipe content:**

```
type: TABLE
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
      AND page != ''
```

```bash
curl -X POST http://localhost:7181/v0/pipes \
  -H "Content-Type: application/json" \
  -d '{"content": "type: TABLE\nname: clean_page_views\ndescription: Drops incomplete page view rows and normalizes timestamp to date.\ndestination: clean_page_views\n\nsources:\n  - raw: page_views\n\npipeline:\n  @result:\n    SELECT\n      toDate(timestamp) AS day,\n      user_id,\n      page\n    FROM raw\n    WHERE user_id != '\'\''\n      AND page != '\'''\'"}'
```

Verify:

```bash
curl http://localhost:7181/v0/pipes/clean_page_views/meta
# → type: TABLE, destination: clean_page_views
```

---

## Step 3 — Pipe 2: Clean User Profiles

Create a second `TABLE` pipe that filters `user_profiles`, dropping rows with a missing `plan`. Its destination `clean_profiles` also becomes a source automatically.

**Pipe content:**

```
type: TABLE
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
      AND plan != ''
```

```bash
curl -X POST http://localhost:7181/v0/pipes \
  -H "Content-Type: application/json" \
  -d '{"content": "type: TABLE\nname: clean_profiles\ndescription: Drops user profiles with missing plan tier.\ndestination: clean_profiles\n\nsources:\n  - raw: user_profiles\n\npipeline:\n  @result:\n    SELECT\n      user_id,\n      plan,\n      country,\n      signup_date\n    FROM raw\n    WHERE user_id != '\'\''\n      AND plan != '\'''\'"}'
```

Verify:

```bash
curl http://localhost:7181/v0/pipes/clean_profiles/meta
# → type: TABLE, destination: clean_profiles
```

---

## Step 4 — Pipe 3: Join Into Enriched Sessions

Create a `TABLE` pipe that joins the outputs of Pipe 1 and Pipe 2 on `user_id`. Both source names (`views` and `profiles`) reference the destinations that the previous pipes created — not the raw tables.

**Pipe content:**

```
type: TABLE
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
    INNER JOIN profiles ON views.user_id = profiles.user_id
```

```bash
curl -X POST http://localhost:7181/v0/pipes \
  -H "Content-Type: application/json" \
  -d '{"content": "type: TABLE\nname: enrich_sessions\ndescription: Joins clean page views with user profiles to produce enriched sessions.\ndestination: enriched_sessions\n\nsources:\n  - views:    clean_page_views\n  - profiles: clean_profiles\n\npipeline:\n  @result:\n    SELECT\n      views.day,\n      views.user_id,\n      views.page,\n      profiles.plan,\n      profiles.country\n    FROM views\n    INNER JOIN profiles ON views.user_id = profiles.user_id"}'
```

Verify:

```bash
curl http://localhost:7181/v0/pipes/enrich_sessions/meta
# → type: TABLE, destination: enriched_sessions, two sources
```

---

## Step 5 — Pipe 4: Expose as an Endpoint

Create the final `ENDPOINT` pipe that reads from `enriched_sessions` and aggregates daily unique users by plan and country. It accepts two optional filter parameters.

The pipeline has two nodes: `@filter` applies the parameter conditions, `@result` aggregates. `@result` reads from `@filter` — demonstrating intra-pipe node chaining on top of the three-hop cross-pipe chain.

**Pipe content:**

```
type: ENDPOINT
name: daily_stats
description: Daily active users segmented by plan and country, with optional filters.

sources:
  - enriched: enriched_sessions

params:
  plan:      { type: string, default: "" }
  country:   { type: string, default: "" }

pipeline:
  @filter:
    SELECT day, user_id, plan, country
    FROM enriched
    WHERE ({{ params.plan }}    = '' OR plan    = {{ params.plan }})
      AND ({{ params.country }} = '' OR country = {{ params.country }})

  @result:
    SELECT
      day,
      plan,
      country,
      uniqExact(user_id) AS dau
    FROM @filter
    GROUP BY day, plan, country
    ORDER BY day DESC, dau DESC
```

```bash
curl -X POST http://localhost:7181/v0/pipes \
  -H "Content-Type: application/json" \
  -d '{"content": "type: ENDPOINT\nname: daily_stats\ndescription: Daily active users segmented by plan and country, with optional filters.\n\nsources:\n  - enriched: enriched_sessions\n\nparams:\n  plan:    { type: string, default: \"\" }\n  country: { type: string, default: \"\" }\n\npipeline:\n  @filter:\n    SELECT day, user_id, plan, country\n    FROM enriched\n    WHERE ({{ params.plan }} = '\'''\'' OR plan = {{ params.plan }})\n      AND ({{ params.country }} = '\'''\'' OR country = {{ params.country }})\n\n  @result:\n    SELECT\n      day,\n      plan,\n      country,\n      uniqExact(user_id) AS dau\n    FROM @filter\n    GROUP BY day, plan, country\n    ORDER BY day DESC, dau DESC"}'
```

Verify all four pipes are listed:

```bash
curl http://localhost:7181/v0/pipes
# → clean_page_views, clean_profiles, enrich_sessions, daily_stats
```

---

## Step 6 — Execute the Endpoint

Call with no filters to get the full result set:

```bash
curl "http://localhost:7181/v0/pipes/daily_stats"
```

Expected shape:

```json
{
  "data": [
    {"day": "2024-01-16", "plan": "pro",   "country": "US", "dau": 3},
    {"day": "2024-01-16", "plan": "free",  "country": "DE", "dau": 1},
    {"day": "2024-01-15", "plan": "pro",   "country": "US", "dau": 2}
  ]
}
```

Filter by plan:

```bash
curl "http://localhost:7181/v0/pipes/daily_stats?plan=pro"
# → only pro-tier rows

curl "http://localhost:7181/v0/pipes/daily_stats?plan=free&country=DE"
# → free-tier rows in Germany only

curl "http://localhost:7181/v0/pipes/daily_stats?country=JP"
# → empty result set (no JP data)
```

---

## Step 7 — Update the Endpoint (Add a Date Parameter)

Extend `daily_stats` with a `start_day` parameter. Send the full updated pipe content:

```bash
curl -X PUT http://localhost:7181/v0/pipes/daily_stats \
  -H "Content-Type: application/json" \
  -d '{"content": "type: ENDPOINT\nname: daily_stats\ndescription: Daily active users segmented by plan and country, with optional filters.\n\nsources:\n  - enriched: enriched_sessions\n\nparams:\n  plan:      { type: string, default: \"\" }\n  country:   { type: string, default: \"\" }\n  start_day: { type: string, default: \"2000-01-01\" }\n\npipeline:\n  @filter:\n    SELECT day, user_id, plan, country\n    FROM enriched\n    WHERE ({{ params.plan }} = '\'''\'' OR plan = {{ params.plan }})\n      AND ({{ params.country }} = '\'''\'' OR country = {{ params.country }})\n      AND day >= toDate({{ params.start_day }})\n\n  @result:\n    SELECT\n      day,\n      plan,\n      country,\n      uniqExact(user_id) AS dau\n    FROM @filter\n    GROUP BY day, plan, country\n    ORDER BY day DESC, dau DESC"}'
```

Verify the new parameter works:

```bash
curl "http://localhost:7181/v0/pipes/daily_stats?start_day=2024-01-16"
# → only rows from January 16 onward

curl "http://localhost:7181/v0/pipes/daily_stats?plan=pro&start_day=2024-01-16"
# → pro-tier rows from January 16 onward
```

---

## Step 8 — Inspect via the UI

1. Open `http://localhost:5173` and go to **Pipes**
2. All four pipes appear in the list with their types
3. Click **enrich_sessions** — detail view shows two sources (`clean_page_views`, `clean_profiles`) and type `TABLE`
4. Click **daily_stats** — detail view shows two pipeline nodes (`@filter`, `@result`) and three params (`plan`, `country`, `start_day`)

---

## Step 9 — Edit via the UI

1. Click **Edit** on `daily_stats`
2. Add `LIMIT 100` at the end of the `@result` node
3. Click **Save Changes** — confirm redirect to detail view
4. Execute and confirm the response is capped:

```bash
curl "http://localhost:7181/v0/pipes/daily_stats"
# → at most 100 rows
```

---

## Step 10 — Delete All Pipes

Delete downstream first to avoid dangling dependencies:

```bash
curl -X DELETE http://localhost:7181/v0/pipes/daily_stats
curl -X DELETE http://localhost:7181/v0/pipes/enrich_sessions
curl -X DELETE http://localhost:7181/v0/pipes/clean_profiles
curl -X DELETE http://localhost:7181/v0/pipes/clean_page_views

curl http://localhost:7181/v0/pipes
# → empty list
```

---

## Dependency Map

```
page_views (source)          user_profiles (source)
      │                              │
      ▼                              ▼
[Pipe 1: clean_page_views]   [Pipe 2: clean_profiles]
      │  destination→               │  destination→
      ▼                              ▼
 clean_page_views              clean_profiles
      │                              │
      └──────────────┬───────────────┘
                     ▼
          [Pipe 3: enrich_sessions]
                     │  destination→
                     ▼
            enriched_sessions
                     │
                     ▼
           [Pipe 4: daily_stats]   type=ENDPOINT
                     │  exposes
                     ▼
GET /v0/pipes/daily_stats?plan=pro&country=US&start_day=2024-01-16
```

---

## Verification Checklist

- [ ] Both sources registered and visible in `GET /v0/sources`
- [ ] `clean_page_views` pipe: `type: TABLE`, `destination: clean_page_views`
- [ ] `clean_profiles` pipe: `type: TABLE`, `destination: clean_profiles`
- [ ] `enrich_sessions` pipe: `type: TABLE`, two sources, `destination: enriched_sessions`
- [ ] `daily_stats` pipe: `type: ENDPOINT`, two nodes, two params
- [ ] `GET /v0/pipes/daily_stats` returns rows aggregated across the full chain
- [ ] `?plan=pro` and `?country=DE` filter correctly; `?country=JP` returns empty
- [ ] `PUT` adds `start_day`; `?start_day=2024-01-16` filters by date
- [ ] UI shows all four pipes; detail view for `enrich_sessions` shows two sources
- [ ] UI edit on `daily_stats` saves and reflects in API response
- [ ] All four pipes deleted in downstream-first order
