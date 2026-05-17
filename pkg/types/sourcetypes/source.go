package sourcetypes

import "github.com/gear6io/pragmata/pkg/types/querybuildertypes"

// Source is a table in the pragmata_source ClickHouse database.
type Source struct {
	Name     string                    `json:"name"`
	Database string                    `json:"-"` // internal — always "pragmata_source"
	Engine   string                    `json:"engine,omitempty"`
	Fields   []querybuildertypes.Field `json:"fields,omitempty"`
}
