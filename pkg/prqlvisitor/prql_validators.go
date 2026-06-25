package prqlvisitor

import (
	"fmt"
	"slices"

	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
)

// Validator validates a single identifier (source table or field name).
type Validator func(name string) error

type SourceValidator func(src string, sources []sourcetypes.Source) error

func NewSourceValidator(src string, availableSources []sourcetypes.Source) error {
	has := slices.ContainsFunc(availableSources, func(source sourcetypes.Source) bool {
		return source.Name == src
	})
	if !has {
		return fmt.Errorf("source[%s] not found", src)
	}

	return nil
}
