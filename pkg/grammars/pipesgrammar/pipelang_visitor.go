// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PipeLang.
type PipeLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PipeLang#pipeFile.
	VisitPipeFile(ctx *PipeFileContext) interface{}

	// Visit a parse tree produced by PipeLang#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by PipeLang#descriptionDir.
	VisitDescriptionDir(ctx *DescriptionDirContext) interface{}

	// Visit a parse tree produced by PipeLang#tagsDir.
	VisitTagsDir(ctx *TagsDirContext) interface{}

	// Visit a parse tree produced by PipeLang#typeDir.
	VisitTypeDir(ctx *TypeDirContext) interface{}

	// Visit a parse tree produced by PipeLang#datasourceDir.
	VisitDatasourceDir(ctx *DatasourceDirContext) interface{}

	// Visit a parse tree produced by PipeLang#targetDatasourceDir.
	VisitTargetDatasourceDir(ctx *TargetDatasourceDirContext) interface{}

	// Visit a parse tree produced by PipeLang#copyScheduleDir.
	VisitCopyScheduleDir(ctx *CopyScheduleDirContext) interface{}

	// Visit a parse tree produced by PipeLang#nodeBlock.
	VisitNodeBlock(ctx *NodeBlockContext) interface{}

	// Visit a parse tree produced by PipeLang#sqlBlock.
	VisitSqlBlock(ctx *SqlBlockContext) interface{}

	// Visit a parse tree produced by PipeLang#sqlBody.
	VisitSqlBody(ctx *SqlBodyContext) interface{}
}
