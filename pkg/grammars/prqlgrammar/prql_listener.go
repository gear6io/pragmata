// Code generated from PRQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package prqlgrammar // PRQL
import "github.com/antlr4-go/antlr/v4"

// PRQLListener is a complete listener for a parse tree produced by PRQL.
type PRQLListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterFilterClause is called when entering the filterClause production.
	EnterFilterClause(c *FilterClauseContext)

	// EnterFilterBody is called when entering the filterBody production.
	EnterFilterBody(c *FilterBodyContext)

	// EnterDeriveClause is called when entering the deriveClause production.
	EnterDeriveClause(c *DeriveClauseContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterGroupClause is called when entering the groupClause production.
	EnterGroupClause(c *GroupClauseContext)

	// EnterJoinClause is called when entering the joinClause production.
	EnterJoinClause(c *JoinClauseContext)

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

	// EnterArrayJoinClause is called when entering the arrayJoinClause production.
	EnterArrayJoinClause(c *ArrayJoinClauseContext)

	// EnterSortClause is called when entering the sortClause production.
	EnterSortClause(c *SortClauseContext)

	// EnterTakeClause is called when entering the takeClause production.
	EnterTakeClause(c *TakeClauseContext)

	// EnterSkipClause is called when entering the skipClause production.
	EnterSkipClause(c *SkipClauseContext)

	// EnterWindowClause is called when entering the windowClause production.
	EnterWindowClause(c *WindowClauseContext)

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

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitFilterClause is called when exiting the filterClause production.
	ExitFilterClause(c *FilterClauseContext)

	// ExitFilterBody is called when exiting the filterBody production.
	ExitFilterBody(c *FilterBodyContext)

	// ExitDeriveClause is called when exiting the deriveClause production.
	ExitDeriveClause(c *DeriveClauseContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitGroupClause is called when exiting the groupClause production.
	ExitGroupClause(c *GroupClauseContext)

	// ExitJoinClause is called when exiting the joinClause production.
	ExitJoinClause(c *JoinClauseContext)

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

	// ExitArrayJoinClause is called when exiting the arrayJoinClause production.
	ExitArrayJoinClause(c *ArrayJoinClauseContext)

	// ExitSortClause is called when exiting the sortClause production.
	ExitSortClause(c *SortClauseContext)

	// ExitTakeClause is called when exiting the takeClause production.
	ExitTakeClause(c *TakeClauseContext)

	// ExitSkipClause is called when exiting the skipClause production.
	ExitSkipClause(c *SkipClauseContext)

	// ExitWindowClause is called when exiting the windowClause production.
	ExitWindowClause(c *WindowClauseContext)

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
