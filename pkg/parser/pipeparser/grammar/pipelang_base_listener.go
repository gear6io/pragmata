// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

// BasePipeLangListener is a complete listener for a parse tree produced by PipeLang.
type BasePipeLangListener struct{}

var _ PipeLangListener = &BasePipeLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePipeLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePipeLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePipeLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePipeLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPipeFile is called when production pipeFile is entered.
func (s *BasePipeLangListener) EnterPipeFile(ctx *PipeFileContext) {}

// ExitPipeFile is called when production pipeFile is exited.
func (s *BasePipeLangListener) ExitPipeFile(ctx *PipeFileContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePipeLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePipeLangListener) ExitStatement(ctx *StatementContext) {}

// EnterDescriptionDir is called when production descriptionDir is entered.
func (s *BasePipeLangListener) EnterDescriptionDir(ctx *DescriptionDirContext) {}

// ExitDescriptionDir is called when production descriptionDir is exited.
func (s *BasePipeLangListener) ExitDescriptionDir(ctx *DescriptionDirContext) {}

// EnterTagsDir is called when production tagsDir is entered.
func (s *BasePipeLangListener) EnterTagsDir(ctx *TagsDirContext) {}

// ExitTagsDir is called when production tagsDir is exited.
func (s *BasePipeLangListener) ExitTagsDir(ctx *TagsDirContext) {}

// EnterTypeDir is called when production typeDir is entered.
func (s *BasePipeLangListener) EnterTypeDir(ctx *TypeDirContext) {}

// ExitTypeDir is called when production typeDir is exited.
func (s *BasePipeLangListener) ExitTypeDir(ctx *TypeDirContext) {}

// EnterDatasourceDir is called when production datasourceDir is entered.
func (s *BasePipeLangListener) EnterDatasourceDir(ctx *DatasourceDirContext) {}

// ExitDatasourceDir is called when production datasourceDir is exited.
func (s *BasePipeLangListener) ExitDatasourceDir(ctx *DatasourceDirContext) {}

// EnterTargetDatasourceDir is called when production targetDatasourceDir is entered.
func (s *BasePipeLangListener) EnterTargetDatasourceDir(ctx *TargetDatasourceDirContext) {}

// ExitTargetDatasourceDir is called when production targetDatasourceDir is exited.
func (s *BasePipeLangListener) ExitTargetDatasourceDir(ctx *TargetDatasourceDirContext) {}

// EnterCopyScheduleDir is called when production copyScheduleDir is entered.
func (s *BasePipeLangListener) EnterCopyScheduleDir(ctx *CopyScheduleDirContext) {}

// ExitCopyScheduleDir is called when production copyScheduleDir is exited.
func (s *BasePipeLangListener) ExitCopyScheduleDir(ctx *CopyScheduleDirContext) {}

// EnterNodeBlock is called when production nodeBlock is entered.
func (s *BasePipeLangListener) EnterNodeBlock(ctx *NodeBlockContext) {}

// ExitNodeBlock is called when production nodeBlock is exited.
func (s *BasePipeLangListener) ExitNodeBlock(ctx *NodeBlockContext) {}

// EnterSqlBlock is called when production sqlBlock is entered.
func (s *BasePipeLangListener) EnterSqlBlock(ctx *SqlBlockContext) {}

// ExitSqlBlock is called when production sqlBlock is exited.
func (s *BasePipeLangListener) ExitSqlBlock(ctx *SqlBlockContext) {}

// EnterSqlBody is called when production sqlBody is entered.
func (s *BasePipeLangListener) EnterSqlBody(ctx *SqlBodyContext) {}

// ExitSqlBody is called when production sqlBody is exited.
func (s *BasePipeLangListener) ExitSqlBody(ctx *SqlBodyContext) {}
