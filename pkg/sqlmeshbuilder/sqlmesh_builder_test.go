package sqlmeshbuilder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromContent(t *testing.T) {
	tests := []struct {
		name     string
		pipeName string
		content  string
		want     string
	}{
		{
			name:     "simple view with filter",
			pipeName: "user_filter",
			content: `type: VIEW
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromContent(tt.pipeName, tt.content)
			if err != nil {
				require.NoError(t, err)
			}
			if got != tt.want {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
