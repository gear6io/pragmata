package prqlvisitor

import (
	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/huandu/go-sqlbuilder"
)

const SourceDatabase = "pragmata_source"

// BuildSQL compiles PRQL nodes and assembles a CTE chain with source aliases.
// Sources are emitted first so node CTEs can reference them.
// All CTEs are registered in a single With() call — repeated With() calls replace each other.
func BuildSQL(nodes pipetypes.Nodes, sources pipetypes.Sources) (string, error) {
	if len(nodes) == 0 {
		return "", errors.NewInvalidInputf(CodeInvalidPipeContent, "no nodes provided")
	}
	sbs := make([]*sqlbuilder.SelectBuilder, len(nodes))
	for i, node := range nodes {
		sb, err := Visit(node.SQL, PRQLVisitorOpts{})
		if err != nil {
			return "", errors.WithAdditionalf(err, "node %q", node.Name)
		}
		sbs[i] = sb
	}
	last := len(nodes) - 1
	root := sbs[last]
	var ctes []*sqlbuilder.CTEQueryBuilder
	for _, src := range sources {
		srcSB := sqlbuilder.NewSelectBuilder().Select("*").From(SourceDatabase + "." + src.Table)
		ctes = append(ctes, sqlbuilder.CTEQuery(src.Alias).As(srcSB))
	}
	for i, node := range nodes[:last] {
		ctes = append(ctes, sqlbuilder.CTEQuery(node.Name).As(sbs[i]))
	}
	if len(ctes) > 0 {
		root.With(sqlbuilder.With(ctes...))
	}
	sql, _ := root.BuildWithFlavor(sqlbuilder.ClickHouse)
	return sql, nil
}
