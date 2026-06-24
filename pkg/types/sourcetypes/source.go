package sourcetypes

import (
	"github.com/gear6io/pragmata/pkg/types/querybuildertypes"
	"github.com/gear6io/pragmata/pkg/valuer"
)

type Engine struct {
	valuer.RawString
}

var (
	EngineUndefined          = Engine{valuer.NewRawString("")}
	EngineMergeTree          = Engine{valuer.NewRawString("MergeTree()")}
	EngineReplacingMergeTree = Engine{valuer.NewRawString("ReplacingMergeTree()")}
	EngineSummingMergeTree   = Engine{valuer.NewRawString("SummingMergeTree()")}
)

// Source is a table in the pragmata_source ClickHouse database.
type Source struct {
	Name     string `json:"name"`
	Database string `json:"-"` // internal — always "pragmata_source"
	Engine   Engine `json:"engine,omitempty"`
	// User level fields; These are top-level columns
	//
	// along with these _attrs_ is JSON() column, accpeting additional
	// fields ingested in a source;
	// Fields here are also OrderBy
	Fields []querybuildertypes.Field `json:"fields,omitempty"`
}
