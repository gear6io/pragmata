// Code generated from FilterLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filtergrammar // FilterLang
import "github.com/antlr4-go/antlr/v4"

type BaseFilterLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFilterLangVisitor) VisitFilterExpr(ctx *FilterExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitGroupExpr(ctx *GroupExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitCondExpr(ctx *CondExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitCastCompare(ctx *CastCompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitCompare(ctx *CompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitInList(ctx *InListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitNotInList(ctx *NotInListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitHasElement(ctx *HasElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitHasAny(ctx *HasAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitHasAll(ctx *HasAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitContainsOp(ctx *ContainsOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitNotContainsOp(ctx *NotContainsOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitLikeOp(ctx *LikeOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitILikeOp(ctx *ILikeOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitNotLikeOp(ctx *NotLikeOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitMatchesOp(ctx *MatchesOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitExistsCheck(ctx *ExistsCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitNotExistsCheck(ctx *NotExistsCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitIsNull(ctx *IsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitIsNotNull(ctx *IsNotNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitScalarList(ctx *ScalarListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitScalarString(ctx *ScalarStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitScalarFloat(ctx *ScalarFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitScalarInt(ctx *ScalarIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitScalarTrue(ctx *ScalarTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitScalarFalse(ctx *ScalarFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitScalarNull(ctx *ScalarNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFilterLangVisitor) VisitScalarBareWord(ctx *ScalarBareWordContext) interface{} {
	return v.VisitChildren(ctx)
}
