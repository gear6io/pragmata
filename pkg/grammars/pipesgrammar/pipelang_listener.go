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

	// EnterSourcesDir is called when entering the sourcesDir production.
	EnterSourcesDir(c *SourcesDirContext)

	// EnterParamsDir is called when entering the paramsDir production.
	EnterParamsDir(c *ParamsDirContext)

	// EnterPipelineDir is called when entering the pipelineDir production.
	EnterPipelineDir(c *PipelineDirContext)

	// EnterPipelineBlock is called when entering the pipelineBlock production.
	EnterPipelineBlock(c *PipelineBlockContext)

	// EnterPipelineNode is called when entering the pipelineNode production.
	EnterPipelineNode(c *PipelineNodeContext)

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

	// ExitSourcesDir is called when exiting the sourcesDir production.
	ExitSourcesDir(c *SourcesDirContext)

	// ExitParamsDir is called when exiting the paramsDir production.
	ExitParamsDir(c *ParamsDirContext)

	// ExitPipelineDir is called when exiting the pipelineDir production.
	ExitPipelineDir(c *PipelineDirContext)

	// ExitPipelineBlock is called when exiting the pipelineBlock production.
	ExitPipelineBlock(c *PipelineBlockContext)

	// ExitPipelineNode is called when exiting the pipelineNode production.
	ExitPipelineNode(c *PipelineNodeContext)
}
