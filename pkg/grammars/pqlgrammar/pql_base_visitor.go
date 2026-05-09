// Code generated from PQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package pqlgrammar // PQL
import "github.com/antlr4-go/antlr/v4"

type BasePQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePQLVisitor) VisitPipeline(ctx *PipelineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitTransform(ctx *TransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitFromTransform(ctx *FromTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitFilterTransform(ctx *FilterTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitFilterBody(ctx *FilterBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitDeriveTransform(ctx *DeriveTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitSelectTransform(ctx *SelectTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitGroupTransform(ctx *GroupTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitJoinTransform(ctx *JoinTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitJoinSide(ctx *JoinSideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitSelfJoinCond(ctx *SelfJoinCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitExplicitJoinCond(ctx *ExplicitJoinCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitJoinCondExpr(ctx *JoinCondExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitJoinCondToken(ctx *JoinCondTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitJoinCondInner(ctx *JoinCondInnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitArrayJoinTransform(ctx *ArrayJoinTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitSortTransform(ctx *SortTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitTakeTransform(ctx *TakeTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitSkipTransform(ctx *SkipTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitWindowTransform(ctx *WindowTransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitAssignmentList(ctx *AssignmentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitSelectionList(ctx *SelectionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitAliasedSelection(ctx *AliasedSelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitBareSelection(ctx *BareSelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitKeyList(ctx *KeyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitComputedKey(ctx *ComputedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitColumnKey(ctx *ColumnKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitSortList(ctx *SortListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitDescSort(ctx *DescSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitAscSortExplicit(ctx *AscSortExplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitAscSort(ctx *AscSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitOpaqueExpr(ctx *OpaqueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitOpaqueToken(ctx *OpaqueTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePQLVisitor) VisitOpaqueInner(ctx *OpaqueInnerContext) interface{} {
	return v.VisitChildren(ctx)
}
