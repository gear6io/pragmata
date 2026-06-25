// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

// PipeLangListener is a complete listener for a parse tree produced by PipeLang.
type PipeLangListener interface {
	antlr.ParseTreeListener

	// EnterPipeFile is called when entering the pipeFile production.
	EnterPipeFile(c *PipeFileContext)

	// EnterTypeDir is called when entering the typeDir production.
	EnterTypeDir(c *TypeDirContext)

	// EnterNameDir is called when entering the nameDir production.
	EnterNameDir(c *NameDirContext)

	// EnterDescriptionDir is called when entering the descriptionDir production.
	EnterDescriptionDir(c *DescriptionDirContext)

	// EnterDescriptionMLDir is called when entering the descriptionMLDir production.
	EnterDescriptionMLDir(c *DescriptionMLDirContext)

	// EnterTagsDir is called when entering the tagsDir production.
	EnterTagsDir(c *TagsDirContext)

	// EnterOwnerDir is called when entering the ownerDir production.
	EnterOwnerDir(c *OwnerDirContext)

	// EnterDestinationDir is called when entering the destinationDir production.
	EnterDestinationDir(c *DestinationDirContext)

	// EnterScheduleDir is called when entering the scheduleDir production.
	EnterScheduleDir(c *ScheduleDirContext)

	// EnterUniqueKeyDir is called when entering the uniqueKeyDir production.
	EnterUniqueKeyDir(c *UniqueKeyDirContext)

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

	// ExitTypeDir is called when exiting the typeDir production.
	ExitTypeDir(c *TypeDirContext)

	// ExitNameDir is called when exiting the nameDir production.
	ExitNameDir(c *NameDirContext)

	// ExitDescriptionDir is called when exiting the descriptionDir production.
	ExitDescriptionDir(c *DescriptionDirContext)

	// ExitDescriptionMLDir is called when exiting the descriptionMLDir production.
	ExitDescriptionMLDir(c *DescriptionMLDirContext)

	// ExitTagsDir is called when exiting the tagsDir production.
	ExitTagsDir(c *TagsDirContext)

	// ExitOwnerDir is called when exiting the ownerDir production.
	ExitOwnerDir(c *OwnerDirContext)

	// ExitDestinationDir is called when exiting the destinationDir production.
	ExitDestinationDir(c *DestinationDirContext)

	// ExitScheduleDir is called when exiting the scheduleDir production.
	ExitScheduleDir(c *ScheduleDirContext)

	// ExitUniqueKeyDir is called when exiting the uniqueKeyDir production.
	ExitUniqueKeyDir(c *UniqueKeyDirContext)

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
