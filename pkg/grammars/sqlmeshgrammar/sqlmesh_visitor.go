// Code generated from SQLMesh.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // SQLMesh
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SQLMesh.
type SQLMeshVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SQLMesh#sqlmeshFile.
	VisitSqlmeshFile(ctx *SqlmeshFileContext) interface{}

	// Visit a parse tree produced by SQLMesh#topStatement.
	VisitTopStatement(ctx *TopStatementContext) interface{}

	// Visit a parse tree produced by SQLMesh#modelDef.
	VisitModelDef(ctx *ModelDefContext) interface{}

	// Visit a parse tree produced by SQLMesh#nameProp.
	VisitNameProp(ctx *NamePropContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindProp.
	VisitKindProp(ctx *KindPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#dialectProp.
	VisitDialectProp(ctx *DialectPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#ownerProp.
	VisitOwnerProp(ctx *OwnerPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#cronProp.
	VisitCronProp(ctx *CronPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#grainProp.
	VisitGrainProp(ctx *GrainPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#partitionedByProp.
	VisitPartitionedByProp(ctx *PartitionedByPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#clusteredByProp.
	VisitClusteredByProp(ctx *ClusteredByPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#storageFormatProp.
	VisitStorageFormatProp(ctx *StorageFormatPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#tableFormatProp.
	VisitTableFormatProp(ctx *TableFormatPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#retentionProp.
	VisitRetentionProp(ctx *RetentionPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#tagsProp.
	VisitTagsProp(ctx *TagsPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#descriptionProp.
	VisitDescriptionProp(ctx *DescriptionPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#stampProp.
	VisitStampProp(ctx *StampPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#columnsProp.
	VisitColumnsProp(ctx *ColumnsPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#columnDescriptionsProp.
	VisitColumnDescriptionsProp(ctx *ColumnDescriptionsPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#physicalPropertiesProp.
	VisitPhysicalPropertiesProp(ctx *PhysicalPropertiesPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#preStatementsProp.
	VisitPreStatementsProp(ctx *PreStatementsPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#postStatementsProp.
	VisitPostStatementsProp(ctx *PostStatementsPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#onVirtualUpdateProp.
	VisitOnVirtualUpdateProp(ctx *OnVirtualUpdatePropContext) interface{}

	// Visit a parse tree produced by SQLMesh#auditsProp.
	VisitAuditsProp(ctx *AuditsPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#timeColumnProp.
	VisitTimeColumnProp(ctx *TimeColumnPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#uniqueKeyProp.
	VisitUniqueKeyProp(ctx *UniqueKeyPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#invalidateHardDeletesProp.
	VisitInvalidateHardDeletesProp(ctx *InvalidateHardDeletesPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#disableRestatementProp.
	VisitDisableRestatementProp(ctx *DisableRestatementPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#pathProp.
	VisitPathProp(ctx *PathPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#csvSettingsProp.
	VisitCsvSettingsProp(ctx *CsvSettingsPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindFull.
	VisitKindFull(ctx *KindFullContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindIncrementalByTimeRange.
	VisitKindIncrementalByTimeRange(ctx *KindIncrementalByTimeRangeContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindIncrementalByUniqueKey.
	VisitKindIncrementalByUniqueKey(ctx *KindIncrementalByUniqueKeyContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindIncrementalByPartition.
	VisitKindIncrementalByPartition(ctx *KindIncrementalByPartitionContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindView.
	VisitKindView(ctx *KindViewContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindSeed.
	VisitKindSeed(ctx *KindSeedContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindExternal.
	VisitKindExternal(ctx *KindExternalContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindEmbedded.
	VisitKindEmbedded(ctx *KindEmbeddedContext) interface{}

	// Visit a parse tree produced by SQLMesh#kindScdType2.
	VisitKindScdType2(ctx *KindScdType2Context) interface{}

	// Visit a parse tree produced by SQLMesh#auditDef.
	VisitAuditDef(ctx *AuditDefContext) interface{}

	// Visit a parse tree produced by SQLMesh#queryProp.
	VisitQueryProp(ctx *QueryPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#auditDialectProp.
	VisitAuditDialectProp(ctx *AuditDialectPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#auditArray.
	VisitAuditArray(ctx *AuditArrayContext) interface{}

	// Visit a parse tree produced by SQLMesh#auditEntry.
	VisitAuditEntry(ctx *AuditEntryContext) interface{}

	// Visit a parse tree produced by SQLMesh#auditEntryNameProp.
	VisitAuditEntryNameProp(ctx *AuditEntryNamePropContext) interface{}

	// Visit a parse tree produced by SQLMesh#auditEntryQueryProp.
	VisitAuditEntryQueryProp(ctx *AuditEntryQueryPropContext) interface{}

	// Visit a parse tree produced by SQLMesh#propValue.
	VisitPropValue(ctx *PropValueContext) interface{}

	// Visit a parse tree produced by SQLMesh#stringLit.
	VisitStringLit(ctx *StringLitContext) interface{}

	// Visit a parse tree produced by SQLMesh#intLit.
	VisitIntLit(ctx *IntLitContext) interface{}

	// Visit a parse tree produced by SQLMesh#floatLit.
	VisitFloatLit(ctx *FloatLitContext) interface{}

	// Visit a parse tree produced by SQLMesh#trueLit.
	VisitTrueLit(ctx *TrueLitContext) interface{}

	// Visit a parse tree produced by SQLMesh#falseLit.
	VisitFalseLit(ctx *FalseLitContext) interface{}

	// Visit a parse tree produced by SQLMesh#qualifiedIdentLit.
	VisitQualifiedIdentLit(ctx *QualifiedIdentLitContext) interface{}

	// Visit a parse tree produced by SQLMesh#identLit.
	VisitIdentLit(ctx *IdentLitContext) interface{}

	// Visit a parse tree produced by SQLMesh#qualifiedIdent.
	VisitQualifiedIdent(ctx *QualifiedIdentContext) interface{}

	// Visit a parse tree produced by SQLMesh#arrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by SQLMesh#dictValue.
	VisitDictValue(ctx *DictValueContext) interface{}

	// Visit a parse tree produced by SQLMesh#dictEntry.
	VisitDictEntry(ctx *DictEntryContext) interface{}

	// Visit a parse tree produced by SQLMesh#macroBracedRef.
	VisitMacroBracedRef(ctx *MacroBracedRefContext) interface{}

	// Visit a parse tree produced by SQLMesh#macroIdentRef.
	VisitMacroIdentRef(ctx *MacroIdentRefContext) interface{}

	// Visit a parse tree produced by SQLMesh#sqlBody.
	VisitSqlBody(ctx *SqlBodyContext) interface{}
}
