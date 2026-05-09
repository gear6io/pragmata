// Code generated from Jinja.g4 by ANTLR 4.13.2. DO NOT EDIT.

package jinjagrammar // Jinja
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Jinja.
type JinjaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Jinja#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by Jinja#rawText.
	VisitRawText(ctx *RawTextContext) interface{}

	// Visit a parse tree produced by Jinja#exprTag.
	VisitExprTag(ctx *ExprTagContext) interface{}

	// Visit a parse tree produced by Jinja#ifTag.
	VisitIfTag(ctx *IfTagContext) interface{}

	// Visit a parse tree produced by Jinja#forTag.
	VisitForTag(ctx *ForTagContext) interface{}

	// Visit a parse tree produced by Jinja#macroTag.
	VisitMacroTag(ctx *MacroTagContext) interface{}

	// Visit a parse tree produced by Jinja#setTag.
	VisitSetTag(ctx *SetTagContext) interface{}

	// Visit a parse tree produced by Jinja#commentTag.
	VisitCommentTag(ctx *CommentTagContext) interface{}

	// Visit a parse tree produced by Jinja#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by Jinja#fieldAccess.
	VisitFieldAccess(ctx *FieldAccessContext) interface{}

	// Visit a parse tree produced by Jinja#filterExpr.
	VisitFilterExpr(ctx *FilterExprContext) interface{}

	// Visit a parse tree produced by Jinja#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by Jinja#nameExpr.
	VisitNameExpr(ctx *NameExprContext) interface{}

	// Visit a parse tree produced by Jinja#stringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by Jinja#numberExpr.
	VisitNumberExpr(ctx *NumberExprContext) interface{}

	// Visit a parse tree produced by Jinja#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by Jinja#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by Jinja#kwarg.
	VisitKwarg(ctx *KwargContext) interface{}

	// Visit a parse tree produced by Jinja#posarg.
	VisitPosarg(ctx *PosargContext) interface{}

	// Visit a parse tree produced by Jinja#ifBlock.
	VisitIfBlock(ctx *IfBlockContext) interface{}

	// Visit a parse tree produced by Jinja#elifClause.
	VisitElifClause(ctx *ElifClauseContext) interface{}

	// Visit a parse tree produced by Jinja#elseClause.
	VisitElseClause(ctx *ElseClauseContext) interface{}

	// Visit a parse tree produced by Jinja#blockExpr.
	VisitBlockExpr(ctx *BlockExprContext) interface{}

	// Visit a parse tree produced by Jinja#forBlock.
	VisitForBlock(ctx *ForBlockContext) interface{}

	// Visit a parse tree produced by Jinja#macroBlock.
	VisitMacroBlock(ctx *MacroBlockContext) interface{}

	// Visit a parse tree produced by Jinja#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by Jinja#namedParam.
	VisitNamedParam(ctx *NamedParamContext) interface{}

	// Visit a parse tree produced by Jinja#setStmt.
	VisitSetStmt(ctx *SetStmtContext) interface{}
}
