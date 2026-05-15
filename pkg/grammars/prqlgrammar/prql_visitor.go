// Code generated from PRQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package prqlgrammar // PRQL
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PRQL.
type PRQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PRQL#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by PRQL#clause.
	VisitClause(ctx *ClauseContext) interface{}

	// Visit a parse tree produced by PRQL#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by PRQL#filterClause.
	VisitFilterClause(ctx *FilterClauseContext) interface{}

	// Visit a parse tree produced by PRQL#filterBody.
	VisitFilterBody(ctx *FilterBodyContext) interface{}

	// Visit a parse tree produced by PRQL#deriveClause.
	VisitDeriveClause(ctx *DeriveClauseContext) interface{}

	// Visit a parse tree produced by PRQL#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by PRQL#groupClause.
	VisitGroupClause(ctx *GroupClauseContext) interface{}

	// Visit a parse tree produced by PRQL#joinClause.
	VisitJoinClause(ctx *JoinClauseContext) interface{}

	// Visit a parse tree produced by PRQL#joinSide.
	VisitJoinSide(ctx *JoinSideContext) interface{}

	// Visit a parse tree produced by PRQL#SelfJoinCond.
	VisitSelfJoinCond(ctx *SelfJoinCondContext) interface{}

	// Visit a parse tree produced by PRQL#ExplicitJoinCond.
	VisitExplicitJoinCond(ctx *ExplicitJoinCondContext) interface{}

	// Visit a parse tree produced by PRQL#joinCondExpr.
	VisitJoinCondExpr(ctx *JoinCondExprContext) interface{}

	// Visit a parse tree produced by PRQL#joinCondToken.
	VisitJoinCondToken(ctx *JoinCondTokenContext) interface{}

	// Visit a parse tree produced by PRQL#joinCondInner.
	VisitJoinCondInner(ctx *JoinCondInnerContext) interface{}

	// Visit a parse tree produced by PRQL#arrayJoinClause.
	VisitArrayJoinClause(ctx *ArrayJoinClauseContext) interface{}

	// Visit a parse tree produced by PRQL#sortClause.
	VisitSortClause(ctx *SortClauseContext) interface{}

	// Visit a parse tree produced by PRQL#takeClause.
	VisitTakeClause(ctx *TakeClauseContext) interface{}

	// Visit a parse tree produced by PRQL#skipClause.
	VisitSkipClause(ctx *SkipClauseContext) interface{}

	// Visit a parse tree produced by PRQL#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by PRQL#assignmentList.
	VisitAssignmentList(ctx *AssignmentListContext) interface{}

	// Visit a parse tree produced by PRQL#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by PRQL#selectionList.
	VisitSelectionList(ctx *SelectionListContext) interface{}

	// Visit a parse tree produced by PRQL#AliasedSelection.
	VisitAliasedSelection(ctx *AliasedSelectionContext) interface{}

	// Visit a parse tree produced by PRQL#BareSelection.
	VisitBareSelection(ctx *BareSelectionContext) interface{}

	// Visit a parse tree produced by PRQL#keyList.
	VisitKeyList(ctx *KeyListContext) interface{}

	// Visit a parse tree produced by PRQL#ComputedKey.
	VisitComputedKey(ctx *ComputedKeyContext) interface{}

	// Visit a parse tree produced by PRQL#ColumnKey.
	VisitColumnKey(ctx *ColumnKeyContext) interface{}

	// Visit a parse tree produced by PRQL#sortList.
	VisitSortList(ctx *SortListContext) interface{}

	// Visit a parse tree produced by PRQL#DescSort.
	VisitDescSort(ctx *DescSortContext) interface{}

	// Visit a parse tree produced by PRQL#AscSortExplicit.
	VisitAscSortExplicit(ctx *AscSortExplicitContext) interface{}

	// Visit a parse tree produced by PRQL#AscSort.
	VisitAscSort(ctx *AscSortContext) interface{}

	// Visit a parse tree produced by PRQL#opaqueExpr.
	VisitOpaqueExpr(ctx *OpaqueExprContext) interface{}

	// Visit a parse tree produced by PRQL#opaqueToken.
	VisitOpaqueToken(ctx *OpaqueTokenContext) interface{}

	// Visit a parse tree produced by PRQL#opaqueInner.
	VisitOpaqueInner(ctx *OpaqueInnerContext) interface{}
}
