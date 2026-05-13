// Code generated from FilterQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filtergrammar // FilterQuery
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FilterQuery.
type FilterQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FilterQuery#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by FilterQuery#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by FilterQuery#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by FilterQuery#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by FilterQuery#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by FilterQuery#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by FilterQuery#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by FilterQuery#inClause.
	VisitInClause(ctx *InClauseContext) interface{}

	// Visit a parse tree produced by FilterQuery#notInClause.
	VisitNotInClause(ctx *NotInClauseContext) interface{}

	// Visit a parse tree produced by FilterQuery#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by FilterQuery#fullText.
	VisitFullText(ctx *FullTextContext) interface{}

	// Visit a parse tree produced by FilterQuery#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by FilterQuery#functionParamList.
	VisitFunctionParamList(ctx *FunctionParamListContext) interface{}

	// Visit a parse tree produced by FilterQuery#functionParam.
	VisitFunctionParam(ctx *FunctionParamContext) interface{}

	// Visit a parse tree produced by FilterQuery#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by FilterQuery#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by FilterQuery#key.
	VisitKey(ctx *KeyContext) interface{}
}
