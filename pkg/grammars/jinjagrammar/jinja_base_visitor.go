// Code generated from Jinja.g4 by ANTLR 4.13.2. DO NOT EDIT.

package jinjagrammar // Jinja
import "github.com/antlr4-go/antlr/v4"

type BaseJinjaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJinjaVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitRawText(ctx *RawTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitExprTag(ctx *ExprTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitIfTag(ctx *IfTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitForTag(ctx *ForTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitMacroTag(ctx *MacroTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitSetTag(ctx *SetTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitCommentTag(ctx *CommentTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitFieldAccess(ctx *FieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitFilterExpr(ctx *FilterExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitNameExpr(ctx *NameExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitNumberExpr(ctx *NumberExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitKwarg(ctx *KwargContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitPosarg(ctx *PosargContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitIfBlock(ctx *IfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitElifClause(ctx *ElifClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitElseClause(ctx *ElseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitBlockExpr(ctx *BlockExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitForBlock(ctx *ForBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitMacroBlock(ctx *MacroBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitNamedParam(ctx *NamedParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitSetStmt(ctx *SetStmtContext) interface{} {
	return v.VisitChildren(ctx)
}
