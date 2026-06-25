// Package prqlvisitor visits PRQL (Pipeline Query Language) source using the ANTLR4-generated grammar
// and compiles it to a ClickHouse SelectBuilder.
// Run scripts/generate-prql-grammar.sh to regenerate grammar/ after editing PRQLLexer.g4 or PRQL.g4.
package prqlvisitor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	prql "github.com/gear6io/pragmata/pkg/grammars/prqlgrammar"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

// PRQLVisitorOpts configures optional validation hooks for the PRQL visitor.
type PRQLVisitorOpts struct {
	SourceValidator Validator
	FieldValidator  Validator
}

// prqlVisitor walks the ANTLR parse tree, accumulating a SelectBuilder.
type prqlVisitor struct {
	sb      *sqlbuilder.SelectBuilder
	inGroup bool
	err     error
	opts    PRQLVisitorOpts
}

// Visit parses pqlSrc as a PRQL pipeline and returns a SelectBuilder pre-populated with the
// compiled ClickHouse query. The caller may further compose the builder before calling Build().
// pqlSrc must be the body of a single pipeline node (no @name: header).
func Visit(pqlSrc string, opts PRQLVisitorOpts) (*sqlbuilder.SelectBuilder, error) {
	if !strings.HasSuffix(pqlSrc, "\n") {
		pqlSrc += "\n"
	}

	el := &errListener{}

	lexer := prql.NewPRQLLexer(antlr.NewInputStream(pqlSrc))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := prql.NewPRQL(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	tree := p.Query()
	if el.msg != "" {
		return nil, fmt.Errorf("prql: %s", el.msg)
	}

	v := &prqlVisitor{sb: sqlbuilder.NewSelectBuilder(), opts: opts}
	v.visitQuery(tree)
	if v.err != nil {
		return nil, v.err
	}
	return v.sb, nil
}

// ── Visitor methods ────────────────────────────────────────────────────────────

func (v *prqlVisitor) visitQuery(ctx prql.IQueryContext) {
	for _, cl := range ctx.AllClause() {
		if v.err != nil {
			return
		}
		v.visitClause(cl)
	}
}

func (v *prqlVisitor) visitClause(ctx prql.IClauseContext) {
	switch {
	case ctx.FromClause() != nil:
		v.visitFromClause(ctx.FromClause())
	case ctx.FilterClause() != nil:
		v.visitFilterClause(ctx.FilterClause())
	case ctx.GroupClause() != nil:
		v.inGroup = true
		v.visitGroupClause(ctx.GroupClause())
	case ctx.DeriveClause() != nil:
		v.visitDeriveClause(ctx.DeriveClause())
	case ctx.SelectClause() != nil:
		v.visitSelectClause(ctx.SelectClause())
	case ctx.JoinClause() != nil:
		v.visitJoinClause(ctx.JoinClause())
	case ctx.ArrayJoinClause() != nil:
		v.visitArrayJoinClause(ctx.ArrayJoinClause())
	case ctx.SortClause() != nil:
		v.visitSortClause(ctx.SortClause())
	case ctx.TakeClause() != nil:
		v.visitTakeClause(ctx.TakeClause())
	case ctx.SkipClause() != nil:
		v.visitSkipClause(ctx.SkipClause())
	case ctx.WindowClause() != nil:
		v.visitWindowClause(ctx.WindowClause())
	}
}

func (v *prqlVisitor) visitFromClause(ctx prql.IFromClauseContext) {
	ident := ctx.IDENT().GetText()
	if v.opts.SourceValidator != nil {
		if err := v.opts.SourceValidator(ident); err != nil {
			v.err = err
			return
		}
	}
	from := ident
	if ctx.KW_FINAL() != nil {
		from += " FINAL"
	}
	v.sb.Select("*")
	v.sb.From(from)
}

func (v *prqlVisitor) visitFilterClause(ctx prql.IFilterClauseContext) {
	cond := filterLines(ctx.FilterBody().AllFILTER_LINE())
	if v.inGroup {
		v.sb.Having(cond)
	} else {
		v.sb.Where(cond)
	}
}

func (v *prqlVisitor) visitDeriveClause(ctx prql.IDeriveClauseContext) {
	for _, a := range ctx.AssignmentList().AllAssignment() {
		v.sb.SelectMore(fmt.Sprintf("%s AS %s", a.OpaqueExpr().GetText(), a.IDENT().GetText()))
	}
}

func (v *prqlVisitor) visitSelectClause(ctx prql.ISelectClauseContext) {
	var cols []string
	for _, item := range ctx.SelectionList().AllSelectionItem() {
		cols = append(cols, selectionItem(item))
	}
	v.sb.Select(cols...)
}

func (v *prqlVisitor) visitGroupClause(ctx prql.IGroupClauseContext) {
	var selExprs, grpExprs, aggExprs []string
	for _, ki := range ctx.KeyList().AllKeyItem() {
		sel, grp := keyItemExprs(ki)
		selExprs = append(selExprs, sel)
		grpExprs = append(grpExprs, grp)
	}
	for _, a := range ctx.AssignmentList().AllAssignment() {
		aggExprs = append(aggExprs, fmt.Sprintf("%s AS %s", a.OpaqueExpr().GetText(), a.IDENT().GetText()))
	}
	v.sb.Select(append(selExprs, aggExprs...)...)
	v.sb.GroupBy(grpExprs...)
}

func (v *prqlVisitor) visitJoinClause(ctx prql.IJoinClauseContext) {
	opt, tableExpr := buildJoinExpr(ctx)
	v.sb.JoinWithOption(opt, tableExpr)
}

func (v *prqlVisitor) visitArrayJoinClause(ctx prql.IArrayJoinClauseContext) {
	idents := ctx.AllIDENT()
	col := idents[0].GetText()
	var expr string
	if len(idents) == 2 {
		expr = fmt.Sprintf("%s AS %s", col, idents[1].GetText())
	} else {
		expr = col
	}
	v.sb.JoinWithOption(sqlbuilder.JoinOption("ARRAY"), expr)
}

func (v *prqlVisitor) visitSortClause(ctx prql.ISortClauseContext) {
	items := make([]string, 0, len(ctx.SortList().AllSortItem()))
	for _, si := range ctx.SortList().AllSortItem() {
		items = append(items, sortItem(si))
	}
	v.sb.OrderBy(items...)
}

func (v *prqlVisitor) visitTakeClause(ctx prql.ITakeClauseContext) {
	if ctx.RANGE() != nil {
		n := mustInt(ctx.INTEGER(0).GetText())
		m := mustInt(ctx.INTEGER(1).GetText())
		v.sb.Limit(m - n)
		v.sb.Offset(n)
	} else {
		v.sb.Limit(mustInt(ctx.INTEGER(0).GetText()))
	}
}

func (v *prqlVisitor) visitSkipClause(ctx prql.ISkipClauseContext) {
	v.sb.Offset(mustInt(ctx.INTEGER().GetText()))
}

func (v *prqlVisitor) visitWindowClause(ctx prql.IWindowClauseContext) {
	for _, a := range ctx.AssignmentList().AllAssignment() {
		v.sb.SelectMore(fmt.Sprintf("%s AS %s", a.OpaqueExpr().GetText(), a.IDENT().GetText()))
	}
}

// ── Helpers ────────────────────────────────────────────────────────────────────

func filterLines(tokens []antlr.TerminalNode) string {
	parts := make([]string, 0, len(tokens))
	for _, tok := range tokens {
		line := strings.TrimSpace(strings.TrimRight(tok.GetText(), "\r\n"))
		if line != "" {
			parts = append(parts, line)
		}
	}
	return strings.Join(parts, " ")
}

func keyItemExprs(item prql.IKeyItemContext) (selectExpr, groupExpr string) {
	switch k := item.(type) {
	case *prql.ComputedKeyContext:
		expr := k.OpaqueExpr().GetText()
		return fmt.Sprintf("%s AS %s", expr, k.IDENT().GetText()), expr
	case *prql.ColumnKeyContext:
		name := k.IDENT().GetText()
		return name, name
	}
	return "", ""
}

func selectionItem(item prql.ISelectionItemContext) string {
	switch s := item.(type) {
	case *prql.AliasedSelectionContext:
		return fmt.Sprintf("%s AS %s", s.OpaqueExpr().GetText(), s.IDENT().GetText())
	case *prql.BareSelectionContext:
		return s.OpaqueExpr().GetText()
	}
	return ""
}

func buildJoinExpr(j prql.IJoinClauseContext) (sqlbuilder.JoinOption, string) {
	var option sqlbuilder.JoinOption
	if side := j.JoinSide(); side != nil {
		switch {
		case side.KW_LEFT() != nil:
			option = sqlbuilder.LeftJoin
		case side.KW_RIGHT() != nil:
			option = sqlbuilder.RightJoin
		case side.KW_INNER() != nil:
			option = sqlbuilder.InnerJoin
		case side.KW_FULL() != nil:
			option = sqlbuilder.FullOuterJoin
		}
	}
	table := j.IDENT().GetText()
	switch jc := j.JoinCond().(type) {
	case *prql.SelfJoinCondContext:
		return option, fmt.Sprintf("%s USING (%s)", table, jc.IDENT().GetText())
	case *prql.ExplicitJoinCondContext:
		return option, fmt.Sprintf("%s ON %s", table, jc.JoinCondExpr().GetText())
	}
	return option, table
}

func sortItem(item prql.ISortItemContext) string {
	switch s := item.(type) {
	case *prql.DescSortContext:
		return s.IDENT().GetText() + " DESC"
	case *prql.AscSortExplicitContext:
		return s.IDENT().GetText() + " ASC"
	case *prql.AscSortContext:
		return s.IDENT().GetText()
	}
	return ""
}

func mustInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// ── Error listener ─────────────────────────────────────────────────────────────

type errListener struct {
	*antlr.DefaultErrorListener
	msg string
}

func (e *errListener) SyntaxError(
	_ antlr.Recognizer, _ interface{},
	line, col int, msg string,
	_ antlr.RecognitionException,
) {
	if e.msg == "" {
		e.msg = fmt.Sprintf("line %d:%d %s", line, col, msg)
	}
}
