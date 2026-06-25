package pipetypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gear6io/pragmata/pkg/types"
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
	PipeTypeTable        = PipeType{valuer.NewString("TABLE")}
	PipeTypeView         = PipeType{valuer.NewString("VIEW")}
	PipeTypeIncremental  = PipeType{valuer.NewString("INCREMENTAL")}
	PipeTypeSnapshot     = PipeType{valuer.NewString("SNAPSHOT")}
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

// Source is one entry in the sources: section — an alias mapped to an actual table.
type Source struct {
	Alias string `json:"alias"`
	Table string `json:"table"`
}

// Sources is a JSON-serialised slice of Source, stored as TEXT in SQLite.
type Sources []Source

func (s Sources) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	b, err := json.Marshal([]Source(s))
	return string(b), err
}

func (s *Sources) Scan(src any) error {
	if src == nil {
		*s = Sources{}
		return nil
	}
	var b []byte
	switch v := src.(type) {
	case []byte:
		b = v
	case string:
		b = []byte(v)
	default:
		return fmt.Errorf("Sources: cannot scan %T", src)
	}
	return json.Unmarshal(b, (*[]Source)(s))
}

// ParamDef describes one parameter in the params: section.
type ParamDef struct {
	Name         string `json:"name"`
	DataType     string `json:"dataType"`
	DefaultValue string `json:"defaultValue"`
}

// ParamDefs is a JSON-serialised slice of ParamDef, stored as TEXT in SQLite.
type ParamDefs []ParamDef

func (p ParamDefs) Value() (driver.Value, error) {
	if len(p) == 0 {
		return "[]", nil
	}
	b, err := json.Marshal([]ParamDef(p))
	return string(b), err
}

func (p *ParamDefs) Scan(src any) error {
	if src == nil {
		*p = ParamDefs{}
		return nil
	}
	var b []byte
	switch v := src.(type) {
	case []byte:
		b = v
	case string:
		b = []byte(v)
	default:
		return fmt.Errorf("ParamDefs: cannot scan %T", src)
	}
	return json.Unmarshal(b, (*[]ParamDef)(p))
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
	Name             string   `bun:"name,pk" json:"name"`
	Type             PipeType `bun:"type,notnull" json:"type"`
	Description      string   `bun:"description" json:"description,omitempty"`
	Tags             Tags     `bun:"tags" json:"tags,omitempty"`
	Content          string   `bun:"content,notnull" json:"content"`
	Datasource       string   `bun:"datasource" json:"datasource,omitempty"`
	TargetDatasource string   `bun:"target_datasource" json:"targetDatasource,omitempty"`
	CopySchedule     string   `bun:"copy_schedule" json:"copySchedule,omitempty"`

	// Virtual / computed fields
	Nodes       Nodes     `bun:"-" json:"-"`
	Owner       string    `bun:"-" json:"-"`
	Destination string    `bun:"-" json:"-"`
	Schedule    string    `bun:"-" json:"-"`
	Sources     Sources   `bun:"-" json:"-"`
	Params      ParamDefs `bun:"-" json:"-"`
}

type StorablePipe struct {
	bun.BaseModel `bun:"table:pipes"`
	types.Identifiable
	types.UserAuditable
	types.TimeAuditable

	Pipe
}
