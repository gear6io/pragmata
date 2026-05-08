package pipetypes

// PipeType is the execution model for a pipe.
type PipeType string

const (
	PipeTypeEndpoint     PipeType = "ENDPOINT"
	PipeTypeMaterialized PipeType = "MATERIALIZED"
	PipeTypeCopy         PipeType = "COPY"
)

// Node is one SQL step in a pipe. Nodes are chained into a CTE at query time.
type Node struct {
	Name        string `json:"name"`
	SQL         string `json:"sql"`
	IsTemplated bool   `json:"is_templated"` // true when the SQL block starts with %
}

// Pipe is the core abstraction — a named chain of SQL nodes with an execution type.
type Pipe struct {
	Name             string   `json:"name"`
	Description      string   `json:"description,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	Nodes            []Node   `json:"nodes"`
	Type             PipeType `json:"type"`
	Datasource       string   `json:"datasource,omitempty"`        // MATERIALIZED: target datasource
	TargetDatasource string   `json:"target_datasource,omitempty"` // COPY: destination table
	CopySchedule     string   `json:"copy_schedule,omitempty"`     // COPY: cron expression
}

// Column describes a single column in a query result.
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ResultMeta holds column metadata for a query result.
type ResultMeta struct {
	Columns []Column `json:"columns"`
}

// ResultStats holds performance information for a query result.
type ResultStats struct {
	Elapsed  float64 `json:"elapsed"`
	RowsRead uint64  `json:"rows_read"`
}

// ExecuteResult is the response envelope for a pipe execution.
type ExecuteResult struct {
	Data  []map[string]any `json:"data"`
	Meta  ResultMeta       `json:"meta"`
	Stats ResultStats      `json:"stats"`
}
