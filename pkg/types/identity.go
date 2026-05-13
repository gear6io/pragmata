package types

import (
	"github.com/gear6io/pragmata/pkg/valuer"
)

type Identifiable struct {
	ID valuer.UUID `json:"id" bun:"id,pk,type:text" required:"true"`
}
