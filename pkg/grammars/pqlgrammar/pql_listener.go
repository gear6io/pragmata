// Code generated from PQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package pqlgrammar // PQL
import "github.com/antlr4-go/antlr/v4"

// PQLListener is a complete listener for a parse tree produced by PQL.
type PQLListener interface {
	antlr.ParseTreeListener

	// EnterPipeline is called when entering the pipeline production.
	EnterPipeline(c *PipelineContext)

	// EnterTransform is called when entering the transform production.
	EnterTransform(c *TransformContext)

	// EnterFromTransform is called when entering the fromTransform production.
	EnterFromTransform(c *FromTransformContext)

	// EnterFilterTransform is called when entering the filterTransform production.
	EnterFilterTransform(c *FilterTransformContext)

	// EnterFilterBody is called when entering the filterBody production.
	EnterFilterBody(c *FilterBodyContext)

	// EnterDeriveTransform is called when entering the deriveTransform production.
	EnterDeriveTransform(c *DeriveTransformContext)

	// EnterSelectTransform is called when entering the selectTransform production.
	EnterSelectTransform(c *SelectTransformContext)

	// EnterGroupTransform is called when entering the groupTransform production.
	EnterGroupTransform(c *GroupTransformContext)

	// EnterJoinTransform is called when entering the joinTransform production.
	EnterJoinTransform(c *JoinTransformContext)

	// EnterJoinSide is called when entering the joinSide production.
	EnterJoinSide(c *JoinSideContext)

	// EnterSelfJoinCond is called when entering the SelfJoinCond production.
	EnterSelfJoinCond(c *SelfJoinCondContext)

	// EnterExplicitJoinCond is called when entering the ExplicitJoinCond production.
	EnterExplicitJoinCond(c *ExplicitJoinCondContext)

	// EnterJoinCondExpr is called when entering the joinCondExpr production.
	EnterJoinCondExpr(c *JoinCondExprContext)

	// EnterJoinCondToken is called when entering the joinCondToken production.
	EnterJoinCondToken(c *JoinCondTokenContext)

	// EnterJoinCondInner is called when entering the joinCondInner production.
	EnterJoinCondInner(c *JoinCondInnerContext)

	// EnterArrayJoinTransform is called when entering the arrayJoinTransform production.
	EnterArrayJoinTransform(c *ArrayJoinTransformContext)

	// EnterSortTransform is called when entering the sortTransform production.
	EnterSortTransform(c *SortTransformContext)

	// EnterTakeTransform is called when entering the takeTransform production.
	EnterTakeTransform(c *TakeTransformContext)

	// EnterSkipTransform is called when entering the skipTransform production.
	EnterSkipTransform(c *SkipTransformContext)

	// EnterWindowTransform is called when entering the windowTransform production.
	EnterWindowTransform(c *WindowTransformContext)

	// EnterAssignmentList is called when entering the assignmentList production.
	EnterAssignmentList(c *AssignmentListContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterSelectionList is called when entering the selectionList production.
	EnterSelectionList(c *SelectionListContext)

	// EnterAliasedSelection is called when entering the AliasedSelection production.
	EnterAliasedSelection(c *AliasedSelectionContext)

	// EnterBareSelection is called when entering the BareSelection production.
	EnterBareSelection(c *BareSelectionContext)

	// EnterKeyList is called when entering the keyList production.
	EnterKeyList(c *KeyListContext)

	// EnterComputedKey is called when entering the ComputedKey production.
	EnterComputedKey(c *ComputedKeyContext)

	// EnterColumnKey is called when entering the ColumnKey production.
	EnterColumnKey(c *ColumnKeyContext)

	// EnterSortList is called when entering the sortList production.
	EnterSortList(c *SortListContext)

	// EnterDescSort is called when entering the DescSort production.
	EnterDescSort(c *DescSortContext)

	// EnterAscSortExplicit is called when entering the AscSortExplicit production.
	EnterAscSortExplicit(c *AscSortExplicitContext)

	// EnterAscSort is called when entering the AscSort production.
	EnterAscSort(c *AscSortContext)

	// EnterOpaqueExpr is called when entering the opaqueExpr production.
	EnterOpaqueExpr(c *OpaqueExprContext)

	// EnterOpaqueToken is called when entering the opaqueToken production.
	EnterOpaqueToken(c *OpaqueTokenContext)

	// EnterOpaqueInner is called when entering the opaqueInner production.
	EnterOpaqueInner(c *OpaqueInnerContext)

	// ExitPipeline is called when exiting the pipeline production.
	ExitPipeline(c *PipelineContext)

	// ExitTransform is called when exiting the transform production.
	ExitTransform(c *TransformContext)

	// ExitFromTransform is called when exiting the fromTransform production.
	ExitFromTransform(c *FromTransformContext)

	// ExitFilterTransform is called when exiting the filterTransform production.
	ExitFilterTransform(c *FilterTransformContext)

	// ExitFilterBody is called when exiting the filterBody production.
	ExitFilterBody(c *FilterBodyContext)

	// ExitDeriveTransform is called when exiting the deriveTransform production.
	ExitDeriveTransform(c *DeriveTransformContext)

	// ExitSelectTransform is called when exiting the selectTransform production.
	ExitSelectTransform(c *SelectTransformContext)

	// ExitGroupTransform is called when exiting the groupTransform production.
	ExitGroupTransform(c *GroupTransformContext)

	// ExitJoinTransform is called when exiting the joinTransform production.
	ExitJoinTransform(c *JoinTransformContext)

	// ExitJoinSide is called when exiting the joinSide production.
	ExitJoinSide(c *JoinSideContext)

	// ExitSelfJoinCond is called when exiting the SelfJoinCond production.
	ExitSelfJoinCond(c *SelfJoinCondContext)

	// ExitExplicitJoinCond is called when exiting the ExplicitJoinCond production.
	ExitExplicitJoinCond(c *ExplicitJoinCondContext)

	// ExitJoinCondExpr is called when exiting the joinCondExpr production.
	ExitJoinCondExpr(c *JoinCondExprContext)

	// ExitJoinCondToken is called when exiting the joinCondToken production.
	ExitJoinCondToken(c *JoinCondTokenContext)

	// ExitJoinCondInner is called when exiting the joinCondInner production.
	ExitJoinCondInner(c *JoinCondInnerContext)

	// ExitArrayJoinTransform is called when exiting the arrayJoinTransform production.
	ExitArrayJoinTransform(c *ArrayJoinTransformContext)

	// ExitSortTransform is called when exiting the sortTransform production.
	ExitSortTransform(c *SortTransformContext)

	// ExitTakeTransform is called when exiting the takeTransform production.
	ExitTakeTransform(c *TakeTransformContext)

	// ExitSkipTransform is called when exiting the skipTransform production.
	ExitSkipTransform(c *SkipTransformContext)

	// ExitWindowTransform is called when exiting the windowTransform production.
	ExitWindowTransform(c *WindowTransformContext)

	// ExitAssignmentList is called when exiting the assignmentList production.
	ExitAssignmentList(c *AssignmentListContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitSelectionList is called when exiting the selectionList production.
	ExitSelectionList(c *SelectionListContext)

	// ExitAliasedSelection is called when exiting the AliasedSelection production.
	ExitAliasedSelection(c *AliasedSelectionContext)

	// ExitBareSelection is called when exiting the BareSelection production.
	ExitBareSelection(c *BareSelectionContext)

	// ExitKeyList is called when exiting the keyList production.
	ExitKeyList(c *KeyListContext)

	// ExitComputedKey is called when exiting the ComputedKey production.
	ExitComputedKey(c *ComputedKeyContext)

	// ExitColumnKey is called when exiting the ColumnKey production.
	ExitColumnKey(c *ColumnKeyContext)

	// ExitSortList is called when exiting the sortList production.
	ExitSortList(c *SortListContext)

	// ExitDescSort is called when exiting the DescSort production.
	ExitDescSort(c *DescSortContext)

	// ExitAscSortExplicit is called when exiting the AscSortExplicit production.
	ExitAscSortExplicit(c *AscSortExplicitContext)

	// ExitAscSort is called when exiting the AscSort production.
	ExitAscSort(c *AscSortContext)

	// ExitOpaqueExpr is called when exiting the opaqueExpr production.
	ExitOpaqueExpr(c *OpaqueExprContext)

	// ExitOpaqueToken is called when exiting the opaqueToken production.
	ExitOpaqueToken(c *OpaqueTokenContext)

	// ExitOpaqueInner is called when exiting the opaqueInner production.
	ExitOpaqueInner(c *OpaqueInnerContext)
}
