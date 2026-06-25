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

// EnterType is called when production type is entered.
func (s *BasePipeLangListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasePipeLangListener) ExitType(ctx *TypeContext) {}

// EnterName is called when production name is entered.
func (s *BasePipeLangListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasePipeLangListener) ExitName(ctx *NameContext) {}

// EnterDescription is called when production description is entered.
func (s *BasePipeLangListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BasePipeLangListener) ExitDescription(ctx *DescriptionContext) {}

// EnterDescriptionML is called when production descriptionML is entered.
func (s *BasePipeLangListener) EnterDescriptionML(ctx *DescriptionMLContext) {}

// ExitDescriptionML is called when production descriptionML is exited.
func (s *BasePipeLangListener) ExitDescriptionML(ctx *DescriptionMLContext) {}

// EnterTags is called when production tags is entered.
func (s *BasePipeLangListener) EnterTags(ctx *TagsContext) {}

// ExitTags is called when production tags is exited.
func (s *BasePipeLangListener) ExitTags(ctx *TagsContext) {}

// EnterOwner is called when production owner is entered.
func (s *BasePipeLangListener) EnterOwner(ctx *OwnerContext) {}

// ExitOwner is called when production owner is exited.
func (s *BasePipeLangListener) ExitOwner(ctx *OwnerContext) {}

// EnterDestination is called when production destination is entered.
func (s *BasePipeLangListener) EnterDestination(ctx *DestinationContext) {}

// ExitDestination is called when production destination is exited.
func (s *BasePipeLangListener) ExitDestination(ctx *DestinationContext) {}

// EnterSchedule is called when production schedule is entered.
func (s *BasePipeLangListener) EnterSchedule(ctx *ScheduleContext) {}

// ExitSchedule is called when production schedule is exited.
func (s *BasePipeLangListener) ExitSchedule(ctx *ScheduleContext) {}

// EnterUniqueKey is called when production uniqueKey is entered.
func (s *BasePipeLangListener) EnterUniqueKey(ctx *UniqueKeyContext) {}

// ExitUniqueKey is called when production uniqueKey is exited.
func (s *BasePipeLangListener) ExitUniqueKey(ctx *UniqueKeyContext) {}

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

// EnterPipelineBlock is called when production pipelineBlock is entered.
func (s *BasePipeLangListener) EnterPipelineBlock(ctx *PipelineBlockContext) {}

// ExitPipelineBlock is called when production pipelineBlock is exited.
func (s *BasePipeLangListener) ExitPipelineBlock(ctx *PipelineBlockContext) {}

// EnterPipelineNode is called when production pipelineNode is entered.
func (s *BasePipeLangListener) EnterPipelineNode(ctx *PipelineNodeContext) {}

// ExitPipelineNode is called when production pipelineNode is exited.
func (s *BasePipeLangListener) ExitPipelineNode(ctx *PipelineNodeContext) {}
