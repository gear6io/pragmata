// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PipeLang.
type PipeLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PipeLang#pipeFile.
	VisitPipeFile(ctx *PipeFileContext) interface{}

	// Visit a parse tree produced by PipeLang#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by PipeLang#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by PipeLang#description.
	VisitDescription(ctx *DescriptionContext) interface{}

	// Visit a parse tree produced by PipeLang#descriptionML.
	VisitDescriptionML(ctx *DescriptionMLContext) interface{}

	// Visit a parse tree produced by PipeLang#tags.
	VisitTags(ctx *TagsContext) interface{}

	// Visit a parse tree produced by PipeLang#owner.
	VisitOwner(ctx *OwnerContext) interface{}

	// Visit a parse tree produced by PipeLang#destination.
	VisitDestination(ctx *DestinationContext) interface{}

	// Visit a parse tree produced by PipeLang#schedule.
	VisitSchedule(ctx *ScheduleContext) interface{}

	// Visit a parse tree produced by PipeLang#uniqueKey.
	VisitUniqueKey(ctx *UniqueKeyContext) interface{}

	// Visit a parse tree produced by PipeLang#sourcesClause.
	VisitSourcesClause(ctx *SourcesClauseContext) interface{}

	// Visit a parse tree produced by PipeLang#paramsClause.
	VisitParamsClause(ctx *ParamsClauseContext) interface{}

	// Visit a parse tree produced by PipeLang#pipelineClause.
	VisitPipelineClause(ctx *PipelineClauseContext) interface{}

	// Visit a parse tree produced by PipeLang#pipelineBlock.
	VisitPipelineBlock(ctx *PipelineBlockContext) interface{}

	// Visit a parse tree produced by PipeLang#pipelineNode.
	VisitPipelineNode(ctx *PipelineNodeContext) interface{}
}
