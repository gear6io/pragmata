// Code generated from PRQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package prqlgrammar // PRQL
import "github.com/antlr4-go/antlr/v4"

type BasePRQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePRQLVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitClause(ctx *ClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitFilterClause(ctx *FilterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitFilterBody(ctx *FilterBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitDeriveClause(ctx *DeriveClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitGroupClause(ctx *GroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitJoinClause(ctx *JoinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitJoinSide(ctx *JoinSideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitSelfJoinCond(ctx *SelfJoinCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitExplicitJoinCond(ctx *ExplicitJoinCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitJoinCondExpr(ctx *JoinCondExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitJoinCondToken(ctx *JoinCondTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitJoinCondInner(ctx *JoinCondInnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitArrayJoinClause(ctx *ArrayJoinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitSortClause(ctx *SortClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitTakeClause(ctx *TakeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitSkipClause(ctx *SkipClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitAssignmentList(ctx *AssignmentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitSelectionList(ctx *SelectionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitAliasedSelection(ctx *AliasedSelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitBareSelection(ctx *BareSelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitKeyList(ctx *KeyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitComputedKey(ctx *ComputedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitColumnKey(ctx *ColumnKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitSortList(ctx *SortListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitDescSort(ctx *DescSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitAscSortExplicit(ctx *AscSortExplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitAscSort(ctx *AscSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitOpaqueExpr(ctx *OpaqueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitOpaqueToken(ctx *OpaqueTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePRQLVisitor) VisitOpaqueInner(ctx *OpaqueInnerContext) interface{} {
	return v.VisitChildren(ctx)
}
