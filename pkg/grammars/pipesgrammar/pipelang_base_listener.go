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

// EnterSourcesDir is called when production sourcesDir is entered.
func (s *BasePipeLangListener) EnterSourcesDir(ctx *SourcesDirContext) {}

// ExitSourcesDir is called when production sourcesDir is exited.
func (s *BasePipeLangListener) ExitSourcesDir(ctx *SourcesDirContext) {}

// EnterParamsDir is called when production paramsDir is entered.
func (s *BasePipeLangListener) EnterParamsDir(ctx *ParamsDirContext) {}

// ExitParamsDir is called when production paramsDir is exited.
func (s *BasePipeLangListener) ExitParamsDir(ctx *ParamsDirContext) {}

// EnterPipelineDir is called when production pipelineDir is entered.
func (s *BasePipeLangListener) EnterPipelineDir(ctx *PipelineDirContext) {}

// ExitPipelineDir is called when production pipelineDir is exited.
func (s *BasePipeLangListener) ExitPipelineDir(ctx *PipelineDirContext) {}

// EnterPipelineBlock is called when production pipelineBlock is entered.
func (s *BasePipeLangListener) EnterPipelineBlock(ctx *PipelineBlockContext) {}

// ExitPipelineBlock is called when production pipelineBlock is exited.
func (s *BasePipeLangListener) ExitPipelineBlock(ctx *PipelineBlockContext) {}

// EnterPipelineNode is called when production pipelineNode is entered.
func (s *BasePipeLangListener) EnterPipelineNode(ctx *PipelineNodeContext) {}

// ExitPipelineNode is called when production pipelineNode is exited.
func (s *BasePipeLangListener) ExitPipelineNode(ctx *PipelineNodeContext) {}
