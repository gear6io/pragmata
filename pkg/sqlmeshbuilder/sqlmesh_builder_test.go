package sqlmeshbuilder

import (
	"testing"

	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/stretchr/testify/require"
)

func TestFromPipe(t *testing.T) {
	tests := []struct {
		name     string
		pipeName string
		content  string
		want     string
		wantErr  string
	}{
		{
			name:     "simple view with filter",
			pipeName: "user_filter",
			content: `type: VIEW
destination: user_filter
pipeline:
  @main:
    from users
    filter age > 18
`,
			want: `MODEL (
  name = user_filter,
  kind = VIEW,
  dialect = 'clickhouse',
);

SELECT * FROM users WHERE age > 18
`,
		},
		{
			name:     "materialized type maps to FULL kind",
			pipeName: "daily_summary",
			content: `type: MATERIALIZED
destination: daily_summary
pipeline:
  @main:
    from events
`,
			want: `MODEL (
  name = daily_summary,
  kind = FULL,
  dialect = 'clickhouse',
);

SELECT * FROM events
`,
		},
		{
			name:     "optional props description tags schedule",
			pipeName: "tagged_view",
			content: `type: VIEW
destination: tagged_view
description: Active users
tags: [analytics, monitoring]
schedule: @daily
pipeline:
  @main:
    from users
`,
			want: `MODEL (
  name = tagged_view,
  kind = VIEW,
  dialect = 'clickhouse',
  description = 'Active users',
  tags = ['analytics', 'monitoring'],
  cron = '@daily',
);

SELECT * FROM users
`,
		},
		{
			name:     "multi-node pipe generates CTE",
			pipeName: "cte_pipe",
			content: `type: VIEW
destination: cte_pipe
pipeline:
  @filtered:
    from orders
    filter status = 'active'
  @main:
    from filtered
`,
			want: `MODEL (
  name = cte_pipe,
  kind = VIEW,
  dialect = 'clickhouse',
);

WITH filtered AS (SELECT * FROM orders WHERE status = 'active') SELECT * FROM filtered
`,
		},
		{
			name:     "select sort take pipeline",
			pipeName: "top_users",
			content: `type: VIEW
destination: top_users
pipeline:
  @main:
    from users
    select { id, name, age }
    sort { -age }
    take 5
`,
			want: `MODEL (
  name = top_users,
  kind = VIEW,
  dialect = 'clickhouse',
);

SELECT id, name, age FROM users ORDER BY age DESC LIMIT ?
`,
		},
		{
			name:     "source alias emits CTE before node CTEs",
			pipeName: "source_pipe",
			content: `type: VIEW
destination: source_pipe
sources:
  - views
pipeline:
  @main:
    from views
`,
			want: `MODEL (
  name = source_pipe,
  kind = VIEW,
  dialect = 'clickhouse',
);

WITH views AS (SELECT * FROM pragmata_source.views) SELECT * FROM views
`,
		},
		{
			name:     "multi-source explicit join",
			pipeName: "enrich_sessions",
			content: `type: TABLE
destination: enriched_sessions
sources:
  - views: clean_page_views
  - profiles: clean_profiles
pipeline:
  @result:
    from views
    join side:inner profiles (views.user_id = profiles.user_id)
    select { views.day, views.user_id, views.page, profiles.plan, profiles.country }
`,
			want: `MODEL (
  name = enriched_sessions,
  kind = FULL,
  dialect = 'clickhouse',
);

WITH views AS (SELECT * FROM pragmata_source.clean_page_views), profiles AS (SELECT * FROM pragmata_source.clean_profiles) SELECT views.day, views.user_id, views.page, profiles.plan, profiles.country FROM views INNER JOIN profiles ON views.user_id=profiles.user_id
`,
		},
		{
			name:     "missing destination returns error",
			pipeName: "no_dest",
			content: `type: VIEW
pipeline:
  @main:
    from users
`,
			wantErr: `must set destination:`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromPipe(&pipetypes.Pipe{Name: tt.pipeName, Content: tt.content})
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
