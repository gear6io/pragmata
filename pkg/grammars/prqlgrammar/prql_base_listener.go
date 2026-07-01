// Code generated from PRQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package prqlgrammar // PRQL
import "github.com/antlr4-go/antlr/v4"

// BasePRQLListener is a complete listener for a parse tree produced by PRQL.
type BasePRQLListener struct{}

var _ PRQLListener = &BasePRQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePRQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePRQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePRQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePRQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BasePRQLListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BasePRQLListener) ExitQuery(ctx *QueryContext) {}

// EnterClause is called when production clause is entered.
func (s *BasePRQLListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BasePRQLListener) ExitClause(ctx *ClauseContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BasePRQLListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BasePRQLListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterFilterClause is called when production filterClause is entered.
func (s *BasePRQLListener) EnterFilterClause(ctx *FilterClauseContext) {}

// ExitFilterClause is called when production filterClause is exited.
func (s *BasePRQLListener) ExitFilterClause(ctx *FilterClauseContext) {}

// EnterFilterBody is called when production filterBody is entered.
func (s *BasePRQLListener) EnterFilterBody(ctx *FilterBodyContext) {}

// ExitFilterBody is called when production filterBody is exited.
func (s *BasePRQLListener) ExitFilterBody(ctx *FilterBodyContext) {}

// EnterFilterToken is called when production filterToken is entered.
func (s *BasePRQLListener) EnterFilterToken(ctx *FilterTokenContext) {}

// ExitFilterToken is called when production filterToken is exited.
func (s *BasePRQLListener) ExitFilterToken(ctx *FilterTokenContext) {}

// EnterFilterInner is called when production filterInner is entered.
func (s *BasePRQLListener) EnterFilterInner(ctx *FilterInnerContext) {}

// ExitFilterInner is called when production filterInner is exited.
func (s *BasePRQLListener) ExitFilterInner(ctx *FilterInnerContext) {}

// EnterDeriveClause is called when production deriveClause is entered.
func (s *BasePRQLListener) EnterDeriveClause(ctx *DeriveClauseContext) {}

// ExitDeriveClause is called when production deriveClause is exited.
func (s *BasePRQLListener) ExitDeriveClause(ctx *DeriveClauseContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BasePRQLListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BasePRQLListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterGroupClause is called when production groupClause is entered.
func (s *BasePRQLListener) EnterGroupClause(ctx *GroupClauseContext) {}

// ExitGroupClause is called when production groupClause is exited.
func (s *BasePRQLListener) ExitGroupClause(ctx *GroupClauseContext) {}

// EnterJoinClause is called when production joinClause is entered.
func (s *BasePRQLListener) EnterJoinClause(ctx *JoinClauseContext) {}

// ExitJoinClause is called when production joinClause is exited.
func (s *BasePRQLListener) ExitJoinClause(ctx *JoinClauseContext) {}

// EnterJoinSide is called when production joinSide is entered.
func (s *BasePRQLListener) EnterJoinSide(ctx *JoinSideContext) {}

// ExitJoinSide is called when production joinSide is exited.
func (s *BasePRQLListener) ExitJoinSide(ctx *JoinSideContext) {}

// EnterSelfJoinCond is called when production SelfJoinCond is entered.
func (s *BasePRQLListener) EnterSelfJoinCond(ctx *SelfJoinCondContext) {}

// ExitSelfJoinCond is called when production SelfJoinCond is exited.
func (s *BasePRQLListener) ExitSelfJoinCond(ctx *SelfJoinCondContext) {}

// EnterExplicitJoinCond is called when production ExplicitJoinCond is entered.
func (s *BasePRQLListener) EnterExplicitJoinCond(ctx *ExplicitJoinCondContext) {}

// ExitExplicitJoinCond is called when production ExplicitJoinCond is exited.
func (s *BasePRQLListener) ExitExplicitJoinCond(ctx *ExplicitJoinCondContext) {}

// EnterJoinCondExpr is called when production joinCondExpr is entered.
func (s *BasePRQLListener) EnterJoinCondExpr(ctx *JoinCondExprContext) {}

// ExitJoinCondExpr is called when production joinCondExpr is exited.
func (s *BasePRQLListener) ExitJoinCondExpr(ctx *JoinCondExprContext) {}

// EnterJoinCondToken is called when production joinCondToken is entered.
func (s *BasePRQLListener) EnterJoinCondToken(ctx *JoinCondTokenContext) {}

// ExitJoinCondToken is called when production joinCondToken is exited.
func (s *BasePRQLListener) ExitJoinCondToken(ctx *JoinCondTokenContext) {}

// EnterJoinCondInner is called when production joinCondInner is entered.
func (s *BasePRQLListener) EnterJoinCondInner(ctx *JoinCondInnerContext) {}

// ExitJoinCondInner is called when production joinCondInner is exited.
func (s *BasePRQLListener) ExitJoinCondInner(ctx *JoinCondInnerContext) {}

// EnterArrayJoinClause is called when production arrayJoinClause is entered.
func (s *BasePRQLListener) EnterArrayJoinClause(ctx *ArrayJoinClauseContext) {}

// ExitArrayJoinClause is called when production arrayJoinClause is exited.
func (s *BasePRQLListener) ExitArrayJoinClause(ctx *ArrayJoinClauseContext) {}

// EnterSortClause is called when production sortClause is entered.
func (s *BasePRQLListener) EnterSortClause(ctx *SortClauseContext) {}

// ExitSortClause is called when production sortClause is exited.
func (s *BasePRQLListener) ExitSortClause(ctx *SortClauseContext) {}

// EnterTakeClause is called when production takeClause is entered.
func (s *BasePRQLListener) EnterTakeClause(ctx *TakeClauseContext) {}

// ExitTakeClause is called when production takeClause is exited.
func (s *BasePRQLListener) ExitTakeClause(ctx *TakeClauseContext) {}

// EnterSkipClause is called when production skipClause is entered.
func (s *BasePRQLListener) EnterSkipClause(ctx *SkipClauseContext) {}

// ExitSkipClause is called when production skipClause is exited.
func (s *BasePRQLListener) ExitSkipClause(ctx *SkipClauseContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BasePRQLListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BasePRQLListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BasePRQLListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BasePRQLListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasePRQLListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasePRQLListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSelectionList is called when production selectionList is entered.
func (s *BasePRQLListener) EnterSelectionList(ctx *SelectionListContext) {}

// ExitSelectionList is called when production selectionList is exited.
func (s *BasePRQLListener) ExitSelectionList(ctx *SelectionListContext) {}

// EnterAliasedSelection is called when production AliasedSelection is entered.
func (s *BasePRQLListener) EnterAliasedSelection(ctx *AliasedSelectionContext) {}

// ExitAliasedSelection is called when production AliasedSelection is exited.
func (s *BasePRQLListener) ExitAliasedSelection(ctx *AliasedSelectionContext) {}

// EnterBareSelection is called when production BareSelection is entered.
func (s *BasePRQLListener) EnterBareSelection(ctx *BareSelectionContext) {}

// ExitBareSelection is called when production BareSelection is exited.
func (s *BasePRQLListener) ExitBareSelection(ctx *BareSelectionContext) {}

// EnterKeyList is called when production keyList is entered.
func (s *BasePRQLListener) EnterKeyList(ctx *KeyListContext) {}

// ExitKeyList is called when production keyList is exited.
func (s *BasePRQLListener) ExitKeyList(ctx *KeyListContext) {}

// EnterComputedKey is called when production ComputedKey is entered.
func (s *BasePRQLListener) EnterComputedKey(ctx *ComputedKeyContext) {}

// ExitComputedKey is called when production ComputedKey is exited.
func (s *BasePRQLListener) ExitComputedKey(ctx *ComputedKeyContext) {}

// EnterColumnKey is called when production ColumnKey is entered.
func (s *BasePRQLListener) EnterColumnKey(ctx *ColumnKeyContext) {}

// ExitColumnKey is called when production ColumnKey is exited.
func (s *BasePRQLListener) ExitColumnKey(ctx *ColumnKeyContext) {}

// EnterSortList is called when production sortList is entered.
func (s *BasePRQLListener) EnterSortList(ctx *SortListContext) {}

// ExitSortList is called when production sortList is exited.
func (s *BasePRQLListener) ExitSortList(ctx *SortListContext) {}

// EnterDescSort is called when production DescSort is entered.
func (s *BasePRQLListener) EnterDescSort(ctx *DescSortContext) {}

// ExitDescSort is called when production DescSort is exited.
func (s *BasePRQLListener) ExitDescSort(ctx *DescSortContext) {}

// EnterAscSortExplicit is called when production AscSortExplicit is entered.
func (s *BasePRQLListener) EnterAscSortExplicit(ctx *AscSortExplicitContext) {}

// ExitAscSortExplicit is called when production AscSortExplicit is exited.
func (s *BasePRQLListener) ExitAscSortExplicit(ctx *AscSortExplicitContext) {}

// EnterAscSort is called when production AscSort is entered.
func (s *BasePRQLListener) EnterAscSort(ctx *AscSortContext) {}

// ExitAscSort is called when production AscSort is exited.
func (s *BasePRQLListener) ExitAscSort(ctx *AscSortContext) {}

// EnterOpaqueExpr is called when production opaqueExpr is entered.
func (s *BasePRQLListener) EnterOpaqueExpr(ctx *OpaqueExprContext) {}

// ExitOpaqueExpr is called when production opaqueExpr is exited.
func (s *BasePRQLListener) ExitOpaqueExpr(ctx *OpaqueExprContext) {}

// EnterOpaqueToken is called when production opaqueToken is entered.
func (s *BasePRQLListener) EnterOpaqueToken(ctx *OpaqueTokenContext) {}

// ExitOpaqueToken is called when production opaqueToken is exited.
func (s *BasePRQLListener) ExitOpaqueToken(ctx *OpaqueTokenContext) {}

// EnterOpaqueInner is called when production opaqueInner is entered.
func (s *BasePRQLListener) EnterOpaqueInner(ctx *OpaqueInnerContext) {}

// ExitOpaqueInner is called when production opaqueInner is exited.
func (s *BasePRQLListener) ExitOpaqueInner(ctx *OpaqueInnerContext) {}
