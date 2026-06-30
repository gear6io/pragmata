package pipetypes

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/types"
	"github.com/gear6io/pragmata/pkg/types/querybuildertypes"
	"github.com/gear6io/pragmata/pkg/valuer"
	"github.com/uptrace/bun"
)

// PipeType is the execution model for a pipe.
type PipeType struct {
	valuer.RawString
}

var (
	PipeTypeUndefined    = PipeType{valuer.NewRawString("")}
	PipeTypeEndpoint     = PipeType{valuer.NewRawString("ENDPOINT")}
	PipeTypeMaterialized = PipeType{valuer.NewRawString("MATERIALIZED")}
	PipeTypeCopy         = PipeType{valuer.NewRawString("COPY")}
	PipeTypeTable        = PipeType{valuer.NewRawString("TABLE")}
	PipeTypeView         = PipeType{valuer.NewRawString("VIEW")}
	PipeTypeIncremental  = PipeType{valuer.NewRawString("INCREMENTAL")}
	PipeTypeSnapshot     = PipeType{valuer.NewRawString("SNAPSHOT")}
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
		return errors.NewInternalf(errors.CodeInternal, "Nodes: cannot scan %T", src)
	}
	return json.Unmarshal(b, (*[]Node)(n))
}

// Source is one entry in the sources: section — an alias mapped to an actual table.
type Source struct {
	Alias string `json:"alias"`
	Table string `json:"table"`
}

func (s *Source) String() string {
	if s.Alias != "" {
		return s.Alias
	}
	return s.Table
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
		return errors.NewInternalf(errors.CodeInternal, "Sources: cannot scan %T", src)
	}
	return json.Unmarshal(b, (*[]Source)(s))
}

// ParamDef describes one parameter in the params: section.
type ParamDef struct {
	Name         string                          `json:"name"`
	DataType     querybuildertypes.FieldDataType `json:"dataType"`
	DefaultValue string                          `json:"defaultValue"`
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
		return errors.NewInternalf(errors.CodeInternal, "ParamDefs: cannot scan %T", src)
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
		return errors.NewInternalf(errors.CodeInternal, "Tags: cannot scan %T", src)
	}
	return json.Unmarshal(b, (*[]string)(t))
}

// Pipe holds the persisted fields for a pipe definition.
type Pipe struct {
	Name             string   `bun:"name,pk" json:"name"`
	Type             PipeType `bun:"type,notnull" json:"type"`
	Description      string   `bun:"description" json:"description,omitempty"`
	Tags             Tags     `bun:"tags" json:"tags,omitempty"`
	Content          string   `bun:"content,notnull" json:"content"`
	Datasource       string   `bun:"datasource" json:"datasource,omitempty"`
	TargetDatasource string   `bun:"target_datasource" json:"targetDatasource,omitempty"`
	CopySchedule     string   `bun:"copy_schedule" json:"copySchedule,omitempty"`
}

// GettablePipe is what the API returns to clients.
// It exposes the full StorablePipe including ID and audit fields.
type GettablePipe = StorablePipe

// ExecutablePipe is a parsed pipe ready for execution.
// It extends Pipe with fields derived by parsing Content at runtime.
type ExecutablePipe struct {
	Pipe
	Nodes       Nodes
	Owner       string
	Destination string
	Schedule    string
	Sources     Sources
	Params      ParamDefs
}

// PostablePipe is the create/update request body for pipe endpoints.
// The raw PipeLang content is parsed server-side to populate all pipe fields.
type PostablePipe struct {
	Content string `json:"content"`
}

// ExecuteResult is the response shape for ENDPOINT pipe execution.
type ExecuteResult struct {
	Data []map[string]any `json:"data"`
}

type StorablePipe struct {
	bun.BaseModel `bun:"table:pipes"`
	types.Identifiable
	types.UserAuditable
	types.TimeAuditable

	Pipe
}
