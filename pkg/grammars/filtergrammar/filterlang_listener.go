// Code generated from FilterLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filtergrammar // FilterLang
import "github.com/antlr4-go/antlr/v4"

// FilterLangListener is a complete listener for a parse tree produced by FilterLang.
type FilterLangListener interface {
	antlr.ParseTreeListener

	// EnterFilterExpr is called when entering the filterExpr production.
	EnterFilterExpr(c *FilterExprContext)

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterGroupExpr is called when entering the GroupExpr production.
	EnterGroupExpr(c *GroupExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterCondExpr is called when entering the CondExpr production.
	EnterCondExpr(c *CondExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterCastCompare is called when entering the CastCompare production.
	EnterCastCompare(c *CastCompareContext)

	// EnterCompare is called when entering the Compare production.
	EnterCompare(c *CompareContext)

	// EnterInList is called when entering the InList production.
	EnterInList(c *InListContext)

	// EnterNotInList is called when entering the NotInList production.
	EnterNotInList(c *NotInListContext)

	// EnterHasElement is called when entering the HasElement production.
	EnterHasElement(c *HasElementContext)

	// EnterHasAny is called when entering the HasAny production.
	EnterHasAny(c *HasAnyContext)

	// EnterHasAll is called when entering the HasAll production.
	EnterHasAll(c *HasAllContext)

	// EnterContainsOp is called when entering the ContainsOp production.
	EnterContainsOp(c *ContainsOpContext)

	// EnterNotContainsOp is called when entering the NotContainsOp production.
	EnterNotContainsOp(c *NotContainsOpContext)

	// EnterLikeOp is called when entering the LikeOp production.
	EnterLikeOp(c *LikeOpContext)

	// EnterILikeOp is called when entering the ILikeOp production.
	EnterILikeOp(c *ILikeOpContext)

	// EnterNotLikeOp is called when entering the NotLikeOp production.
	EnterNotLikeOp(c *NotLikeOpContext)

	// EnterMatchesOp is called when entering the MatchesOp production.
	EnterMatchesOp(c *MatchesOpContext)

	// EnterExistsCheck is called when entering the ExistsCheck production.
	EnterExistsCheck(c *ExistsCheckContext)

	// EnterNotExistsCheck is called when entering the NotExistsCheck production.
	EnterNotExistsCheck(c *NotExistsCheckContext)

	// EnterIsNull is called when entering the IsNull production.
	EnterIsNull(c *IsNullContext)

	// EnterIsNotNull is called when entering the IsNotNull production.
	EnterIsNotNull(c *IsNotNullContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterScalarList is called when entering the scalarList production.
	EnterScalarList(c *ScalarListContext)

	// EnterScalarString is called when entering the ScalarString production.
	EnterScalarString(c *ScalarStringContext)

	// EnterScalarFloat is called when entering the ScalarFloat production.
	EnterScalarFloat(c *ScalarFloatContext)

	// EnterScalarInt is called when entering the ScalarInt production.
	EnterScalarInt(c *ScalarIntContext)

	// EnterScalarTrue is called when entering the ScalarTrue production.
	EnterScalarTrue(c *ScalarTrueContext)

	// EnterScalarFalse is called when entering the ScalarFalse production.
	EnterScalarFalse(c *ScalarFalseContext)

	// EnterScalarNull is called when entering the ScalarNull production.
	EnterScalarNull(c *ScalarNullContext)

	// EnterScalarBareWord is called when entering the ScalarBareWord production.
	EnterScalarBareWord(c *ScalarBareWordContext)

	// ExitFilterExpr is called when exiting the filterExpr production.
	ExitFilterExpr(c *FilterExprContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitGroupExpr is called when exiting the GroupExpr production.
	ExitGroupExpr(c *GroupExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitCondExpr is called when exiting the CondExpr production.
	ExitCondExpr(c *CondExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitCastCompare is called when exiting the CastCompare production.
	ExitCastCompare(c *CastCompareContext)

	// ExitCompare is called when exiting the Compare production.
	ExitCompare(c *CompareContext)

	// ExitInList is called when exiting the InList production.
	ExitInList(c *InListContext)

	// ExitNotInList is called when exiting the NotInList production.
	ExitNotInList(c *NotInListContext)

	// ExitHasElement is called when exiting the HasElement production.
	ExitHasElement(c *HasElementContext)

	// ExitHasAny is called when exiting the HasAny production.
	ExitHasAny(c *HasAnyContext)

	// ExitHasAll is called when exiting the HasAll production.
	ExitHasAll(c *HasAllContext)

	// ExitContainsOp is called when exiting the ContainsOp production.
	ExitContainsOp(c *ContainsOpContext)

	// ExitNotContainsOp is called when exiting the NotContainsOp production.
	ExitNotContainsOp(c *NotContainsOpContext)

	// ExitLikeOp is called when exiting the LikeOp production.
	ExitLikeOp(c *LikeOpContext)

	// ExitILikeOp is called when exiting the ILikeOp production.
	ExitILikeOp(c *ILikeOpContext)

	// ExitNotLikeOp is called when exiting the NotLikeOp production.
	ExitNotLikeOp(c *NotLikeOpContext)

	// ExitMatchesOp is called when exiting the MatchesOp production.
	ExitMatchesOp(c *MatchesOpContext)

	// ExitExistsCheck is called when exiting the ExistsCheck production.
	ExitExistsCheck(c *ExistsCheckContext)

	// ExitNotExistsCheck is called when exiting the NotExistsCheck production.
	ExitNotExistsCheck(c *NotExistsCheckContext)

	// ExitIsNull is called when exiting the IsNull production.
	ExitIsNull(c *IsNullContext)

	// ExitIsNotNull is called when exiting the IsNotNull production.
	ExitIsNotNull(c *IsNotNullContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitScalarList is called when exiting the scalarList production.
	ExitScalarList(c *ScalarListContext)

	// ExitScalarString is called when exiting the ScalarString production.
	ExitScalarString(c *ScalarStringContext)

	// ExitScalarFloat is called when exiting the ScalarFloat production.
	ExitScalarFloat(c *ScalarFloatContext)

	// ExitScalarInt is called when exiting the ScalarInt production.
	ExitScalarInt(c *ScalarIntContext)

	// ExitScalarTrue is called when exiting the ScalarTrue production.
	ExitScalarTrue(c *ScalarTrueContext)

	// ExitScalarFalse is called when exiting the ScalarFalse production.
	ExitScalarFalse(c *ScalarFalseContext)

	// ExitScalarNull is called when exiting the ScalarNull production.
	ExitScalarNull(c *ScalarNullContext)

	// ExitScalarBareWord is called when exiting the ScalarBareWord production.
	ExitScalarBareWord(c *ScalarBareWordContext)
}
