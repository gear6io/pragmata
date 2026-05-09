// Code generated from Jinja.g4 by ANTLR 4.13.2. DO NOT EDIT.

package jinjagrammar // Jinja
import "github.com/antlr4-go/antlr/v4"

// BaseJinjaListener is a complete listener for a parse tree produced by Jinja.
type BaseJinjaListener struct{}

var _ JinjaListener = &BaseJinjaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJinjaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJinjaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJinjaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJinjaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTemplate is called when production template is entered.
func (s *BaseJinjaListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production template is exited.
func (s *BaseJinjaListener) ExitTemplate(ctx *TemplateContext) {}

// EnterRawText is called when production rawText is entered.
func (s *BaseJinjaListener) EnterRawText(ctx *RawTextContext) {}

// ExitRawText is called when production rawText is exited.
func (s *BaseJinjaListener) ExitRawText(ctx *RawTextContext) {}

// EnterExprTag is called when production exprTag is entered.
func (s *BaseJinjaListener) EnterExprTag(ctx *ExprTagContext) {}

// ExitExprTag is called when production exprTag is exited.
func (s *BaseJinjaListener) ExitExprTag(ctx *ExprTagContext) {}

// EnterIfTag is called when production ifTag is entered.
func (s *BaseJinjaListener) EnterIfTag(ctx *IfTagContext) {}

// ExitIfTag is called when production ifTag is exited.
func (s *BaseJinjaListener) ExitIfTag(ctx *IfTagContext) {}

// EnterForTag is called when production forTag is entered.
func (s *BaseJinjaListener) EnterForTag(ctx *ForTagContext) {}

// ExitForTag is called when production forTag is exited.
func (s *BaseJinjaListener) ExitForTag(ctx *ForTagContext) {}

// EnterMacroTag is called when production macroTag is entered.
func (s *BaseJinjaListener) EnterMacroTag(ctx *MacroTagContext) {}

// ExitMacroTag is called when production macroTag is exited.
func (s *BaseJinjaListener) ExitMacroTag(ctx *MacroTagContext) {}

// EnterSetTag is called when production setTag is entered.
func (s *BaseJinjaListener) EnterSetTag(ctx *SetTagContext) {}

// ExitSetTag is called when production setTag is exited.
func (s *BaseJinjaListener) ExitSetTag(ctx *SetTagContext) {}

// EnterCommentTag is called when production commentTag is entered.
func (s *BaseJinjaListener) EnterCommentTag(ctx *CommentTagContext) {}

// ExitCommentTag is called when production commentTag is exited.
func (s *BaseJinjaListener) ExitCommentTag(ctx *CommentTagContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseJinjaListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseJinjaListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterFieldAccess is called when production fieldAccess is entered.
func (s *BaseJinjaListener) EnterFieldAccess(ctx *FieldAccessContext) {}

// ExitFieldAccess is called when production fieldAccess is exited.
func (s *BaseJinjaListener) ExitFieldAccess(ctx *FieldAccessContext) {}

// EnterFilterExpr is called when production filterExpr is entered.
func (s *BaseJinjaListener) EnterFilterExpr(ctx *FilterExprContext) {}

// ExitFilterExpr is called when production filterExpr is exited.
func (s *BaseJinjaListener) ExitFilterExpr(ctx *FilterExprContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BaseJinjaListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BaseJinjaListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterNameExpr is called when production nameExpr is entered.
func (s *BaseJinjaListener) EnterNameExpr(ctx *NameExprContext) {}

// ExitNameExpr is called when production nameExpr is exited.
func (s *BaseJinjaListener) ExitNameExpr(ctx *NameExprContext) {}

// EnterStringExpr is called when production stringExpr is entered.
func (s *BaseJinjaListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production stringExpr is exited.
func (s *BaseJinjaListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterNumberExpr is called when production numberExpr is entered.
func (s *BaseJinjaListener) EnterNumberExpr(ctx *NumberExprContext) {}

// ExitNumberExpr is called when production numberExpr is exited.
func (s *BaseJinjaListener) ExitNumberExpr(ctx *NumberExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseJinjaListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseJinjaListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseJinjaListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseJinjaListener) ExitArgList(ctx *ArgListContext) {}

// EnterKwarg is called when production kwarg is entered.
func (s *BaseJinjaListener) EnterKwarg(ctx *KwargContext) {}

// ExitKwarg is called when production kwarg is exited.
func (s *BaseJinjaListener) ExitKwarg(ctx *KwargContext) {}

// EnterPosarg is called when production posarg is entered.
func (s *BaseJinjaListener) EnterPosarg(ctx *PosargContext) {}

// ExitPosarg is called when production posarg is exited.
func (s *BaseJinjaListener) ExitPosarg(ctx *PosargContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseJinjaListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseJinjaListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterElifClause is called when production elifClause is entered.
func (s *BaseJinjaListener) EnterElifClause(ctx *ElifClauseContext) {}

// ExitElifClause is called when production elifClause is exited.
func (s *BaseJinjaListener) ExitElifClause(ctx *ElifClauseContext) {}

// EnterElseClause is called when production elseClause is entered.
func (s *BaseJinjaListener) EnterElseClause(ctx *ElseClauseContext) {}

// ExitElseClause is called when production elseClause is exited.
func (s *BaseJinjaListener) ExitElseClause(ctx *ElseClauseContext) {}

// EnterBlockExpr is called when production blockExpr is entered.
func (s *BaseJinjaListener) EnterBlockExpr(ctx *BlockExprContext) {}

// ExitBlockExpr is called when production blockExpr is exited.
func (s *BaseJinjaListener) ExitBlockExpr(ctx *BlockExprContext) {}

// EnterForBlock is called when production forBlock is entered.
func (s *BaseJinjaListener) EnterForBlock(ctx *ForBlockContext) {}

// ExitForBlock is called when production forBlock is exited.
func (s *BaseJinjaListener) ExitForBlock(ctx *ForBlockContext) {}

// EnterMacroBlock is called when production macroBlock is entered.
func (s *BaseJinjaListener) EnterMacroBlock(ctx *MacroBlockContext) {}

// ExitMacroBlock is called when production macroBlock is exited.
func (s *BaseJinjaListener) ExitMacroBlock(ctx *MacroBlockContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseJinjaListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseJinjaListener) ExitParamList(ctx *ParamListContext) {}

// EnterNamedParam is called when production namedParam is entered.
func (s *BaseJinjaListener) EnterNamedParam(ctx *NamedParamContext) {}

// ExitNamedParam is called when production namedParam is exited.
func (s *BaseJinjaListener) ExitNamedParam(ctx *NamedParamContext) {}

// EnterSetStmt is called when production setStmt is entered.
func (s *BaseJinjaListener) EnterSetStmt(ctx *SetStmtContext) {}

// ExitSetStmt is called when production setStmt is exited.
func (s *BaseJinjaListener) ExitSetStmt(ctx *SetStmtContext) {}
