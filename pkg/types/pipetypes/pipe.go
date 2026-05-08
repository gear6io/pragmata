package pipetypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gear6io/pragmata/pkg/valuer"
	"github.com/uptrace/bun"
)

// PipeType is the execution model for a pipe.
type PipeType struct {
	valuer.String
}

var (
	PipeTypeUndefined    = PipeType{valuer.NewString("")}
	PipeTypeEndpoint     = PipeType{valuer.NewString("ENDPOINT")}
	PipeTypeMaterialized = PipeType{valuer.NewString("MATERIALIZED")}
	PipeTypeCopy         = PipeType{valuer.NewString("COPY")}
)

// Node is one SQL step in a pipe. Nodes are chained into a CTE at query time.
type Node struct {
	Name        string `json:"name"`
	SQL         string `json:"sql"`
	IsTemplated bool   `json:"isTemplated"`
}

// Nodes is a JSON-serialised slice of Node, stored as TEXT in SQLite.
type Nodes []Node

func (n Nodes) Value() (driver.Value, error) {
	if len(n) == 0 {
		return "[]", nil
	}
	b, err := json.Marshal([]Node(n))
	return string(b), err
}

func (n *Nodes) Scan(src any) error {
	if src == nil {
		*n = Nodes{}
		return nil
	}
	var b []byte
	switch v := src.(type) {
	case []byte:
		b = v
	case string:
		b = []byte(v)
	default:
		return fmt.Errorf("Nodes: cannot scan %T", src)
	}
	return json.Unmarshal(b, (*[]Node)(n))
}

// Tags is a JSON-serialised string slice, stored as TEXT in SQLite.
type Tags []string

func (t Tags) Value() (driver.Value, error) {
	if len(t) == 0 {
		return "[]", nil
	}
	b, err := json.Marshal([]string(t))
	return string(b), err
}

func (t *Tags) Scan(src any) error {
	if src == nil {
		*t = Tags{}
		return nil
	}
	var b []byte
	switch v := src.(type) {
	case []byte:
		b = v
	case string:
		b = []byte(v)
	default:
		return fmt.Errorf("Tags: cannot scan %T", src)
	}
	return json.Unmarshal(b, (*[]string)(t))
}

// Pipe is the core abstraction — a named chain of SQL nodes with an execution type.
type Pipe struct {
	bun.BaseModel    `bun:"table:pipes"`
	Name             string    `bun:"name,pk"                                       json:"name"`
	Type             PipeType  `bun:"type,notnull"                                  json:"type"`
	Description      string    `bun:"description"                                   json:"description,omitempty"`
	Tags             Tags      `bun:"tags"                                          json:"tags,omitempty"`
	Nodes            Nodes     `bun:"nodes,notnull"                                 json:"nodes"`
	Datasource       string    `bun:"datasource"                                    json:"datasource,omitempty"`
	TargetDatasource string    `bun:"target_datasource"                             json:"targetDatasource,omitempty"`
	CopySchedule     string    `bun:"copy_schedule"                                 json:"copySchedule,omitempty"`
	CreatedAt        time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"-"`
	UpdatedAt        time.Time `bun:"updated_at,nullzero,default:current_timestamp" json:"-"`
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
	RowsRead uint64  `json:"rowsRead"`
}

// ExecuteResult is the response envelope for a pipe execution.
type ExecuteResult struct {
	Data  []map[string]any `json:"data"`
	Meta  ResultMeta       `json:"meta"`
	Stats ResultStats      `json:"stats"`
}
