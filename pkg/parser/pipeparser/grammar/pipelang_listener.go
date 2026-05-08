// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

// PipeLangListener is a complete listener for a parse tree produced by PipeLang.
type PipeLangListener interface {
	antlr.ParseTreeListener

	// EnterPipeFile is called when entering the pipeFile production.
	EnterPipeFile(c *PipeFileContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDescriptionDir is called when entering the descriptionDir production.
	EnterDescriptionDir(c *DescriptionDirContext)

	// EnterTagsDir is called when entering the tagsDir production.
	EnterTagsDir(c *TagsDirContext)

	// EnterTypeDir is called when entering the typeDir production.
	EnterTypeDir(c *TypeDirContext)

	// EnterDatasourceDir is called when entering the datasourceDir production.
	EnterDatasourceDir(c *DatasourceDirContext)

	// EnterTargetDatasourceDir is called when entering the targetDatasourceDir production.
	EnterTargetDatasourceDir(c *TargetDatasourceDirContext)

	// EnterCopyScheduleDir is called when entering the copyScheduleDir production.
	EnterCopyScheduleDir(c *CopyScheduleDirContext)

	// EnterNodeBlock is called when entering the nodeBlock production.
	EnterNodeBlock(c *NodeBlockContext)

	// EnterSqlBlock is called when entering the sqlBlock production.
	EnterSqlBlock(c *SqlBlockContext)

	// EnterSqlBody is called when entering the sqlBody production.
	EnterSqlBody(c *SqlBodyContext)

	// ExitPipeFile is called when exiting the pipeFile production.
	ExitPipeFile(c *PipeFileContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDescriptionDir is called when exiting the descriptionDir production.
	ExitDescriptionDir(c *DescriptionDirContext)

	// ExitTagsDir is called when exiting the tagsDir production.
	ExitTagsDir(c *TagsDirContext)

	// ExitTypeDir is called when exiting the typeDir production.
	ExitTypeDir(c *TypeDirContext)

	// ExitDatasourceDir is called when exiting the datasourceDir production.
	ExitDatasourceDir(c *DatasourceDirContext)

	// ExitTargetDatasourceDir is called when exiting the targetDatasourceDir production.
	ExitTargetDatasourceDir(c *TargetDatasourceDirContext)

	// ExitCopyScheduleDir is called when exiting the copyScheduleDir production.
	ExitCopyScheduleDir(c *CopyScheduleDirContext)

	// ExitNodeBlock is called when exiting the nodeBlock production.
	ExitNodeBlock(c *NodeBlockContext)

	// ExitSqlBlock is called when exiting the sqlBlock production.
	ExitSqlBlock(c *SqlBlockContext)

	// ExitSqlBody is called when exiting the sqlBody production.
	ExitSqlBody(c *SqlBodyContext)
}
