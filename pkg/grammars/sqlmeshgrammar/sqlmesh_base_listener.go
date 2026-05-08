// Code generated from SQLMesh.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // SQLMesh
import "github.com/antlr4-go/antlr/v4"

// BaseSQLMeshListener is a complete listener for a parse tree produced by SQLMesh.
type BaseSQLMeshListener struct{}

var _ SQLMeshListener = &BaseSQLMeshListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLMeshListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLMeshListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLMeshListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLMeshListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSqlmeshFile is called when production sqlmeshFile is entered.
func (s *BaseSQLMeshListener) EnterSqlmeshFile(ctx *SqlmeshFileContext) {}

// ExitSqlmeshFile is called when production sqlmeshFile is exited.
func (s *BaseSQLMeshListener) ExitSqlmeshFile(ctx *SqlmeshFileContext) {}

// EnterTopStatement is called when production topStatement is entered.
func (s *BaseSQLMeshListener) EnterTopStatement(ctx *TopStatementContext) {}

// ExitTopStatement is called when production topStatement is exited.
func (s *BaseSQLMeshListener) ExitTopStatement(ctx *TopStatementContext) {}

// EnterModelDef is called when production modelDef is entered.
func (s *BaseSQLMeshListener) EnterModelDef(ctx *ModelDefContext) {}

// ExitModelDef is called when production modelDef is exited.
func (s *BaseSQLMeshListener) ExitModelDef(ctx *ModelDefContext) {}

// EnterNameProp is called when production nameProp is entered.
func (s *BaseSQLMeshListener) EnterNameProp(ctx *NamePropContext) {}

// ExitNameProp is called when production nameProp is exited.
func (s *BaseSQLMeshListener) ExitNameProp(ctx *NamePropContext) {}

// EnterKindProp is called when production kindProp is entered.
func (s *BaseSQLMeshListener) EnterKindProp(ctx *KindPropContext) {}

// ExitKindProp is called when production kindProp is exited.
func (s *BaseSQLMeshListener) ExitKindProp(ctx *KindPropContext) {}

// EnterDialectProp is called when production dialectProp is entered.
func (s *BaseSQLMeshListener) EnterDialectProp(ctx *DialectPropContext) {}

// ExitDialectProp is called when production dialectProp is exited.
func (s *BaseSQLMeshListener) ExitDialectProp(ctx *DialectPropContext) {}

// EnterOwnerProp is called when production ownerProp is entered.
func (s *BaseSQLMeshListener) EnterOwnerProp(ctx *OwnerPropContext) {}

// ExitOwnerProp is called when production ownerProp is exited.
func (s *BaseSQLMeshListener) ExitOwnerProp(ctx *OwnerPropContext) {}

// EnterCronProp is called when production cronProp is entered.
func (s *BaseSQLMeshListener) EnterCronProp(ctx *CronPropContext) {}

// ExitCronProp is called when production cronProp is exited.
func (s *BaseSQLMeshListener) ExitCronProp(ctx *CronPropContext) {}

// EnterGrainProp is called when production grainProp is entered.
func (s *BaseSQLMeshListener) EnterGrainProp(ctx *GrainPropContext) {}

// ExitGrainProp is called when production grainProp is exited.
func (s *BaseSQLMeshListener) ExitGrainProp(ctx *GrainPropContext) {}

// EnterPartitionedByProp is called when production partitionedByProp is entered.
func (s *BaseSQLMeshListener) EnterPartitionedByProp(ctx *PartitionedByPropContext) {}

// ExitPartitionedByProp is called when production partitionedByProp is exited.
func (s *BaseSQLMeshListener) ExitPartitionedByProp(ctx *PartitionedByPropContext) {}

// EnterClusteredByProp is called when production clusteredByProp is entered.
func (s *BaseSQLMeshListener) EnterClusteredByProp(ctx *ClusteredByPropContext) {}

// ExitClusteredByProp is called when production clusteredByProp is exited.
func (s *BaseSQLMeshListener) ExitClusteredByProp(ctx *ClusteredByPropContext) {}

// EnterStorageFormatProp is called when production storageFormatProp is entered.
func (s *BaseSQLMeshListener) EnterStorageFormatProp(ctx *StorageFormatPropContext) {}

// ExitStorageFormatProp is called when production storageFormatProp is exited.
func (s *BaseSQLMeshListener) ExitStorageFormatProp(ctx *StorageFormatPropContext) {}

// EnterTableFormatProp is called when production tableFormatProp is entered.
func (s *BaseSQLMeshListener) EnterTableFormatProp(ctx *TableFormatPropContext) {}

// ExitTableFormatProp is called when production tableFormatProp is exited.
func (s *BaseSQLMeshListener) ExitTableFormatProp(ctx *TableFormatPropContext) {}

// EnterRetentionProp is called when production retentionProp is entered.
func (s *BaseSQLMeshListener) EnterRetentionProp(ctx *RetentionPropContext) {}

// ExitRetentionProp is called when production retentionProp is exited.
func (s *BaseSQLMeshListener) ExitRetentionProp(ctx *RetentionPropContext) {}

// EnterTagsProp is called when production tagsProp is entered.
func (s *BaseSQLMeshListener) EnterTagsProp(ctx *TagsPropContext) {}

// ExitTagsProp is called when production tagsProp is exited.
func (s *BaseSQLMeshListener) ExitTagsProp(ctx *TagsPropContext) {}

// EnterDescriptionProp is called when production descriptionProp is entered.
func (s *BaseSQLMeshListener) EnterDescriptionProp(ctx *DescriptionPropContext) {}

// ExitDescriptionProp is called when production descriptionProp is exited.
func (s *BaseSQLMeshListener) ExitDescriptionProp(ctx *DescriptionPropContext) {}

// EnterStampProp is called when production stampProp is entered.
func (s *BaseSQLMeshListener) EnterStampProp(ctx *StampPropContext) {}

// ExitStampProp is called when production stampProp is exited.
func (s *BaseSQLMeshListener) ExitStampProp(ctx *StampPropContext) {}

// EnterColumnsProp is called when production columnsProp is entered.
func (s *BaseSQLMeshListener) EnterColumnsProp(ctx *ColumnsPropContext) {}

// ExitColumnsProp is called when production columnsProp is exited.
func (s *BaseSQLMeshListener) ExitColumnsProp(ctx *ColumnsPropContext) {}

// EnterColumnDescriptionsProp is called when production columnDescriptionsProp is entered.
func (s *BaseSQLMeshListener) EnterColumnDescriptionsProp(ctx *ColumnDescriptionsPropContext) {}

// ExitColumnDescriptionsProp is called when production columnDescriptionsProp is exited.
func (s *BaseSQLMeshListener) ExitColumnDescriptionsProp(ctx *ColumnDescriptionsPropContext) {}

// EnterPhysicalPropertiesProp is called when production physicalPropertiesProp is entered.
func (s *BaseSQLMeshListener) EnterPhysicalPropertiesProp(ctx *PhysicalPropertiesPropContext) {}

// ExitPhysicalPropertiesProp is called when production physicalPropertiesProp is exited.
func (s *BaseSQLMeshListener) ExitPhysicalPropertiesProp(ctx *PhysicalPropertiesPropContext) {}

// EnterPreStatementsProp is called when production preStatementsProp is entered.
func (s *BaseSQLMeshListener) EnterPreStatementsProp(ctx *PreStatementsPropContext) {}

// ExitPreStatementsProp is called when production preStatementsProp is exited.
func (s *BaseSQLMeshListener) ExitPreStatementsProp(ctx *PreStatementsPropContext) {}

// EnterPostStatementsProp is called when production postStatementsProp is entered.
func (s *BaseSQLMeshListener) EnterPostStatementsProp(ctx *PostStatementsPropContext) {}

// ExitPostStatementsProp is called when production postStatementsProp is exited.
func (s *BaseSQLMeshListener) ExitPostStatementsProp(ctx *PostStatementsPropContext) {}

// EnterOnVirtualUpdateProp is called when production onVirtualUpdateProp is entered.
func (s *BaseSQLMeshListener) EnterOnVirtualUpdateProp(ctx *OnVirtualUpdatePropContext) {}

// ExitOnVirtualUpdateProp is called when production onVirtualUpdateProp is exited.
func (s *BaseSQLMeshListener) ExitOnVirtualUpdateProp(ctx *OnVirtualUpdatePropContext) {}

// EnterAuditsProp is called when production auditsProp is entered.
func (s *BaseSQLMeshListener) EnterAuditsProp(ctx *AuditsPropContext) {}

// ExitAuditsProp is called when production auditsProp is exited.
func (s *BaseSQLMeshListener) ExitAuditsProp(ctx *AuditsPropContext) {}

// EnterTimeColumnProp is called when production timeColumnProp is entered.
func (s *BaseSQLMeshListener) EnterTimeColumnProp(ctx *TimeColumnPropContext) {}

// ExitTimeColumnProp is called when production timeColumnProp is exited.
func (s *BaseSQLMeshListener) ExitTimeColumnProp(ctx *TimeColumnPropContext) {}

// EnterUniqueKeyProp is called when production uniqueKeyProp is entered.
func (s *BaseSQLMeshListener) EnterUniqueKeyProp(ctx *UniqueKeyPropContext) {}

// ExitUniqueKeyProp is called when production uniqueKeyProp is exited.
func (s *BaseSQLMeshListener) ExitUniqueKeyProp(ctx *UniqueKeyPropContext) {}

// EnterInvalidateHardDeletesProp is called when production invalidateHardDeletesProp is entered.
func (s *BaseSQLMeshListener) EnterInvalidateHardDeletesProp(ctx *InvalidateHardDeletesPropContext) {}

// ExitInvalidateHardDeletesProp is called when production invalidateHardDeletesProp is exited.
func (s *BaseSQLMeshListener) ExitInvalidateHardDeletesProp(ctx *InvalidateHardDeletesPropContext) {}

// EnterDisableRestatementProp is called when production disableRestatementProp is entered.
func (s *BaseSQLMeshListener) EnterDisableRestatementProp(ctx *DisableRestatementPropContext) {}

// ExitDisableRestatementProp is called when production disableRestatementProp is exited.
func (s *BaseSQLMeshListener) ExitDisableRestatementProp(ctx *DisableRestatementPropContext) {}

// EnterPathProp is called when production pathProp is entered.
func (s *BaseSQLMeshListener) EnterPathProp(ctx *PathPropContext) {}

// ExitPathProp is called when production pathProp is exited.
func (s *BaseSQLMeshListener) ExitPathProp(ctx *PathPropContext) {}

// EnterCsvSettingsProp is called when production csvSettingsProp is entered.
func (s *BaseSQLMeshListener) EnterCsvSettingsProp(ctx *CsvSettingsPropContext) {}

// ExitCsvSettingsProp is called when production csvSettingsProp is exited.
func (s *BaseSQLMeshListener) ExitCsvSettingsProp(ctx *CsvSettingsPropContext) {}

// EnterKindFull is called when production kindFull is entered.
func (s *BaseSQLMeshListener) EnterKindFull(ctx *KindFullContext) {}

// ExitKindFull is called when production kindFull is exited.
func (s *BaseSQLMeshListener) ExitKindFull(ctx *KindFullContext) {}

// EnterKindIncrementalByTimeRange is called when production kindIncrementalByTimeRange is entered.
func (s *BaseSQLMeshListener) EnterKindIncrementalByTimeRange(ctx *KindIncrementalByTimeRangeContext) {
}

// ExitKindIncrementalByTimeRange is called when production kindIncrementalByTimeRange is exited.
func (s *BaseSQLMeshListener) ExitKindIncrementalByTimeRange(ctx *KindIncrementalByTimeRangeContext) {
}

// EnterKindIncrementalByUniqueKey is called when production kindIncrementalByUniqueKey is entered.
func (s *BaseSQLMeshListener) EnterKindIncrementalByUniqueKey(ctx *KindIncrementalByUniqueKeyContext) {
}

// ExitKindIncrementalByUniqueKey is called when production kindIncrementalByUniqueKey is exited.
func (s *BaseSQLMeshListener) ExitKindIncrementalByUniqueKey(ctx *KindIncrementalByUniqueKeyContext) {
}

// EnterKindIncrementalByPartition is called when production kindIncrementalByPartition is entered.
func (s *BaseSQLMeshListener) EnterKindIncrementalByPartition(ctx *KindIncrementalByPartitionContext) {
}

// ExitKindIncrementalByPartition is called when production kindIncrementalByPartition is exited.
func (s *BaseSQLMeshListener) ExitKindIncrementalByPartition(ctx *KindIncrementalByPartitionContext) {
}

// EnterKindView is called when production kindView is entered.
func (s *BaseSQLMeshListener) EnterKindView(ctx *KindViewContext) {}

// ExitKindView is called when production kindView is exited.
func (s *BaseSQLMeshListener) ExitKindView(ctx *KindViewContext) {}

// EnterKindSeed is called when production kindSeed is entered.
func (s *BaseSQLMeshListener) EnterKindSeed(ctx *KindSeedContext) {}

// ExitKindSeed is called when production kindSeed is exited.
func (s *BaseSQLMeshListener) ExitKindSeed(ctx *KindSeedContext) {}

// EnterKindExternal is called when production kindExternal is entered.
func (s *BaseSQLMeshListener) EnterKindExternal(ctx *KindExternalContext) {}

// ExitKindExternal is called when production kindExternal is exited.
func (s *BaseSQLMeshListener) ExitKindExternal(ctx *KindExternalContext) {}

// EnterKindEmbedded is called when production kindEmbedded is entered.
func (s *BaseSQLMeshListener) EnterKindEmbedded(ctx *KindEmbeddedContext) {}

// ExitKindEmbedded is called when production kindEmbedded is exited.
func (s *BaseSQLMeshListener) ExitKindEmbedded(ctx *KindEmbeddedContext) {}

// EnterKindScdType2 is called when production kindScdType2 is entered.
func (s *BaseSQLMeshListener) EnterKindScdType2(ctx *KindScdType2Context) {}

// ExitKindScdType2 is called when production kindScdType2 is exited.
func (s *BaseSQLMeshListener) ExitKindScdType2(ctx *KindScdType2Context) {}

// EnterAuditDef is called when production auditDef is entered.
func (s *BaseSQLMeshListener) EnterAuditDef(ctx *AuditDefContext) {}

// ExitAuditDef is called when production auditDef is exited.
func (s *BaseSQLMeshListener) ExitAuditDef(ctx *AuditDefContext) {}

// EnterQueryProp is called when production queryProp is entered.
func (s *BaseSQLMeshListener) EnterQueryProp(ctx *QueryPropContext) {}

// ExitQueryProp is called when production queryProp is exited.
func (s *BaseSQLMeshListener) ExitQueryProp(ctx *QueryPropContext) {}

// EnterAuditDialectProp is called when production auditDialectProp is entered.
func (s *BaseSQLMeshListener) EnterAuditDialectProp(ctx *AuditDialectPropContext) {}

// ExitAuditDialectProp is called when production auditDialectProp is exited.
func (s *BaseSQLMeshListener) ExitAuditDialectProp(ctx *AuditDialectPropContext) {}

// EnterAuditArray is called when production auditArray is entered.
func (s *BaseSQLMeshListener) EnterAuditArray(ctx *AuditArrayContext) {}

// ExitAuditArray is called when production auditArray is exited.
func (s *BaseSQLMeshListener) ExitAuditArray(ctx *AuditArrayContext) {}

// EnterAuditEntry is called when production auditEntry is entered.
func (s *BaseSQLMeshListener) EnterAuditEntry(ctx *AuditEntryContext) {}

// ExitAuditEntry is called when production auditEntry is exited.
func (s *BaseSQLMeshListener) ExitAuditEntry(ctx *AuditEntryContext) {}

// EnterAuditEntryNameProp is called when production auditEntryNameProp is entered.
func (s *BaseSQLMeshListener) EnterAuditEntryNameProp(ctx *AuditEntryNamePropContext) {}

// ExitAuditEntryNameProp is called when production auditEntryNameProp is exited.
func (s *BaseSQLMeshListener) ExitAuditEntryNameProp(ctx *AuditEntryNamePropContext) {}

// EnterAuditEntryQueryProp is called when production auditEntryQueryProp is entered.
func (s *BaseSQLMeshListener) EnterAuditEntryQueryProp(ctx *AuditEntryQueryPropContext) {}

// ExitAuditEntryQueryProp is called when production auditEntryQueryProp is exited.
func (s *BaseSQLMeshListener) ExitAuditEntryQueryProp(ctx *AuditEntryQueryPropContext) {}

// EnterPropValue is called when production propValue is entered.
func (s *BaseSQLMeshListener) EnterPropValue(ctx *PropValueContext) {}

// ExitPropValue is called when production propValue is exited.
func (s *BaseSQLMeshListener) ExitPropValue(ctx *PropValueContext) {}

// EnterStringLit is called when production stringLit is entered.
func (s *BaseSQLMeshListener) EnterStringLit(ctx *StringLitContext) {}

// ExitStringLit is called when production stringLit is exited.
func (s *BaseSQLMeshListener) ExitStringLit(ctx *StringLitContext) {}

// EnterIntLit is called when production intLit is entered.
func (s *BaseSQLMeshListener) EnterIntLit(ctx *IntLitContext) {}

// ExitIntLit is called when production intLit is exited.
func (s *BaseSQLMeshListener) ExitIntLit(ctx *IntLitContext) {}

// EnterFloatLit is called when production floatLit is entered.
func (s *BaseSQLMeshListener) EnterFloatLit(ctx *FloatLitContext) {}

// ExitFloatLit is called when production floatLit is exited.
func (s *BaseSQLMeshListener) ExitFloatLit(ctx *FloatLitContext) {}

// EnterTrueLit is called when production trueLit is entered.
func (s *BaseSQLMeshListener) EnterTrueLit(ctx *TrueLitContext) {}

// ExitTrueLit is called when production trueLit is exited.
func (s *BaseSQLMeshListener) ExitTrueLit(ctx *TrueLitContext) {}

// EnterFalseLit is called when production falseLit is entered.
func (s *BaseSQLMeshListener) EnterFalseLit(ctx *FalseLitContext) {}

// ExitFalseLit is called when production falseLit is exited.
func (s *BaseSQLMeshListener) ExitFalseLit(ctx *FalseLitContext) {}

// EnterQualifiedIdentLit is called when production qualifiedIdentLit is entered.
func (s *BaseSQLMeshListener) EnterQualifiedIdentLit(ctx *QualifiedIdentLitContext) {}

// ExitQualifiedIdentLit is called when production qualifiedIdentLit is exited.
func (s *BaseSQLMeshListener) ExitQualifiedIdentLit(ctx *QualifiedIdentLitContext) {}

// EnterIdentLit is called when production identLit is entered.
func (s *BaseSQLMeshListener) EnterIdentLit(ctx *IdentLitContext) {}

// ExitIdentLit is called when production identLit is exited.
func (s *BaseSQLMeshListener) ExitIdentLit(ctx *IdentLitContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *BaseSQLMeshListener) EnterQualifiedIdent(ctx *QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *BaseSQLMeshListener) ExitQualifiedIdent(ctx *QualifiedIdentContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseSQLMeshListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseSQLMeshListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterDictValue is called when production dictValue is entered.
func (s *BaseSQLMeshListener) EnterDictValue(ctx *DictValueContext) {}

// ExitDictValue is called when production dictValue is exited.
func (s *BaseSQLMeshListener) ExitDictValue(ctx *DictValueContext) {}

// EnterDictEntry is called when production dictEntry is entered.
func (s *BaseSQLMeshListener) EnterDictEntry(ctx *DictEntryContext) {}

// ExitDictEntry is called when production dictEntry is exited.
func (s *BaseSQLMeshListener) ExitDictEntry(ctx *DictEntryContext) {}

// EnterMacroBracedRef is called when production macroBracedRef is entered.
func (s *BaseSQLMeshListener) EnterMacroBracedRef(ctx *MacroBracedRefContext) {}

// ExitMacroBracedRef is called when production macroBracedRef is exited.
func (s *BaseSQLMeshListener) ExitMacroBracedRef(ctx *MacroBracedRefContext) {}

// EnterMacroIdentRef is called when production macroIdentRef is entered.
func (s *BaseSQLMeshListener) EnterMacroIdentRef(ctx *MacroIdentRefContext) {}

// ExitMacroIdentRef is called when production macroIdentRef is exited.
func (s *BaseSQLMeshListener) ExitMacroIdentRef(ctx *MacroIdentRefContext) {}

// EnterSqlBody is called when production sqlBody is entered.
func (s *BaseSQLMeshListener) EnterSqlBody(ctx *SqlBodyContext) {}

// ExitSqlBody is called when production sqlBody is exited.
func (s *BaseSQLMeshListener) ExitSqlBody(ctx *SqlBodyContext) {}
