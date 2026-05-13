// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PipeLang.
type PipeLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PipeLang#pipeFile.
	VisitPipeFile(ctx *PipeFileContext) interface{}

	// Visit a parse tree produced by PipeLang#typeDir.
	VisitTypeDir(ctx *TypeDirContext) interface{}

	// Visit a parse tree produced by PipeLang#nameDir.
	VisitNameDir(ctx *NameDirContext) interface{}

	// Visit a parse tree produced by PipeLang#descriptionDir.
	VisitDescriptionDir(ctx *DescriptionDirContext) interface{}

	// Visit a parse tree produced by PipeLang#descriptionMLDir.
	VisitDescriptionMLDir(ctx *DescriptionMLDirContext) interface{}

	// Visit a parse tree produced by PipeLang#tagsDir.
	VisitTagsDir(ctx *TagsDirContext) interface{}

	// Visit a parse tree produced by PipeLang#ownerDir.
	VisitOwnerDir(ctx *OwnerDirContext) interface{}

	// Visit a parse tree produced by PipeLang#destinationDir.
	VisitDestinationDir(ctx *DestinationDirContext) interface{}

	// Visit a parse tree produced by PipeLang#scheduleDir.
	VisitScheduleDir(ctx *ScheduleDirContext) interface{}

	// Visit a parse tree produced by PipeLang#sourcesDir.
	VisitSourcesDir(ctx *SourcesDirContext) interface{}

	// Visit a parse tree produced by PipeLang#paramsDir.
	VisitParamsDir(ctx *ParamsDirContext) interface{}

	// Visit a parse tree produced by PipeLang#pipelineDir.
	VisitPipelineDir(ctx *PipelineDirContext) interface{}

	// Visit a parse tree produced by PipeLang#pipelineBlock.
	VisitPipelineBlock(ctx *PipelineBlockContext) interface{}

	// Visit a parse tree produced by PipeLang#pipelineNode.
	VisitPipelineNode(ctx *PipelineNodeContext) interface{}
}
