package datasourcetypes

import "github.com/gear6io/pragmata/pkg/valuer"

type DataSource struct {
	ID       valuer.UUID   `json:"id"`
	Database valuer.String `json:"database"`
}

type DataSourceSchema struct {
	ID           valuer.UUID   `json:"id"`
	DataSourceID valuer.UUID   `json:"dataSourceId"`
	Schema       valuer.String `json:"schema"`
}

type Schema struct {
	
}