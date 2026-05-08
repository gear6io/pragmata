// Code generated from SQLMesh.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // SQLMesh
import "github.com/antlr4-go/antlr/v4"

type BaseSQLMeshVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSQLMeshVisitor) VisitSqlmeshFile(ctx *SqlmeshFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitTopStatement(ctx *TopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitModelDef(ctx *ModelDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitNameProp(ctx *NamePropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindProp(ctx *KindPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitDialectProp(ctx *DialectPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitOwnerProp(ctx *OwnerPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitCronProp(ctx *CronPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitGrainProp(ctx *GrainPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitPartitionedByProp(ctx *PartitionedByPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitClusteredByProp(ctx *ClusteredByPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitStorageFormatProp(ctx *StorageFormatPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitTableFormatProp(ctx *TableFormatPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitRetentionProp(ctx *RetentionPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitTagsProp(ctx *TagsPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitDescriptionProp(ctx *DescriptionPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitStampProp(ctx *StampPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitColumnsProp(ctx *ColumnsPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitColumnDescriptionsProp(ctx *ColumnDescriptionsPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitPhysicalPropertiesProp(ctx *PhysicalPropertiesPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitPreStatementsProp(ctx *PreStatementsPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitPostStatementsProp(ctx *PostStatementsPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitOnVirtualUpdateProp(ctx *OnVirtualUpdatePropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitAuditsProp(ctx *AuditsPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitTimeColumnProp(ctx *TimeColumnPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitUniqueKeyProp(ctx *UniqueKeyPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitInvalidateHardDeletesProp(ctx *InvalidateHardDeletesPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitDisableRestatementProp(ctx *DisableRestatementPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitPathProp(ctx *PathPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitCsvSettingsProp(ctx *CsvSettingsPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindFull(ctx *KindFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindIncrementalByTimeRange(ctx *KindIncrementalByTimeRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindIncrementalByUniqueKey(ctx *KindIncrementalByUniqueKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindIncrementalByPartition(ctx *KindIncrementalByPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindView(ctx *KindViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindSeed(ctx *KindSeedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindExternal(ctx *KindExternalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindEmbedded(ctx *KindEmbeddedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitKindScdType2(ctx *KindScdType2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitAuditDef(ctx *AuditDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitQueryProp(ctx *QueryPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitAuditDialectProp(ctx *AuditDialectPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitAuditArray(ctx *AuditArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitAuditEntry(ctx *AuditEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitAuditEntryNameProp(ctx *AuditEntryNamePropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitAuditEntryQueryProp(ctx *AuditEntryQueryPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitPropValue(ctx *PropValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitStringLit(ctx *StringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitIntLit(ctx *IntLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitFloatLit(ctx *FloatLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitTrueLit(ctx *TrueLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitFalseLit(ctx *FalseLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitQualifiedIdentLit(ctx *QualifiedIdentLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitIdentLit(ctx *IdentLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitQualifiedIdent(ctx *QualifiedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitDictValue(ctx *DictValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitDictEntry(ctx *DictEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitMacroBracedRef(ctx *MacroBracedRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitMacroIdentRef(ctx *MacroIdentRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLMeshVisitor) VisitSqlBody(ctx *SqlBodyContext) interface{} {
	return v.VisitChildren(ctx)
}
