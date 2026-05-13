// Code generated from PQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package pqlgrammar // PQL
import "github.com/antlr4-go/antlr/v4"

// BasePQLListener is a complete listener for a parse tree produced by PQL.
type BasePQLListener struct{}

var _ PQLListener = &BasePQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPipeline is called when production pipeline is entered.
func (s *BasePQLListener) EnterPipeline(ctx *PipelineContext) {}

// ExitPipeline is called when production pipeline is exited.
func (s *BasePQLListener) ExitPipeline(ctx *PipelineContext) {}

// EnterTransform is called when production transform is entered.
func (s *BasePQLListener) EnterTransform(ctx *TransformContext) {}

// ExitTransform is called when production transform is exited.
func (s *BasePQLListener) ExitTransform(ctx *TransformContext) {}

// EnterFromTransform is called when production fromTransform is entered.
func (s *BasePQLListener) EnterFromTransform(ctx *FromTransformContext) {}

// ExitFromTransform is called when production fromTransform is exited.
func (s *BasePQLListener) ExitFromTransform(ctx *FromTransformContext) {}

// EnterFilterTransform is called when production filterTransform is entered.
func (s *BasePQLListener) EnterFilterTransform(ctx *FilterTransformContext) {}

// ExitFilterTransform is called when production filterTransform is exited.
func (s *BasePQLListener) ExitFilterTransform(ctx *FilterTransformContext) {}

// EnterFilterBody is called when production filterBody is entered.
func (s *BasePQLListener) EnterFilterBody(ctx *FilterBodyContext) {}

// ExitFilterBody is called when production filterBody is exited.
func (s *BasePQLListener) ExitFilterBody(ctx *FilterBodyContext) {}

// EnterDeriveTransform is called when production deriveTransform is entered.
func (s *BasePQLListener) EnterDeriveTransform(ctx *DeriveTransformContext) {}

// ExitDeriveTransform is called when production deriveTransform is exited.
func (s *BasePQLListener) ExitDeriveTransform(ctx *DeriveTransformContext) {}

// EnterSelectTransform is called when production selectTransform is entered.
func (s *BasePQLListener) EnterSelectTransform(ctx *SelectTransformContext) {}

// ExitSelectTransform is called when production selectTransform is exited.
func (s *BasePQLListener) ExitSelectTransform(ctx *SelectTransformContext) {}

// EnterGroupTransform is called when production groupTransform is entered.
func (s *BasePQLListener) EnterGroupTransform(ctx *GroupTransformContext) {}

// ExitGroupTransform is called when production groupTransform is exited.
func (s *BasePQLListener) ExitGroupTransform(ctx *GroupTransformContext) {}

// EnterJoinTransform is called when production joinTransform is entered.
func (s *BasePQLListener) EnterJoinTransform(ctx *JoinTransformContext) {}

// ExitJoinTransform is called when production joinTransform is exited.
func (s *BasePQLListener) ExitJoinTransform(ctx *JoinTransformContext) {}

// EnterJoinSide is called when production joinSide is entered.
func (s *BasePQLListener) EnterJoinSide(ctx *JoinSideContext) {}

// ExitJoinSide is called when production joinSide is exited.
func (s *BasePQLListener) ExitJoinSide(ctx *JoinSideContext) {}

// EnterSelfJoinCond is called when production SelfJoinCond is entered.
func (s *BasePQLListener) EnterSelfJoinCond(ctx *SelfJoinCondContext) {}

// ExitSelfJoinCond is called when production SelfJoinCond is exited.
func (s *BasePQLListener) ExitSelfJoinCond(ctx *SelfJoinCondContext) {}

// EnterExplicitJoinCond is called when production ExplicitJoinCond is entered.
func (s *BasePQLListener) EnterExplicitJoinCond(ctx *ExplicitJoinCondContext) {}

// ExitExplicitJoinCond is called when production ExplicitJoinCond is exited.
func (s *BasePQLListener) ExitExplicitJoinCond(ctx *ExplicitJoinCondContext) {}

// EnterJoinCondExpr is called when production joinCondExpr is entered.
func (s *BasePQLListener) EnterJoinCondExpr(ctx *JoinCondExprContext) {}

// ExitJoinCondExpr is called when production joinCondExpr is exited.
func (s *BasePQLListener) ExitJoinCondExpr(ctx *JoinCondExprContext) {}

// EnterJoinCondToken is called when production joinCondToken is entered.
func (s *BasePQLListener) EnterJoinCondToken(ctx *JoinCondTokenContext) {}

// ExitJoinCondToken is called when production joinCondToken is exited.
func (s *BasePQLListener) ExitJoinCondToken(ctx *JoinCondTokenContext) {}

// EnterJoinCondInner is called when production joinCondInner is entered.
func (s *BasePQLListener) EnterJoinCondInner(ctx *JoinCondInnerContext) {}

// ExitJoinCondInner is called when production joinCondInner is exited.
func (s *BasePQLListener) ExitJoinCondInner(ctx *JoinCondInnerContext) {}

// EnterArrayJoinTransform is called when production arrayJoinTransform is entered.
func (s *BasePQLListener) EnterArrayJoinTransform(ctx *ArrayJoinTransformContext) {}

// ExitArrayJoinTransform is called when production arrayJoinTransform is exited.
func (s *BasePQLListener) ExitArrayJoinTransform(ctx *ArrayJoinTransformContext) {}

// EnterSortTransform is called when production sortTransform is entered.
func (s *BasePQLListener) EnterSortTransform(ctx *SortTransformContext) {}

// ExitSortTransform is called when production sortTransform is exited.
func (s *BasePQLListener) ExitSortTransform(ctx *SortTransformContext) {}

// EnterTakeTransform is called when production takeTransform is entered.
func (s *BasePQLListener) EnterTakeTransform(ctx *TakeTransformContext) {}

// ExitTakeTransform is called when production takeTransform is exited.
func (s *BasePQLListener) ExitTakeTransform(ctx *TakeTransformContext) {}

// EnterSkipTransform is called when production skipTransform is entered.
func (s *BasePQLListener) EnterSkipTransform(ctx *SkipTransformContext) {}

// ExitSkipTransform is called when production skipTransform is exited.
func (s *BasePQLListener) ExitSkipTransform(ctx *SkipTransformContext) {}

// EnterWindowTransform is called when production windowTransform is entered.
func (s *BasePQLListener) EnterWindowTransform(ctx *WindowTransformContext) {}

// ExitWindowTransform is called when production windowTransform is exited.
func (s *BasePQLListener) ExitWindowTransform(ctx *WindowTransformContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BasePQLListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BasePQLListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasePQLListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasePQLListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSelectionList is called when production selectionList is entered.
func (s *BasePQLListener) EnterSelectionList(ctx *SelectionListContext) {}

// ExitSelectionList is called when production selectionList is exited.
func (s *BasePQLListener) ExitSelectionList(ctx *SelectionListContext) {}

// EnterAliasedSelection is called when production AliasedSelection is entered.
func (s *BasePQLListener) EnterAliasedSelection(ctx *AliasedSelectionContext) {}

// ExitAliasedSelection is called when production AliasedSelection is exited.
func (s *BasePQLListener) ExitAliasedSelection(ctx *AliasedSelectionContext) {}

// EnterBareSelection is called when production BareSelection is entered.
func (s *BasePQLListener) EnterBareSelection(ctx *BareSelectionContext) {}

// ExitBareSelection is called when production BareSelection is exited.
func (s *BasePQLListener) ExitBareSelection(ctx *BareSelectionContext) {}

// EnterKeyList is called when production keyList is entered.
func (s *BasePQLListener) EnterKeyList(ctx *KeyListContext) {}

// ExitKeyList is called when production keyList is exited.
func (s *BasePQLListener) ExitKeyList(ctx *KeyListContext) {}

// EnterComputedKey is called when production ComputedKey is entered.
func (s *BasePQLListener) EnterComputedKey(ctx *ComputedKeyContext) {}

// ExitComputedKey is called when production ComputedKey is exited.
func (s *BasePQLListener) ExitComputedKey(ctx *ComputedKeyContext) {}

// EnterColumnKey is called when production ColumnKey is entered.
func (s *BasePQLListener) EnterColumnKey(ctx *ColumnKeyContext) {}

// ExitColumnKey is called when production ColumnKey is exited.
func (s *BasePQLListener) ExitColumnKey(ctx *ColumnKeyContext) {}

// EnterSortList is called when production sortList is entered.
func (s *BasePQLListener) EnterSortList(ctx *SortListContext) {}

// ExitSortList is called when production sortList is exited.
func (s *BasePQLListener) ExitSortList(ctx *SortListContext) {}

// EnterDescSort is called when production DescSort is entered.
func (s *BasePQLListener) EnterDescSort(ctx *DescSortContext) {}

// ExitDescSort is called when production DescSort is exited.
func (s *BasePQLListener) ExitDescSort(ctx *DescSortContext) {}

// EnterAscSortExplicit is called when production AscSortExplicit is entered.
func (s *BasePQLListener) EnterAscSortExplicit(ctx *AscSortExplicitContext) {}

// ExitAscSortExplicit is called when production AscSortExplicit is exited.
func (s *BasePQLListener) ExitAscSortExplicit(ctx *AscSortExplicitContext) {}

// EnterAscSort is called when production AscSort is entered.
func (s *BasePQLListener) EnterAscSort(ctx *AscSortContext) {}

// ExitAscSort is called when production AscSort is exited.
func (s *BasePQLListener) ExitAscSort(ctx *AscSortContext) {}

// EnterOpaqueExpr is called when production opaqueExpr is entered.
func (s *BasePQLListener) EnterOpaqueExpr(ctx *OpaqueExprContext) {}

// ExitOpaqueExpr is called when production opaqueExpr is exited.
func (s *BasePQLListener) ExitOpaqueExpr(ctx *OpaqueExprContext) {}

// EnterOpaqueToken is called when production opaqueToken is entered.
func (s *BasePQLListener) EnterOpaqueToken(ctx *OpaqueTokenContext) {}

// ExitOpaqueToken is called when production opaqueToken is exited.
func (s *BasePQLListener) ExitOpaqueToken(ctx *OpaqueTokenContext) {}

// EnterOpaqueInner is called when production opaqueInner is entered.
func (s *BasePQLListener) EnterOpaqueInner(ctx *OpaqueInnerContext) {}

// ExitOpaqueInner is called when production opaqueInner is exited.
func (s *BasePQLListener) ExitOpaqueInner(ctx *OpaqueInnerContext) {}
