package types

import (
	"time"

	"github.com/gear6io/pragmata/pkg/valuer"
)

type TimeAuditable struct {
	CreatedAt time.Time `bun:"created_at" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at" json:"updatedAt"`
}

type UserAuditable struct {
	CreatedBy valuer.Email `bun:"created_by,type:text" json:"createdBy"`
	UpdatedBy valuer.Email `bun:"updated_by,type:text" json:"updatedBy"`
}
