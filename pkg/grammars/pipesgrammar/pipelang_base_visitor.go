// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

type BasePipeLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePipeLangVisitor) VisitPipeFile(ctx *PipeFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitTypeDir(ctx *TypeDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitNameDir(ctx *NameDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitDescriptionDir(ctx *DescriptionDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitDescriptionMLDir(ctx *DescriptionMLDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitTagsDir(ctx *TagsDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitOwnerDir(ctx *OwnerDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitDestinationDir(ctx *DestinationDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitScheduleDir(ctx *ScheduleDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitSourcesDir(ctx *SourcesDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitParamsDir(ctx *ParamsDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitPipelineDir(ctx *PipelineDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitPipelineBlock(ctx *PipelineBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitPipelineNode(ctx *PipelineNodeContext) interface{} {
	return v.VisitChildren(ctx)
}
