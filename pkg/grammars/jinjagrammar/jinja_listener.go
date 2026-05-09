// Code generated from Jinja.g4 by ANTLR 4.13.2. DO NOT EDIT.

package jinjagrammar // Jinja
import "github.com/antlr4-go/antlr/v4"

// JinjaListener is a complete listener for a parse tree produced by Jinja.
type JinjaListener interface {
	antlr.ParseTreeListener

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterRawText is called when entering the rawText production.
	EnterRawText(c *RawTextContext)

	// EnterExprTag is called when entering the exprTag production.
	EnterExprTag(c *ExprTagContext)

	// EnterIfTag is called when entering the ifTag production.
	EnterIfTag(c *IfTagContext)

	// EnterForTag is called when entering the forTag production.
	EnterForTag(c *ForTagContext)

	// EnterMacroTag is called when entering the macroTag production.
	EnterMacroTag(c *MacroTagContext)

	// EnterSetTag is called when entering the setTag production.
	EnterSetTag(c *SetTagContext)

	// EnterCommentTag is called when entering the commentTag production.
	EnterCommentTag(c *CommentTagContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterFieldAccess is called when entering the fieldAccess production.
	EnterFieldAccess(c *FieldAccessContext)

	// EnterFilterExpr is called when entering the filterExpr production.
	EnterFilterExpr(c *FilterExprContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterNameExpr is called when entering the nameExpr production.
	EnterNameExpr(c *NameExprContext)

	// EnterStringExpr is called when entering the stringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterNumberExpr is called when entering the numberExpr production.
	EnterNumberExpr(c *NumberExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterKwarg is called when entering the kwarg production.
	EnterKwarg(c *KwargContext)

	// EnterPosarg is called when entering the posarg production.
	EnterPosarg(c *PosargContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterElifClause is called when entering the elifClause production.
	EnterElifClause(c *ElifClauseContext)

	// EnterElseClause is called when entering the elseClause production.
	EnterElseClause(c *ElseClauseContext)

	// EnterBlockExpr is called when entering the blockExpr production.
	EnterBlockExpr(c *BlockExprContext)

	// EnterForBlock is called when entering the forBlock production.
	EnterForBlock(c *ForBlockContext)

	// EnterMacroBlock is called when entering the macroBlock production.
	EnterMacroBlock(c *MacroBlockContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterNamedParam is called when entering the namedParam production.
	EnterNamedParam(c *NamedParamContext)

	// EnterSetStmt is called when entering the setStmt production.
	EnterSetStmt(c *SetStmtContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitRawText is called when exiting the rawText production.
	ExitRawText(c *RawTextContext)

	// ExitExprTag is called when exiting the exprTag production.
	ExitExprTag(c *ExprTagContext)

	// ExitIfTag is called when exiting the ifTag production.
	ExitIfTag(c *IfTagContext)

	// ExitForTag is called when exiting the forTag production.
	ExitForTag(c *ForTagContext)

	// ExitMacroTag is called when exiting the macroTag production.
	ExitMacroTag(c *MacroTagContext)

	// ExitSetTag is called when exiting the setTag production.
	ExitSetTag(c *SetTagContext)

	// ExitCommentTag is called when exiting the commentTag production.
	ExitCommentTag(c *CommentTagContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitFieldAccess is called when exiting the fieldAccess production.
	ExitFieldAccess(c *FieldAccessContext)

	// ExitFilterExpr is called when exiting the filterExpr production.
	ExitFilterExpr(c *FilterExprContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitNameExpr is called when exiting the nameExpr production.
	ExitNameExpr(c *NameExprContext)

	// ExitStringExpr is called when exiting the stringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitNumberExpr is called when exiting the numberExpr production.
	ExitNumberExpr(c *NumberExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitKwarg is called when exiting the kwarg production.
	ExitKwarg(c *KwargContext)

	// ExitPosarg is called when exiting the posarg production.
	ExitPosarg(c *PosargContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitElifClause is called when exiting the elifClause production.
	ExitElifClause(c *ElifClauseContext)

	// ExitElseClause is called when exiting the elseClause production.
	ExitElseClause(c *ElseClauseContext)

	// ExitBlockExpr is called when exiting the blockExpr production.
	ExitBlockExpr(c *BlockExprContext)

	// ExitForBlock is called when exiting the forBlock production.
	ExitForBlock(c *ForBlockContext)

	// ExitMacroBlock is called when exiting the macroBlock production.
	ExitMacroBlock(c *MacroBlockContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitNamedParam is called when exiting the namedParam production.
	ExitNamedParam(c *NamedParamContext)

	// ExitSetStmt is called when exiting the setStmt production.
	ExitSetStmt(c *SetStmtContext)
}
