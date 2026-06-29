package prqlvisitor

import (
	"slices"

	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
)

var CodeSourceNotFound = errors.MustNewCode("source_not_found")

// Validator validates a single identifier (source table or field name).
type Validator func(name string) error

type SourceValidator func(src string, sources []sourcetypes.Source) error

// FROMValidator is a PRQL validator, isNode means either source or referring another node
type FromValidator func(table string, isNode bool) error

func NewSourceValidator(src string, availableSources []sourcetypes.Source) error {
	has := slices.ContainsFunc(availableSources, func(source sourcetypes.Source) bool {
		return source.Name == src
	})
	if !has {
		return errors.NewNotFoundf(CodeSourceNotFound, "source[%s] not found", src)
	}

	return nil
}
