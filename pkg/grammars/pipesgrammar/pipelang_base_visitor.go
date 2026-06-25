// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

type BasePipeLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePipeLangVisitor) VisitPipeFile(ctx *PipeFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitDescriptionML(ctx *DescriptionMLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitTags(ctx *TagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitOwner(ctx *OwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitDestination(ctx *DestinationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitSchedule(ctx *ScheduleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitUniqueKey(ctx *UniqueKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitSourcesClause(ctx *SourcesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitParamsClause(ctx *ParamsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitPipelineClause(ctx *PipelineClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitPipelineBlock(ctx *PipelineBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitPipelineNode(ctx *PipelineNodeContext) interface{} {
	return v.VisitChildren(ctx)
}
