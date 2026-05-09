// Code generated from FilterLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filtergrammar // FilterLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FilterLang.
type FilterLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FilterLang#filterExpr.
	VisitFilterExpr(ctx *FilterExprContext) interface{}

	// Visit a parse tree produced by FilterLang#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by FilterLang#GroupExpr.
	VisitGroupExpr(ctx *GroupExprContext) interface{}

	// Visit a parse tree produced by FilterLang#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by FilterLang#CondExpr.
	VisitCondExpr(ctx *CondExprContext) interface{}

	// Visit a parse tree produced by FilterLang#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by FilterLang#CastCompare.
	VisitCastCompare(ctx *CastCompareContext) interface{}

	// Visit a parse tree produced by FilterLang#Compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by FilterLang#InList.
	VisitInList(ctx *InListContext) interface{}

	// Visit a parse tree produced by FilterLang#NotInList.
	VisitNotInList(ctx *NotInListContext) interface{}

	// Visit a parse tree produced by FilterLang#HasElement.
	VisitHasElement(ctx *HasElementContext) interface{}

	// Visit a parse tree produced by FilterLang#HasAny.
	VisitHasAny(ctx *HasAnyContext) interface{}

	// Visit a parse tree produced by FilterLang#HasAll.
	VisitHasAll(ctx *HasAllContext) interface{}

	// Visit a parse tree produced by FilterLang#ContainsOp.
	VisitContainsOp(ctx *ContainsOpContext) interface{}

	// Visit a parse tree produced by FilterLang#NotContainsOp.
	VisitNotContainsOp(ctx *NotContainsOpContext) interface{}

	// Visit a parse tree produced by FilterLang#LikeOp.
	VisitLikeOp(ctx *LikeOpContext) interface{}

	// Visit a parse tree produced by FilterLang#ILikeOp.
	VisitILikeOp(ctx *ILikeOpContext) interface{}

	// Visit a parse tree produced by FilterLang#NotLikeOp.
	VisitNotLikeOp(ctx *NotLikeOpContext) interface{}

	// Visit a parse tree produced by FilterLang#MatchesOp.
	VisitMatchesOp(ctx *MatchesOpContext) interface{}

	// Visit a parse tree produced by FilterLang#ExistsCheck.
	VisitExistsCheck(ctx *ExistsCheckContext) interface{}

	// Visit a parse tree produced by FilterLang#NotExistsCheck.
	VisitNotExistsCheck(ctx *NotExistsCheckContext) interface{}

	// Visit a parse tree produced by FilterLang#IsNull.
	VisitIsNull(ctx *IsNullContext) interface{}

	// Visit a parse tree produced by FilterLang#IsNotNull.
	VisitIsNotNull(ctx *IsNotNullContext) interface{}

	// Visit a parse tree produced by FilterLang#op.
	VisitOp(ctx *OpContext) interface{}

	// Visit a parse tree produced by FilterLang#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by FilterLang#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by FilterLang#scalarList.
	VisitScalarList(ctx *ScalarListContext) interface{}

	// Visit a parse tree produced by FilterLang#ScalarString.
	VisitScalarString(ctx *ScalarStringContext) interface{}

	// Visit a parse tree produced by FilterLang#ScalarFloat.
	VisitScalarFloat(ctx *ScalarFloatContext) interface{}

	// Visit a parse tree produced by FilterLang#ScalarInt.
	VisitScalarInt(ctx *ScalarIntContext) interface{}

	// Visit a parse tree produced by FilterLang#ScalarTrue.
	VisitScalarTrue(ctx *ScalarTrueContext) interface{}

	// Visit a parse tree produced by FilterLang#ScalarFalse.
	VisitScalarFalse(ctx *ScalarFalseContext) interface{}

	// Visit a parse tree produced by FilterLang#ScalarNull.
	VisitScalarNull(ctx *ScalarNullContext) interface{}

	// Visit a parse tree produced by FilterLang#ScalarBareWord.
	VisitScalarBareWord(ctx *ScalarBareWordContext) interface{}
}
