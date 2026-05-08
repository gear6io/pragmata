// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
import "github.com/antlr4-go/antlr/v4"

type BasePipeLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePipeLangVisitor) VisitPipeFile(ctx *PipeFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitDescriptionDir(ctx *DescriptionDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitTagsDir(ctx *TagsDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitTypeDir(ctx *TypeDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitDatasourceDir(ctx *DatasourceDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitTargetDatasourceDir(ctx *TargetDatasourceDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitCopyScheduleDir(ctx *CopyScheduleDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitNodeBlock(ctx *NodeBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitSqlBlock(ctx *SqlBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipeLangVisitor) VisitSqlBody(ctx *SqlBodyContext) interface{} {
	return v.VisitChildren(ctx)
}
