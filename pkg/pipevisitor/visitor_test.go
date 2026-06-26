package pipevisitor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/stretchr/testify/require"
)

func mustReadTestdata(t *testing.T, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", name))
	require.NoError(t, err)
	return string(b)
}

func TestVisit(t *testing.T) {
	tests := []struct {
		name     string
		pipeName string
		file     string // load from testdata/ when set; content field ignored
		content  string
		opts     PipeVisitorOpts
		wantErr  string
		expected pipetypes.ExecutablePipe
	}{
		// ── happy paths ──────────────────────────────────────────────────────────
		{
			name:     "full endpoint: all directives, aliased+simple sources, 3-node pipeline",
			pipeName: "ignored",
			file:     "analytics_endpoint.pipe",
			expected: pipetypes.ExecutablePipe{
				Pipe: pipetypes.Pipe{
					Name:        "dau_report", // name: directive overrides pipeName param
					Type:        pipetypes.PipeTypeEndpoint,
					Description: "Daily active user count grouped by country.",
					Tags:        pipetypes.Tags{"analytics", "dashboard", "pii"},
				},
				Owner:       "data-team@acme.com",
				Destination: "reporting.dau",
				Schedule:    "every 24 hours",
				Sources: pipetypes.Sources{
					{Alias: "evt", Table: "raw.events"},
					{Alias: "users", Table: "users"},
				},
				Params: pipetypes.ParamDefs{
					{Name: "country", DataType: "string", DefaultValue: "US"},
					{Name: "days_back", DataType: "int64", DefaultValue: "30"},
				},
				Nodes: pipetypes.Nodes{
					{Name: "filtered", SQL: "from evt filter country == $country"},
					{Name: "enriched", SQL: "from @filtered join users (==user_id) select { user_id, country, ts }"},
					{Name: "result", SQL: "from @enriched group { country } ( aggregate { cnt = count user_id } ) sort cnt desc"},
				},
			},
		},
		{
			name:     "incremental: single-line desc, node-to-node @ref, param with numeric default",
			pipeName: "incremental_events",
			file:     "incremental_events.pipe",
			expected: pipetypes.ExecutablePipe{
				Pipe: pipetypes.Pipe{
					Name:        "incremental_events", // no name: directive; pipeName kept
					Type:        pipetypes.PipeTypeIncremental,
					Description: "Hourly event aggregations, append-only.",
				},
				Sources: pipetypes.Sources{
					{Alias: "events", Table: "events"},
					{Alias: "meta", Table: "ref.event_metadata"},
				},
				Params: pipetypes.ParamDefs{
					{Name: "lookback_hours", DataType: "int64", DefaultValue: "6"},
				},
				Nodes: pipetypes.Nodes{
					{Name: "raw", SQL: "from events select { event_id, ts, kind }"},
					{Name: "final", SQL: "from @raw join meta (==event_id) select { event_id, ts, kind, meta.label }"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			content := tc.content
			if tc.file != "" {
				content = mustReadTestdata(t, tc.file)
			}
			got, err := Visit(content, tc.opts)
			if tc.wantErr != "" {
				require.ErrorContains(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			expected := tc.expected
			expected.Pipe.Content = content
			require.Equal(t, &expected, got)
		})
	}
}
