package implpipes

import (
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/sqlstore"
)

type module struct {
	sqlStore sqlstore.SQLStore
}

func NewModule(sqlStore sqlstore.SQLStore) pipes.Module {
	return &module{sqlStore: sqlStore}
}
