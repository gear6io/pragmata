package querier

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/stretchr/testify/require"
)

func mustParsePipe(t *testing.T, name string) *pipetypes.ExecutablePipe {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", name))
	require.NoError(t, err)
	pipe, err := pipevisitor.Visit(string(b), pipevisitor.PipeVisitorOpts{})
	require.NoError(t, err)
	return pipe
}

func TestBuildSQL(t *testing.T) {
	tests := []struct {
		name        string
		pipe        func() *pipetypes.ExecutablePipe
		urlParams   map[string]string
		expectedSQL string
		expectedErr string
	}{
		{
			name: "multi-node with aliased and simple sources",
			pipe: func() *pipetypes.ExecutablePipe { return mustParsePipe(t, "simple_endpoint.pipe") },
			expectedSQL: "WITH events AS (SELECT * FROM pragmata_source.events), " +
				"users AS (SELECT * FROM pragmata_source.users), " +
				"raw AS (SELECT event_id, user_id, ts FROM events) " +
				"SELECT event_id, user_id, ts FROM raw JOIN users USING (user_id)",
		},
		{
			name:        "single node with params, default values",
			pipe:        func() *pipetypes.ExecutablePipe { return mustParsePipe(t, "endpoint_with_params.pipe") },
			expectedSQL: "WITH evt AS (SELECT * FROM pragmata_source.raw.events) SELECT event_id, ts FROM evt WHERE status ==  'active'",
		},
		{
			name:        "single node with params, url params override defaults",
			pipe:        func() *pipetypes.ExecutablePipe { return mustParsePipe(t, "endpoint_with_params.pipe") },
			urlParams:   map[string]string{"status": "inactive"},
			expectedSQL: "WITH evt AS (SELECT * FROM pragmata_source.raw.events) SELECT event_id, ts FROM evt WHERE status ==  'inactive'",
		},
		{
			name: "multi-node with group and sort, default params",
			pipe: func() *pipetypes.ExecutablePipe { return mustParsePipe(t, "daily_stats.pipe") },
			expectedSQL: "WITH enriched AS (SELECT * FROM pragmata_source.enriched_sessions), " +
				"filtered AS (SELECT day, user_id, plan, country FROM enriched WHERE ( ''  = '' OR plan =  '' ) AND ( ''  = '' OR country =  '' )) " +
				"SELECT day, plan, country, uniqExact(user_id) AS dau FROM filtered GROUP BY day, plan, country ORDER BY day DESC, dau DESC",
		},
		{
			name:      "multi-node with group and sort, plan param set",
			pipe:      func() *pipetypes.ExecutablePipe { return mustParsePipe(t, "daily_stats.pipe") },
			urlParams: map[string]string{"plan": "pro"},
			expectedSQL: "WITH enriched AS (SELECT * FROM pragmata_source.enriched_sessions), " +
				"filtered AS (SELECT day, user_id, plan, country FROM enriched WHERE ( 'pro'  = '' OR plan =  'pro' ) AND ( ''  = '' OR country =  '' )) " +
				"SELECT day, plan, country, uniqExact(user_id) AS dau FROM filtered GROUP BY day, plan, country ORDER BY day DESC, dau DESC",
		},
		{
			name: "empty nodes returns error",
			pipe: func() *pipetypes.ExecutablePipe {
				return &pipetypes.ExecutablePipe{Pipe: pipetypes.Pipe{Name: "empty"}}
			},
			expectedErr: `pipe "empty" has no nodes`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sql, err := BuildSQL(tc.pipe(), tc.urlParams)
			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedSQL, sql)
		})
	}
}
