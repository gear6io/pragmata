package pipevisitor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
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
		check    func(t *testing.T, p *pipetypes.ExecutablePipe)
	}{
		// ── happy paths ──────────────────────────────────────────────────────────
		{
			name:     "full endpoint: all directives, aliased+simple sources, 3-node pipeline",
			pipeName: "ignored",
			file:     "analytics_endpoint.pipe",
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, pipetypes.PipeTypeEndpoint, p.Type)
				// name directive in file overrides the Visit parameter
				require.Equal(t, "dau_report", p.Name)
				// multi-line description is joined with a single space
				require.Equal(t, "Daily active user count grouped by country.", p.Description)
				require.Equal(t, pipetypes.Tags{"analytics", "dashboard", "pii"}, p.Tags)
				require.Equal(t, "data-team@acme.com", p.Owner)
				require.Equal(t, "reporting.dau", p.Destination)
				require.Equal(t, "every 24 hours", p.Schedule)
				require.Equal(t, pipetypes.Sources{
					{Alias: "evt", Table: "raw.events"},
					{Alias: "users", Table: "users"},
				}, p.Sources)
				require.Equal(t, pipetypes.ParamDefs{
					{Name: "country", DataType: "string", DefaultValue: "US"},
					{Name: "days_back", DataType: "int64", DefaultValue: "30"},
				}, p.Params)
				require.Len(t, p.Nodes, 3)
				require.Equal(t, "filtered", p.Nodes[0].Name)
				require.Equal(t, "enriched", p.Nodes[1].Name)
				require.Equal(t, "result", p.Nodes[2].Name)
			},
		},
		{
			name:     "incremental: single-line desc, node-to-node @ref, param with numeric default",
			pipeName: "incremental_events",
			file:     "incremental_events.pipe",
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, pipetypes.PipeTypeIncremental, p.Type)
				// no name directive; pipeName parameter is kept
				require.Equal(t, "incremental_events", p.Name)
				require.Equal(t, "Hourly event aggregations, append-only.", p.Description)
				require.Empty(t, p.Tags)
				require.Empty(t, p.Owner)
				require.Empty(t, p.Destination)
				require.Empty(t, p.Schedule)
				require.Equal(t, pipetypes.Sources{
					{Alias: "events", Table: "events"},
					{Alias: "meta", Table: "ref.event_metadata"},
				}, p.Sources)
				require.Equal(t, pipetypes.ParamDefs{
					{Name: "lookback_hours", DataType: "int64", DefaultValue: "6"},
				}, p.Params)
				require.Len(t, p.Nodes, 2)
				require.Equal(t, "raw", p.Nodes[0].Name)
				require.Equal(t, "final", p.Nodes[1].Name)
			},
		},
		{
			name:     "minimal: only type + pipeline, all optional fields are zero",
			pipeName: "minimal_pipe",
			content: `type: VIEW
pipeline:
  @main:
    from users
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, pipetypes.PipeTypeView, p.Type)
				require.Equal(t, "minimal_pipe", p.Name)
				require.Empty(t, p.Description)
				require.Empty(t, p.Tags)
				require.Empty(t, p.Owner)
				require.Empty(t, p.Sources)
				require.Empty(t, p.Params)
				require.Len(t, p.Nodes, 1)
				require.Equal(t, "main", p.Nodes[0].Name)
			},
		},
		{
			name:     "name directive overrides the Visit name parameter",
			pipeName: "file_derived_name",
			content: `type: TABLE
name: override_name
pipeline:
  @main:
    from orders
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, "override_name", p.Name)
			},
		},
		{
			name:     "multi-line description: blank lines ignored, non-blank lines joined with space",
			pipeName: "ml_desc",
			content: `type: VIEW
description: |
  First line of description.
  Second line of description.
pipeline:
  @main:
    from users
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, "First line of description. Second line of description.", p.Description)
			},
		},
		{
			name:     "tags: comma-separated, each entry trimmed of whitespace",
			pipeName: "tagged",
			content: `type: VIEW
tags: [analytics, finance, pii]
pipeline:
  @main:
    from users
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, pipetypes.Tags{"analytics", "finance", "pii"}, p.Tags)
			},
		},
		{
			name:     "param with no default: DefaultValue is empty string",
			pipeName: "no_default",
			content: `type: VIEW
params:
  region: { type: string }
pipeline:
  @main:
    from users
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, pipetypes.ParamDefs{
					{Name: "region", DataType: "string", DefaultValue: ""},
				}, p.Params)
			},
		},
		{
			name:     "simple source: Alias and Table are both set to the bare identifier",
			pipeName: "simple_src",
			content: `type: VIEW
sources:
  - events
pipeline:
  @main:
    from events
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, pipetypes.Sources{{Alias: "events", Table: "events"}}, p.Sources)
			},
		},
		{
			name:     "aliased source: Alias differs from Table, dotted table path preserved",
			pipeName: "aliased_src",
			content: `type: VIEW
sources:
  - ev: schema.events
pipeline:
  @main:
    from ev
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, pipetypes.Sources{{Alias: "ev", Table: "schema.events"}}, p.Sources)
			},
		},
		{
			name:     "unknown type: stored verbatim as lowercase custom PipeType",
			pipeName: "custom_type",
			content: `type: STREAMING
pipeline:
  @main:
    from events
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Equal(t, "streaming", p.Type.StringValue())
			},
		},
		{
			name:     "multi-node pipeline: inline SQL spans multiple lines, accumulated per node",
			pipeName: "multi_node",
			content: `type: TABLE
pipeline:
  @step1:
    from orders
    filter status == "active"
  @step2:
    from @step1
    select { id, status }
`,
			check: func(t *testing.T, p *pipetypes.ExecutablePipe) {
				require.Len(t, p.Nodes, 2)
				require.Equal(t, "step1", p.Nodes[0].Name)
				require.Contains(t, p.Nodes[0].SQL, "from orders")
				require.Contains(t, p.Nodes[0].SQL, `filter status == "active"`)
				require.Equal(t, "step2", p.Nodes[1].Name)
				require.Contains(t, p.Nodes[1].SQL, "from @step1")
			},
		},
		// ── error paths ──────────────────────────────────────────────────────────
		{
			name:     "missing type: Visit returns error naming the pipe",
			pipeName: "no_type_pipe",
			content: `pipeline:
  @main:
    from users
`,
			wantErr: "has no type declaration",
		},
		{
			name:    "no pipeline section: Visit returns error when Nodes is empty",
			content: `type: VIEW
sources:
  - users
`,
			wantErr: "has no pipeline nodes",
		},
		{
			name:    "empty pipeline node name: @: triggers error",
			content: `type: VIEW
pipeline:
  @:
    from users
`,
			wantErr: "pipeline node has empty name",
		},
		{
			name:    "lex error: unrecognized character at column zero",
			content: "@ bad content\n",
			wantErr: "lex error",
		},
		{
			name: "opts: SourceValidator set but FetchSources nil returns error before parsing",
			opts: PipeVisitorOpts{
				SourceValidator: func(_ string, _ []sourcetypes.Source) error { return nil },
			},
			wantErr: "FetchSources can not be nil",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			content := tc.content
			if tc.file != "" {
				content = mustReadTestdata(t, tc.file)
			}
			got, err := Visit(tc.pipeName, content, tc.opts)
			if tc.wantErr != "" {
				require.ErrorContains(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, got)
			if tc.check != nil {
				tc.check(t, got)
			}
		})
	}
}
