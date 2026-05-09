// Code generated from FilterLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filtergrammar // FilterLang
import "github.com/antlr4-go/antlr/v4"

// BaseFilterLangListener is a complete listener for a parse tree produced by FilterLang.
type BaseFilterLangListener struct{}

var _ FilterLangListener = &BaseFilterLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFilterLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFilterLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFilterLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFilterLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFilterExpr is called when production filterExpr is entered.
func (s *BaseFilterLangListener) EnterFilterExpr(ctx *FilterExprContext) {}

// ExitFilterExpr is called when production filterExpr is exited.
func (s *BaseFilterLangListener) ExitFilterExpr(ctx *FilterExprContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BaseFilterLangListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BaseFilterLangListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterGroupExpr is called when production GroupExpr is entered.
func (s *BaseFilterLangListener) EnterGroupExpr(ctx *GroupExprContext) {}

// ExitGroupExpr is called when production GroupExpr is exited.
func (s *BaseFilterLangListener) ExitGroupExpr(ctx *GroupExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseFilterLangListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseFilterLangListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterCondExpr is called when production CondExpr is entered.
func (s *BaseFilterLangListener) EnterCondExpr(ctx *CondExprContext) {}

// ExitCondExpr is called when production CondExpr is exited.
func (s *BaseFilterLangListener) ExitCondExpr(ctx *CondExprContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BaseFilterLangListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BaseFilterLangListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterCastCompare is called when production CastCompare is entered.
func (s *BaseFilterLangListener) EnterCastCompare(ctx *CastCompareContext) {}

// ExitCastCompare is called when production CastCompare is exited.
func (s *BaseFilterLangListener) ExitCastCompare(ctx *CastCompareContext) {}

// EnterCompare is called when production Compare is entered.
func (s *BaseFilterLangListener) EnterCompare(ctx *CompareContext) {}

// ExitCompare is called when production Compare is exited.
func (s *BaseFilterLangListener) ExitCompare(ctx *CompareContext) {}

// EnterInList is called when production InList is entered.
func (s *BaseFilterLangListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production InList is exited.
func (s *BaseFilterLangListener) ExitInList(ctx *InListContext) {}

// EnterNotInList is called when production NotInList is entered.
func (s *BaseFilterLangListener) EnterNotInList(ctx *NotInListContext) {}

// ExitNotInList is called when production NotInList is exited.
func (s *BaseFilterLangListener) ExitNotInList(ctx *NotInListContext) {}

// EnterHasElement is called when production HasElement is entered.
func (s *BaseFilterLangListener) EnterHasElement(ctx *HasElementContext) {}

// ExitHasElement is called when production HasElement is exited.
func (s *BaseFilterLangListener) ExitHasElement(ctx *HasElementContext) {}

// EnterHasAny is called when production HasAny is entered.
func (s *BaseFilterLangListener) EnterHasAny(ctx *HasAnyContext) {}

// ExitHasAny is called when production HasAny is exited.
func (s *BaseFilterLangListener) ExitHasAny(ctx *HasAnyContext) {}

// EnterHasAll is called when production HasAll is entered.
func (s *BaseFilterLangListener) EnterHasAll(ctx *HasAllContext) {}

// ExitHasAll is called when production HasAll is exited.
func (s *BaseFilterLangListener) ExitHasAll(ctx *HasAllContext) {}

// EnterContainsOp is called when production ContainsOp is entered.
func (s *BaseFilterLangListener) EnterContainsOp(ctx *ContainsOpContext) {}

// ExitContainsOp is called when production ContainsOp is exited.
func (s *BaseFilterLangListener) ExitContainsOp(ctx *ContainsOpContext) {}

// EnterNotContainsOp is called when production NotContainsOp is entered.
func (s *BaseFilterLangListener) EnterNotContainsOp(ctx *NotContainsOpContext) {}

// ExitNotContainsOp is called when production NotContainsOp is exited.
func (s *BaseFilterLangListener) ExitNotContainsOp(ctx *NotContainsOpContext) {}

// EnterLikeOp is called when production LikeOp is entered.
func (s *BaseFilterLangListener) EnterLikeOp(ctx *LikeOpContext) {}

// ExitLikeOp is called when production LikeOp is exited.
func (s *BaseFilterLangListener) ExitLikeOp(ctx *LikeOpContext) {}

// EnterILikeOp is called when production ILikeOp is entered.
func (s *BaseFilterLangListener) EnterILikeOp(ctx *ILikeOpContext) {}

// ExitILikeOp is called when production ILikeOp is exited.
func (s *BaseFilterLangListener) ExitILikeOp(ctx *ILikeOpContext) {}

// EnterNotLikeOp is called when production NotLikeOp is entered.
func (s *BaseFilterLangListener) EnterNotLikeOp(ctx *NotLikeOpContext) {}

// ExitNotLikeOp is called when production NotLikeOp is exited.
func (s *BaseFilterLangListener) ExitNotLikeOp(ctx *NotLikeOpContext) {}

// EnterMatchesOp is called when production MatchesOp is entered.
func (s *BaseFilterLangListener) EnterMatchesOp(ctx *MatchesOpContext) {}

// ExitMatchesOp is called when production MatchesOp is exited.
func (s *BaseFilterLangListener) ExitMatchesOp(ctx *MatchesOpContext) {}

// EnterExistsCheck is called when production ExistsCheck is entered.
func (s *BaseFilterLangListener) EnterExistsCheck(ctx *ExistsCheckContext) {}

// ExitExistsCheck is called when production ExistsCheck is exited.
func (s *BaseFilterLangListener) ExitExistsCheck(ctx *ExistsCheckContext) {}

// EnterNotExistsCheck is called when production NotExistsCheck is entered.
func (s *BaseFilterLangListener) EnterNotExistsCheck(ctx *NotExistsCheckContext) {}

// ExitNotExistsCheck is called when production NotExistsCheck is exited.
func (s *BaseFilterLangListener) ExitNotExistsCheck(ctx *NotExistsCheckContext) {}

// EnterIsNull is called when production IsNull is entered.
func (s *BaseFilterLangListener) EnterIsNull(ctx *IsNullContext) {}

// ExitIsNull is called when production IsNull is exited.
func (s *BaseFilterLangListener) ExitIsNull(ctx *IsNullContext) {}

// EnterIsNotNull is called when production IsNotNull is entered.
func (s *BaseFilterLangListener) EnterIsNotNull(ctx *IsNotNullContext) {}

// ExitIsNotNull is called when production IsNotNull is exited.
func (s *BaseFilterLangListener) ExitIsNotNull(ctx *IsNotNullContext) {}

// EnterOp is called when production op is entered.
func (s *BaseFilterLangListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseFilterLangListener) ExitOp(ctx *OpContext) {}

// EnterPath is called when production path is entered.
func (s *BaseFilterLangListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseFilterLangListener) ExitPath(ctx *PathContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseFilterLangListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseFilterLangListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterScalarList is called when production scalarList is entered.
func (s *BaseFilterLangListener) EnterScalarList(ctx *ScalarListContext) {}

// ExitScalarList is called when production scalarList is exited.
func (s *BaseFilterLangListener) ExitScalarList(ctx *ScalarListContext) {}

// EnterScalarString is called when production ScalarString is entered.
func (s *BaseFilterLangListener) EnterScalarString(ctx *ScalarStringContext) {}

// ExitScalarString is called when production ScalarString is exited.
func (s *BaseFilterLangListener) ExitScalarString(ctx *ScalarStringContext) {}

// EnterScalarFloat is called when production ScalarFloat is entered.
func (s *BaseFilterLangListener) EnterScalarFloat(ctx *ScalarFloatContext) {}

// ExitScalarFloat is called when production ScalarFloat is exited.
func (s *BaseFilterLangListener) ExitScalarFloat(ctx *ScalarFloatContext) {}

// EnterScalarInt is called when production ScalarInt is entered.
func (s *BaseFilterLangListener) EnterScalarInt(ctx *ScalarIntContext) {}

// ExitScalarInt is called when production ScalarInt is exited.
func (s *BaseFilterLangListener) ExitScalarInt(ctx *ScalarIntContext) {}

// EnterScalarTrue is called when production ScalarTrue is entered.
func (s *BaseFilterLangListener) EnterScalarTrue(ctx *ScalarTrueContext) {}

// ExitScalarTrue is called when production ScalarTrue is exited.
func (s *BaseFilterLangListener) ExitScalarTrue(ctx *ScalarTrueContext) {}

// EnterScalarFalse is called when production ScalarFalse is entered.
func (s *BaseFilterLangListener) EnterScalarFalse(ctx *ScalarFalseContext) {}

// ExitScalarFalse is called when production ScalarFalse is exited.
func (s *BaseFilterLangListener) ExitScalarFalse(ctx *ScalarFalseContext) {}

// EnterScalarNull is called when production ScalarNull is entered.
func (s *BaseFilterLangListener) EnterScalarNull(ctx *ScalarNullContext) {}

// ExitScalarNull is called when production ScalarNull is exited.
func (s *BaseFilterLangListener) ExitScalarNull(ctx *ScalarNullContext) {}

// EnterScalarBareWord is called when production ScalarBareWord is entered.
func (s *BaseFilterLangListener) EnterScalarBareWord(ctx *ScalarBareWordContext) {}

// ExitScalarBareWord is called when production ScalarBareWord is exited.
func (s *BaseFilterLangListener) ExitScalarBareWord(ctx *ScalarBareWordContext) {}
