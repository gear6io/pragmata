// Code generated from PQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package pqlgrammar // PQL
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PQL.
type PQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PQL#pipeline.
	VisitPipeline(ctx *PipelineContext) interface{}

	// Visit a parse tree produced by PQL#transform.
	VisitTransform(ctx *TransformContext) interface{}

	// Visit a parse tree produced by PQL#fromTransform.
	VisitFromTransform(ctx *FromTransformContext) interface{}

	// Visit a parse tree produced by PQL#filterTransform.
	VisitFilterTransform(ctx *FilterTransformContext) interface{}

	// Visit a parse tree produced by PQL#filterBody.
	VisitFilterBody(ctx *FilterBodyContext) interface{}

	// Visit a parse tree produced by PQL#deriveTransform.
	VisitDeriveTransform(ctx *DeriveTransformContext) interface{}

	// Visit a parse tree produced by PQL#selectTransform.
	VisitSelectTransform(ctx *SelectTransformContext) interface{}

	// Visit a parse tree produced by PQL#groupTransform.
	VisitGroupTransform(ctx *GroupTransformContext) interface{}

	// Visit a parse tree produced by PQL#joinTransform.
	VisitJoinTransform(ctx *JoinTransformContext) interface{}

	// Visit a parse tree produced by PQL#joinSide.
	VisitJoinSide(ctx *JoinSideContext) interface{}

	// Visit a parse tree produced by PQL#SelfJoinCond.
	VisitSelfJoinCond(ctx *SelfJoinCondContext) interface{}

	// Visit a parse tree produced by PQL#ExplicitJoinCond.
	VisitExplicitJoinCond(ctx *ExplicitJoinCondContext) interface{}

	// Visit a parse tree produced by PQL#joinCondExpr.
	VisitJoinCondExpr(ctx *JoinCondExprContext) interface{}

	// Visit a parse tree produced by PQL#joinCondToken.
	VisitJoinCondToken(ctx *JoinCondTokenContext) interface{}

	// Visit a parse tree produced by PQL#joinCondInner.
	VisitJoinCondInner(ctx *JoinCondInnerContext) interface{}

	// Visit a parse tree produced by PQL#arrayJoinTransform.
	VisitArrayJoinTransform(ctx *ArrayJoinTransformContext) interface{}

	// Visit a parse tree produced by PQL#sortTransform.
	VisitSortTransform(ctx *SortTransformContext) interface{}

	// Visit a parse tree produced by PQL#takeTransform.
	VisitTakeTransform(ctx *TakeTransformContext) interface{}

	// Visit a parse tree produced by PQL#skipTransform.
	VisitSkipTransform(ctx *SkipTransformContext) interface{}

	// Visit a parse tree produced by PQL#windowTransform.
	VisitWindowTransform(ctx *WindowTransformContext) interface{}

	// Visit a parse tree produced by PQL#assignmentList.
	VisitAssignmentList(ctx *AssignmentListContext) interface{}

	// Visit a parse tree produced by PQL#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by PQL#selectionList.
	VisitSelectionList(ctx *SelectionListContext) interface{}

	// Visit a parse tree produced by PQL#AliasedSelection.
	VisitAliasedSelection(ctx *AliasedSelectionContext) interface{}

	// Visit a parse tree produced by PQL#BareSelection.
	VisitBareSelection(ctx *BareSelectionContext) interface{}

	// Visit a parse tree produced by PQL#keyList.
	VisitKeyList(ctx *KeyListContext) interface{}

	// Visit a parse tree produced by PQL#ComputedKey.
	VisitComputedKey(ctx *ComputedKeyContext) interface{}

	// Visit a parse tree produced by PQL#ColumnKey.
	VisitColumnKey(ctx *ColumnKeyContext) interface{}

	// Visit a parse tree produced by PQL#sortList.
	VisitSortList(ctx *SortListContext) interface{}

	// Visit a parse tree produced by PQL#DescSort.
	VisitDescSort(ctx *DescSortContext) interface{}

	// Visit a parse tree produced by PQL#AscSortExplicit.
	VisitAscSortExplicit(ctx *AscSortExplicitContext) interface{}

	// Visit a parse tree produced by PQL#AscSort.
	VisitAscSort(ctx *AscSortContext) interface{}

	// Visit a parse tree produced by PQL#opaqueExpr.
	VisitOpaqueExpr(ctx *OpaqueExprContext) interface{}

	// Visit a parse tree produced by PQL#opaqueToken.
	VisitOpaqueToken(ctx *OpaqueTokenContext) interface{}

	// Visit a parse tree produced by PQL#opaqueInner.
	VisitOpaqueInner(ctx *OpaqueInnerContext) interface{}
}
