// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

// PipeLangListener is a complete listener for a parse tree produced by PipeLang.
type PipeLangListener interface {
	antlr.ParseTreeListener

	// EnterPipeFile is called when entering the pipeFile production.
	EnterPipeFile(c *PipeFileContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterDescriptionML is called when entering the descriptionML production.
	EnterDescriptionML(c *DescriptionMLContext)

	// EnterTags is called when entering the tags production.
	EnterTags(c *TagsContext)

	// EnterOwner is called when entering the owner production.
	EnterOwner(c *OwnerContext)

	// EnterDestination is called when entering the destination production.
	EnterDestination(c *DestinationContext)

	// EnterSchedule is called when entering the schedule production.
	EnterSchedule(c *ScheduleContext)

	// EnterUniqueKey is called when entering the uniqueKey production.
	EnterUniqueKey(c *UniqueKeyContext)

	// EnterSourcesClause is called when entering the sourcesClause production.
	EnterSourcesClause(c *SourcesClauseContext)

	// EnterParamsClause is called when entering the paramsClause production.
	EnterParamsClause(c *ParamsClauseContext)

	// EnterPipelineClause is called when entering the pipelineClause production.
	EnterPipelineClause(c *PipelineClauseContext)

	// EnterAliasedSource is called when entering the aliasedSource production.
	EnterAliasedSource(c *AliasedSourceContext)

	// EnterSimpleSource is called when entering the simpleSource production.
	EnterSimpleSource(c *SimpleSourceContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterParamValue is called when entering the paramValue production.
	EnterParamValue(c *ParamValueContext)

	// ExitPipeFile is called when exiting the pipeFile production.
	ExitPipeFile(c *PipeFileContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitDescriptionML is called when exiting the descriptionML production.
	ExitDescriptionML(c *DescriptionMLContext)

	// ExitTags is called when exiting the tags production.
	ExitTags(c *TagsContext)

	// ExitOwner is called when exiting the owner production.
	ExitOwner(c *OwnerContext)

	// ExitDestination is called when exiting the destination production.
	ExitDestination(c *DestinationContext)

	// ExitSchedule is called when exiting the schedule production.
	ExitSchedule(c *ScheduleContext)

	// ExitUniqueKey is called when exiting the uniqueKey production.
	ExitUniqueKey(c *UniqueKeyContext)

	// ExitSourcesClause is called when exiting the sourcesClause production.
	ExitSourcesClause(c *SourcesClauseContext)

	// ExitParamsClause is called when exiting the paramsClause production.
	ExitParamsClause(c *ParamsClauseContext)

	// ExitPipelineClause is called when exiting the pipelineClause production.
	ExitPipelineClause(c *PipelineClauseContext)

	// ExitAliasedSource is called when exiting the aliasedSource production.
	ExitAliasedSource(c *AliasedSourceContext)

	// ExitSimpleSource is called when exiting the simpleSource production.
	ExitSimpleSource(c *SimpleSourceContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitParamValue is called when exiting the paramValue production.
	ExitParamValue(c *ParamValueContext)
}
