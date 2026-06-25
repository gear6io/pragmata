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

// EnterTypeDir is called when production typeDir is entered.
func (s *BasePipeLangListener) EnterTypeDir(ctx *TypeDirContext) {}

// ExitTypeDir is called when production typeDir is exited.
func (s *BasePipeLangListener) ExitTypeDir(ctx *TypeDirContext) {}

// EnterNameDir is called when production nameDir is entered.
func (s *BasePipeLangListener) EnterNameDir(ctx *NameDirContext) {}

// ExitNameDir is called when production nameDir is exited.
func (s *BasePipeLangListener) ExitNameDir(ctx *NameDirContext) {}

// EnterDescriptionDir is called when production descriptionDir is entered.
func (s *BasePipeLangListener) EnterDescriptionDir(ctx *DescriptionDirContext) {}

// ExitDescriptionDir is called when production descriptionDir is exited.
func (s *BasePipeLangListener) ExitDescriptionDir(ctx *DescriptionDirContext) {}

// EnterDescriptionMLDir is called when production descriptionMLDir is entered.
func (s *BasePipeLangListener) EnterDescriptionMLDir(ctx *DescriptionMLDirContext) {}

// ExitDescriptionMLDir is called when production descriptionMLDir is exited.
func (s *BasePipeLangListener) ExitDescriptionMLDir(ctx *DescriptionMLDirContext) {}

// EnterTagsDir is called when production tagsDir is entered.
func (s *BasePipeLangListener) EnterTagsDir(ctx *TagsDirContext) {}

// ExitTagsDir is called when production tagsDir is exited.
func (s *BasePipeLangListener) ExitTagsDir(ctx *TagsDirContext) {}

// EnterOwnerDir is called when production ownerDir is entered.
func (s *BasePipeLangListener) EnterOwnerDir(ctx *OwnerDirContext) {}

// ExitOwnerDir is called when production ownerDir is exited.
func (s *BasePipeLangListener) ExitOwnerDir(ctx *OwnerDirContext) {}

// EnterDestinationDir is called when production destinationDir is entered.
func (s *BasePipeLangListener) EnterDestinationDir(ctx *DestinationDirContext) {}

// ExitDestinationDir is called when production destinationDir is exited.
func (s *BasePipeLangListener) ExitDestinationDir(ctx *DestinationDirContext) {}

// EnterScheduleDir is called when production scheduleDir is entered.
func (s *BasePipeLangListener) EnterScheduleDir(ctx *ScheduleDirContext) {}

// ExitScheduleDir is called when production scheduleDir is exited.
func (s *BasePipeLangListener) ExitScheduleDir(ctx *ScheduleDirContext) {}

// EnterUniqueKeyDir is called when production uniqueKeyDir is entered.
func (s *BasePipeLangListener) EnterUniqueKeyDir(ctx *UniqueKeyDirContext) {}

// ExitUniqueKeyDir is called when production uniqueKeyDir is exited.
func (s *BasePipeLangListener) ExitUniqueKeyDir(ctx *UniqueKeyDirContext) {}

// EnterSourcesClause is called when production sourcesClause is entered.
func (s *BasePipeLangListener) EnterSourcesClause(ctx *SourcesClauseContext) {}

// ExitSourcesClause is called when production sourcesClause is exited.
func (s *BasePipeLangListener) ExitSourcesClause(ctx *SourcesClauseContext) {}

// EnterParamsClause is called when production paramsClause is entered.
func (s *BasePipeLangListener) EnterParamsClause(ctx *ParamsClauseContext) {}

// ExitParamsClause is called when production paramsClause is exited.
func (s *BasePipeLangListener) ExitParamsClause(ctx *ParamsClauseContext) {}

// EnterPipelineClause is called when production pipelineClause is entered.
func (s *BasePipeLangListener) EnterPipelineClause(ctx *PipelineClauseContext) {}

// ExitPipelineClause is called when production pipelineClause is exited.
func (s *BasePipeLangListener) ExitPipelineClause(ctx *PipelineClauseContext) {}

// EnterAliasedSource is called when production aliasedSource is entered.
func (s *BasePipeLangListener) EnterAliasedSource(ctx *AliasedSourceContext) {}

// ExitAliasedSource is called when production aliasedSource is exited.
func (s *BasePipeLangListener) ExitAliasedSource(ctx *AliasedSourceContext) {}

// EnterSimpleSource is called when production simpleSource is entered.
func (s *BasePipeLangListener) EnterSimpleSource(ctx *SimpleSourceContext) {}

// ExitSimpleSource is called when production simpleSource is exited.
func (s *BasePipeLangListener) ExitSimpleSource(ctx *SimpleSourceContext) {}

// EnterParam is called when production param is entered.
func (s *BasePipeLangListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasePipeLangListener) ExitParam(ctx *ParamContext) {}

// EnterParamValue is called when production paramValue is entered.
func (s *BasePipeLangListener) EnterParamValue(ctx *ParamValueContext) {}

// ExitParamValue is called when production paramValue is exited.
func (s *BasePipeLangListener) ExitParamValue(ctx *ParamValueContext) {}
