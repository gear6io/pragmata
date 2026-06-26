# ClickHouse Transformation Platform

Define data pipelines as versioned files. Run them against ClickHouse. Expose results as REST endpoints — without dbt, without SQLMesh config, without managed-only lock-in.

Self-hostable. Bring your own ClickHouse cluster.

---

## How it works

Two primitives:

**Sources** — either a raw ClickHouse table receiving data, or the output destination of an existing pipeline. Sources are the inputs to everything else.

**Pipe files** — versioned `.pipe` files that declare what to read, how to transform it, and what to produce. The output of a pipe becomes a source for other pipes, building a DAG without a separate orchestration layer.

Three pipe types:

| Type | What it does |
|------|-------------|
| `endpoint` | Parameterized query exposed as a live REST API |
| `materialized` | Pre-computed ClickHouse materialized view |
| `incremental` | Incremental append-only model, managed via SQLMesh |
| `copy` | Scheduled SQL-to-table write, runs on cron |

---

## Pipe file format

Pipe files have two sections: YAML directives for metadata, and a `pipeline:` block of PRQL nodes.

```yaml
type: endpoint
name: revenue_report
description: Daily revenue by country
tags: [revenue, sales]
sources:
  - orders
params:
  country: { type: string, default: "US" }
  limit:   { type: int64,  default: 100  }
pipeline:
  @filtered:
    from orders
    filter country == $country

  @result:
    from @filtered
    group { country } (
      aggregate { revenue = sum amount, orders = count() }
    )
    sort revenue desc
    take $limit
```

Each `@nodename:` block is a PRQL pipeline. Nodes reference each other with `@` — `from @filtered` reads the output of the `filtered` node. The last node's query becomes the final SELECT; all preceding nodes compile to CTEs.

Parameters are declared in the `params:` directive and referenced as `$param` in PRQL.

### Sources can be aliased

```yaml
sources:
  - evt: raw.events   # aliased: use 'evt' in pipeline
  - users             # simple: use 'users' directly
```

### Full example with join

```yaml
type: endpoint
name: dau_report
description: |
  Daily active user count
  grouped by country.
tags: [analytics, dashboard]
owner: data-team@acme.com
destination: reporting.dau
schedule: every 24 hours
sources:
  - evt: raw.events
  - users
params:
  country: { type: string, default: "US" }
  days_back: { type: int64, default: 30 }
pipeline:
  @filtered:
    from evt
    filter country == $country

  @enriched:
    from @filtered
    join users (==user_id)
    select { user_id, country, ts }

  @result:
    from @enriched
    group { country } (
      aggregate { cnt = count user_id }
    )
    sort cnt desc
```

---

## PRQL — the query language

Node bodies are written in PRQL, a pipelined query language that compiles to ClickHouse SQL. Transforms read top-to-bottom in logical order:

| Transform | What it does |
|-----------|-------------|
| `from table` | Read from a source or `@node` |
| `filter expr` | WHERE clause |
| `select { col, alias = expr }` | Projection |
| `derive { alias = expr }` | Add computed columns |
| `group { key } ( aggregate { ... } )` | GROUP BY + aggregates |
| `sort col desc` / `sort { -col }` | ORDER BY |
| `take n` | LIMIT n |
| `skip n` | OFFSET n |
| `join table (==col)` | JOIN using shared key |
| `join side:left table (col = table.col)` | Explicit JOIN condition |
| `array_join col as alias` | ClickHouse ARRAY JOIN |
| `window { alias = expr }` | Window functions |

`from orders final` maps to ClickHouse `FINAL` for ReplacingMergeTree.

---

## HTTP API

Every `endpoint` pipe is served immediately:

```
GET /v0/pipes/{name}
GET /v0/pipes/{name}.json
GET /v0/pipes/{name}.csv
GET /v0/pipes/{name}.csvwithnames
```

```json
{
  "data": [{ "country": "US", "revenue": 45230.5 }],
  "meta": { "columns": [{ "name": "country", "type": "String" }] },
  "stats": { "elapsed": 0.032, "rows_read": 50000 }
}
```

Auth: `Authorization: Bearer <token>` or `?token=<token>`.

---

## Quick start

**Prerequisites**: Go 1.25+, Docker

```bash
# Start ClickHouse
make dev-up

# Start the server (reads .vscode/config.yaml by default)
make run

# Optional: frontend dev server
make dev-fe
```

Single binary with embedded frontend:

```bash
make build
./pragmata serve
```

---

## CLI

```bash
pragmata serve                        # Start HTTP server (default :7181)
pragmata push <dir>                   # Load/reload pipes from a directory
pragmata run <name> [--param k=v]     # Execute a pipe manually, print results
pragmata token create --name <n>      # Create an API token
pragmata token list
pragmata generate openapi             # Emit OpenAPI spec
```

---

## Architecture

```
┌──────────────────────────────────────────────────────┐
│                    Go binary                          │
│                                                       │
│  HTTP /v0/pipes/{name}                                │
│    → token auth                                       │
│    → pipe parser  (ANTLR4 grammar → AST)              │
│    → PRQL visitor (nodes → CTE chain)                 │
│    → executor  (ClickHouse SQL)                       │
│    → formatter  (JSON / CSV / CSVWithNames)           │
│                                                       │
│  Cron scheduler → SQLMesh subprocess                 │
│  (materialized + incremental + copy pipes)           │
└──────────────────────────────────────────────────────┘
                │ clickhouse-go/v2
         ┌──────▼──────────┐
         │   ClickHouse    │
         │ (any instance)  │
         └─────────────────┘

Internal state: SQLite  |  Frontend: React + Vite (embedded)
```

---

## Development

```bash
make test           # Go tests
make generate-api   # Regenerate OpenAPI spec + TypeScript client
make dev-clean      # Stop Docker services and delete volumes
```
