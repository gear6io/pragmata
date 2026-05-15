// Package pqlcompiler compiles PQL (Pipeline Query Language) parse trees to
// ClickHouse SQL.  The caller owns grammar parsing; this package only walks
// the already-constructed parse tree and emits SQL text.
//
// # Compilation strategy
//
// PQL is a linear sequence of clauses.  They are folded into a single
// SELECT where possible:
//
//	from        → FROM clause + default SELECT *
//	filter      → WHERE (before group) or HAVING (after group, positional)
//	join        → JOIN clause
//	derive      → extra SELECT columns (appended via SelectMore)
//	group       → GROUP BY + replaces SELECT column list
//	select      → overrides the SELECT column list
//	sort        → ORDER BY
//	take        → LIMIT  (range form also sets OFFSET)
//	skip        → OFFSET
//	window      → window-function SELECT columns (appended via SelectMore)
//	array_join  → ARRAY JOIN clause
package sqlmeshbuilder

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	prql "github.com/gear6io/pragmata/pkg/grammars/prqlgrammar"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

// PrepareSQLMesh parses pqlSrc as a PQL pipeline and returns a SelectBuilder
// pre-populated with the compiled query. The caller may further compose the
// builder (extra WHERE conditions, etc.) before calling Build().
// pqlSrc must be the body of a single pipeline node (no @name: header).
func PrepareSQLMesh(pqlSrc string) (*sqlbuilder.SelectBuilder, error) {
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
		return nil, fmt.Errorf("pql: %s", el.msg)
	}

	c := &compiler{sb: sqlbuilder.NewSelectBuilder()}
	return c.compile(tree)
}

// ── Internal state ─────────────────────────────────────────────────────────

type compiler struct {
	sb *sqlbuilder.SelectBuilder
}

func (c *compiler) compile(tree prql.IQueryContext) (*sqlbuilder.SelectBuilder, error) {
	if err := c.VisitQuery(tree); err != nil {
		return nil, err
	}
	return c.sb, nil
}

// ── Visitor dispatch ───────────────────────────────────────────────────────

// Visit dispatches to the specific visit method based on node type.
func (c *compiler) Visit(tree antlr.ParseTree) error {
	if tree == nil {
		return nil
	}
	switch t := tree.(type) {
	case *prql.QueryContext:
		return c.VisitQuery(t)
	case *prql.ClauseContext:
		inGroup := false
		return c.VisitClause(t, &inGroup)
	}
	return nil
}

func (c *compiler) VisitQuery(ctx prql.IQueryContext) error {
	inGroup := false
	for _, cl := range ctx.AllClause() {
		if err := c.VisitClause(cl, &inGroup); err != nil {
			return err
		}
	}
	return nil
}

func (c *compiler) VisitClause(ctx prql.IClauseContext, inGroup *bool) error {
	switch {
	case ctx.FromClause() != nil:
		return c.VisitFromClause(ctx.FromClause())
	case ctx.FilterClause() != nil:
		return c.VisitFilterClause(ctx.FilterClause(), *inGroup)
	case ctx.GroupClause() != nil:
		*inGroup = true
		return c.VisitGroupClause(ctx.GroupClause())
	case ctx.DeriveClause() != nil:
		return c.VisitDeriveClause(ctx.DeriveClause())
	case ctx.SelectClause() != nil:
		return c.VisitSelectClause(ctx.SelectClause())
	case ctx.JoinClause() != nil:
		return c.VisitJoinClause(ctx.JoinClause())
	case ctx.ArrayJoinClause() != nil:
		return c.VisitArrayJoinClause(ctx.ArrayJoinClause())
	case ctx.SortClause() != nil:
		return c.VisitSortClause(ctx.SortClause())
	case ctx.TakeClause() != nil:
		return c.VisitTakeClause(ctx.TakeClause())
	case ctx.SkipClause() != nil:
		return c.VisitSkipClause(ctx.SkipClause())
	case ctx.WindowClause() != nil:
		return c.VisitWindowClause(ctx.WindowClause())
	}
	return nil
}

// ── Clause visitors ────────────────────────────────────────────────────────

func (c *compiler) VisitFromClause(ctx prql.IFromClauseContext) error {
	from := ctx.IDENT().GetText()
	if ctx.KW_FINAL() != nil {
		from += " FINAL"
	}
	c.sb.Select("*")
	c.sb.From(from)
	return nil
}

func (c *compiler) VisitFilterClause(ctx prql.IFilterClauseContext, inGroup bool) error {
	cond := filterLines(ctx.FilterBody().AllFILTER_LINE())
	if inGroup {
		c.sb.Having(cond)
	} else {
		c.sb.Where(cond)
	}
	return nil
}

func (c *compiler) VisitDeriveClause(ctx prql.IDeriveClauseContext) error {
	for _, a := range ctx.AssignmentList().AllAssignment() {
		c.sb.SelectMore(fmt.Sprintf("%s AS %s", a.OpaqueExpr().GetText(), a.IDENT().GetText()))
	}
	return nil
}

func (c *compiler) VisitSelectClause(ctx prql.ISelectClauseContext) error {
	var cols []string
	for _, item := range ctx.SelectionList().AllSelectionItem() {
		cols = append(cols, selectionItem(item))
	}
	c.sb.Select(cols...)
	return nil
}

func (c *compiler) VisitGroupClause(ctx prql.IGroupClauseContext) error {
	var selExprs, grpExprs, aggExprs []string
	for _, ki := range ctx.KeyList().AllKeyItem() {
		sel, grp := keyItemExprs(ki)
		selExprs = append(selExprs, sel)
		grpExprs = append(grpExprs, grp)
	}
	for _, a := range ctx.AssignmentList().AllAssignment() {
		aggExprs = append(aggExprs, fmt.Sprintf("%s AS %s", a.OpaqueExpr().GetText(), a.IDENT().GetText()))
	}
	c.sb.Select(append(selExprs, aggExprs...)...)
	c.sb.GroupBy(grpExprs...)
	return nil
}

func (c *compiler) VisitJoinClause(ctx prql.IJoinClauseContext) error {
	opt, tableExpr := buildJoinExpr(ctx)
	c.sb.JoinWithOption(opt, tableExpr)
	return nil
}

func (c *compiler) VisitArrayJoinClause(ctx prql.IArrayJoinClauseContext) error {
	idents := ctx.AllIDENT()
	col := idents[0].GetText()
	var expr string
	if len(idents) == 2 {
		expr = fmt.Sprintf("%s AS %s", col, idents[1].GetText())
	} else {
		expr = col
	}
	c.sb.JoinWithOption(sqlbuilder.JoinOption("ARRAY"), expr)
	return nil
}

func (c *compiler) VisitSortClause(ctx prql.ISortClauseContext) error {
	items := make([]string, 0, len(ctx.SortList().AllSortItem()))
	for _, si := range ctx.SortList().AllSortItem() {
		items = append(items, sortItem(si))
	}
	c.sb.OrderBy(items...)
	return nil
}

func (c *compiler) VisitTakeClause(ctx prql.ITakeClauseContext) error {
	if ctx.RANGE() != nil {
		n := mustInt(ctx.INTEGER(0).GetText())
		m := mustInt(ctx.INTEGER(1).GetText())
		c.sb.Limit(m - n)
		c.sb.Offset(n)
	} else {
		c.sb.Limit(mustInt(ctx.INTEGER(0).GetText()))
	}
	return nil
}

func (c *compiler) VisitSkipClause(ctx prql.ISkipClauseContext) error {
	c.sb.Offset(mustInt(ctx.INTEGER().GetText()))
	return nil
}

func (c *compiler) VisitWindowClause(ctx prql.IWindowClauseContext) error {
	for _, a := range ctx.AssignmentList().AllAssignment() {
		c.sb.SelectMore(fmt.Sprintf("%s AS %s", a.OpaqueExpr().GetText(), a.IDENT().GetText()))
	}
	return nil
}

// ── Helpers ────────────────────────────────────────────────────────────────

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

// buildJoinExpr returns the JoinOption and table-expression (table + USING/ON condition)
// for the given join clause.
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

// ── Error listener ─────────────────────────────────────────────────────────

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
