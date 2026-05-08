// Code generated from SQLMesh.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // SQLMesh
import "github.com/antlr4-go/antlr/v4"

// SQLMeshListener is a complete listener for a parse tree produced by SQLMesh.
type SQLMeshListener interface {
	antlr.ParseTreeListener

	// EnterSqlmeshFile is called when entering the sqlmeshFile production.
	EnterSqlmeshFile(c *SqlmeshFileContext)

	// EnterTopStatement is called when entering the topStatement production.
	EnterTopStatement(c *TopStatementContext)

	// EnterModelDef is called when entering the modelDef production.
	EnterModelDef(c *ModelDefContext)

	// EnterNameProp is called when entering the nameProp production.
	EnterNameProp(c *NamePropContext)

	// EnterKindProp is called when entering the kindProp production.
	EnterKindProp(c *KindPropContext)

	// EnterDialectProp is called when entering the dialectProp production.
	EnterDialectProp(c *DialectPropContext)

	// EnterOwnerProp is called when entering the ownerProp production.
	EnterOwnerProp(c *OwnerPropContext)

	// EnterCronProp is called when entering the cronProp production.
	EnterCronProp(c *CronPropContext)

	// EnterGrainProp is called when entering the grainProp production.
	EnterGrainProp(c *GrainPropContext)

	// EnterPartitionedByProp is called when entering the partitionedByProp production.
	EnterPartitionedByProp(c *PartitionedByPropContext)

	// EnterClusteredByProp is called when entering the clusteredByProp production.
	EnterClusteredByProp(c *ClusteredByPropContext)

	// EnterStorageFormatProp is called when entering the storageFormatProp production.
	EnterStorageFormatProp(c *StorageFormatPropContext)

	// EnterTableFormatProp is called when entering the tableFormatProp production.
	EnterTableFormatProp(c *TableFormatPropContext)

	// EnterRetentionProp is called when entering the retentionProp production.
	EnterRetentionProp(c *RetentionPropContext)

	// EnterTagsProp is called when entering the tagsProp production.
	EnterTagsProp(c *TagsPropContext)

	// EnterDescriptionProp is called when entering the descriptionProp production.
	EnterDescriptionProp(c *DescriptionPropContext)

	// EnterStampProp is called when entering the stampProp production.
	EnterStampProp(c *StampPropContext)

	// EnterColumnsProp is called when entering the columnsProp production.
	EnterColumnsProp(c *ColumnsPropContext)

	// EnterColumnDescriptionsProp is called when entering the columnDescriptionsProp production.
	EnterColumnDescriptionsProp(c *ColumnDescriptionsPropContext)

	// EnterPhysicalPropertiesProp is called when entering the physicalPropertiesProp production.
	EnterPhysicalPropertiesProp(c *PhysicalPropertiesPropContext)

	// EnterPreStatementsProp is called when entering the preStatementsProp production.
	EnterPreStatementsProp(c *PreStatementsPropContext)

	// EnterPostStatementsProp is called when entering the postStatementsProp production.
	EnterPostStatementsProp(c *PostStatementsPropContext)

	// EnterOnVirtualUpdateProp is called when entering the onVirtualUpdateProp production.
	EnterOnVirtualUpdateProp(c *OnVirtualUpdatePropContext)

	// EnterAuditsProp is called when entering the auditsProp production.
	EnterAuditsProp(c *AuditsPropContext)

	// EnterTimeColumnProp is called when entering the timeColumnProp production.
	EnterTimeColumnProp(c *TimeColumnPropContext)

	// EnterUniqueKeyProp is called when entering the uniqueKeyProp production.
	EnterUniqueKeyProp(c *UniqueKeyPropContext)

	// EnterInvalidateHardDeletesProp is called when entering the invalidateHardDeletesProp production.
	EnterInvalidateHardDeletesProp(c *InvalidateHardDeletesPropContext)

	// EnterDisableRestatementProp is called when entering the disableRestatementProp production.
	EnterDisableRestatementProp(c *DisableRestatementPropContext)

	// EnterPathProp is called when entering the pathProp production.
	EnterPathProp(c *PathPropContext)

	// EnterCsvSettingsProp is called when entering the csvSettingsProp production.
	EnterCsvSettingsProp(c *CsvSettingsPropContext)

	// EnterKindFull is called when entering the kindFull production.
	EnterKindFull(c *KindFullContext)

	// EnterKindIncrementalByTimeRange is called when entering the kindIncrementalByTimeRange production.
	EnterKindIncrementalByTimeRange(c *KindIncrementalByTimeRangeContext)

	// EnterKindIncrementalByUniqueKey is called when entering the kindIncrementalByUniqueKey production.
	EnterKindIncrementalByUniqueKey(c *KindIncrementalByUniqueKeyContext)

	// EnterKindIncrementalByPartition is called when entering the kindIncrementalByPartition production.
	EnterKindIncrementalByPartition(c *KindIncrementalByPartitionContext)

	// EnterKindView is called when entering the kindView production.
	EnterKindView(c *KindViewContext)

	// EnterKindSeed is called when entering the kindSeed production.
	EnterKindSeed(c *KindSeedContext)

	// EnterKindExternal is called when entering the kindExternal production.
	EnterKindExternal(c *KindExternalContext)

	// EnterKindEmbedded is called when entering the kindEmbedded production.
	EnterKindEmbedded(c *KindEmbeddedContext)

	// EnterKindScdType2 is called when entering the kindScdType2 production.
	EnterKindScdType2(c *KindScdType2Context)

	// EnterAuditDef is called when entering the auditDef production.
	EnterAuditDef(c *AuditDefContext)

	// EnterQueryProp is called when entering the queryProp production.
	EnterQueryProp(c *QueryPropContext)

	// EnterAuditDialectProp is called when entering the auditDialectProp production.
	EnterAuditDialectProp(c *AuditDialectPropContext)

	// EnterAuditArray is called when entering the auditArray production.
	EnterAuditArray(c *AuditArrayContext)

	// EnterAuditEntry is called when entering the auditEntry production.
	EnterAuditEntry(c *AuditEntryContext)

	// EnterAuditEntryNameProp is called when entering the auditEntryNameProp production.
	EnterAuditEntryNameProp(c *AuditEntryNamePropContext)

	// EnterAuditEntryQueryProp is called when entering the auditEntryQueryProp production.
	EnterAuditEntryQueryProp(c *AuditEntryQueryPropContext)

	// EnterPropValue is called when entering the propValue production.
	EnterPropValue(c *PropValueContext)

	// EnterStringLit is called when entering the stringLit production.
	EnterStringLit(c *StringLitContext)

	// EnterIntLit is called when entering the intLit production.
	EnterIntLit(c *IntLitContext)

	// EnterFloatLit is called when entering the floatLit production.
	EnterFloatLit(c *FloatLitContext)

	// EnterTrueLit is called when entering the trueLit production.
	EnterTrueLit(c *TrueLitContext)

	// EnterFalseLit is called when entering the falseLit production.
	EnterFalseLit(c *FalseLitContext)

	// EnterQualifiedIdentLit is called when entering the qualifiedIdentLit production.
	EnterQualifiedIdentLit(c *QualifiedIdentLitContext)

	// EnterIdentLit is called when entering the identLit production.
	EnterIdentLit(c *IdentLitContext)

	// EnterQualifiedIdent is called when entering the qualifiedIdent production.
	EnterQualifiedIdent(c *QualifiedIdentContext)

	// EnterArrayValue is called when entering the arrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterDictValue is called when entering the dictValue production.
	EnterDictValue(c *DictValueContext)

	// EnterDictEntry is called when entering the dictEntry production.
	EnterDictEntry(c *DictEntryContext)

	// EnterMacroBracedRef is called when entering the macroBracedRef production.
	EnterMacroBracedRef(c *MacroBracedRefContext)

	// EnterMacroIdentRef is called when entering the macroIdentRef production.
	EnterMacroIdentRef(c *MacroIdentRefContext)

	// EnterSqlBody is called when entering the sqlBody production.
	EnterSqlBody(c *SqlBodyContext)

	// ExitSqlmeshFile is called when exiting the sqlmeshFile production.
	ExitSqlmeshFile(c *SqlmeshFileContext)

	// ExitTopStatement is called when exiting the topStatement production.
	ExitTopStatement(c *TopStatementContext)

	// ExitModelDef is called when exiting the modelDef production.
	ExitModelDef(c *ModelDefContext)

	// ExitNameProp is called when exiting the nameProp production.
	ExitNameProp(c *NamePropContext)

	// ExitKindProp is called when exiting the kindProp production.
	ExitKindProp(c *KindPropContext)

	// ExitDialectProp is called when exiting the dialectProp production.
	ExitDialectProp(c *DialectPropContext)

	// ExitOwnerProp is called when exiting the ownerProp production.
	ExitOwnerProp(c *OwnerPropContext)

	// ExitCronProp is called when exiting the cronProp production.
	ExitCronProp(c *CronPropContext)

	// ExitGrainProp is called when exiting the grainProp production.
	ExitGrainProp(c *GrainPropContext)

	// ExitPartitionedByProp is called when exiting the partitionedByProp production.
	ExitPartitionedByProp(c *PartitionedByPropContext)

	// ExitClusteredByProp is called when exiting the clusteredByProp production.
	ExitClusteredByProp(c *ClusteredByPropContext)

	// ExitStorageFormatProp is called when exiting the storageFormatProp production.
	ExitStorageFormatProp(c *StorageFormatPropContext)

	// ExitTableFormatProp is called when exiting the tableFormatProp production.
	ExitTableFormatProp(c *TableFormatPropContext)

	// ExitRetentionProp is called when exiting the retentionProp production.
	ExitRetentionProp(c *RetentionPropContext)

	// ExitTagsProp is called when exiting the tagsProp production.
	ExitTagsProp(c *TagsPropContext)

	// ExitDescriptionProp is called when exiting the descriptionProp production.
	ExitDescriptionProp(c *DescriptionPropContext)

	// ExitStampProp is called when exiting the stampProp production.
	ExitStampProp(c *StampPropContext)

	// ExitColumnsProp is called when exiting the columnsProp production.
	ExitColumnsProp(c *ColumnsPropContext)

	// ExitColumnDescriptionsProp is called when exiting the columnDescriptionsProp production.
	ExitColumnDescriptionsProp(c *ColumnDescriptionsPropContext)

	// ExitPhysicalPropertiesProp is called when exiting the physicalPropertiesProp production.
	ExitPhysicalPropertiesProp(c *PhysicalPropertiesPropContext)

	// ExitPreStatementsProp is called when exiting the preStatementsProp production.
	ExitPreStatementsProp(c *PreStatementsPropContext)

	// ExitPostStatementsProp is called when exiting the postStatementsProp production.
	ExitPostStatementsProp(c *PostStatementsPropContext)

	// ExitOnVirtualUpdateProp is called when exiting the onVirtualUpdateProp production.
	ExitOnVirtualUpdateProp(c *OnVirtualUpdatePropContext)

	// ExitAuditsProp is called when exiting the auditsProp production.
	ExitAuditsProp(c *AuditsPropContext)

	// ExitTimeColumnProp is called when exiting the timeColumnProp production.
	ExitTimeColumnProp(c *TimeColumnPropContext)

	// ExitUniqueKeyProp is called when exiting the uniqueKeyProp production.
	ExitUniqueKeyProp(c *UniqueKeyPropContext)

	// ExitInvalidateHardDeletesProp is called when exiting the invalidateHardDeletesProp production.
	ExitInvalidateHardDeletesProp(c *InvalidateHardDeletesPropContext)

	// ExitDisableRestatementProp is called when exiting the disableRestatementProp production.
	ExitDisableRestatementProp(c *DisableRestatementPropContext)

	// ExitPathProp is called when exiting the pathProp production.
	ExitPathProp(c *PathPropContext)

	// ExitCsvSettingsProp is called when exiting the csvSettingsProp production.
	ExitCsvSettingsProp(c *CsvSettingsPropContext)

	// ExitKindFull is called when exiting the kindFull production.
	ExitKindFull(c *KindFullContext)

	// ExitKindIncrementalByTimeRange is called when exiting the kindIncrementalByTimeRange production.
	ExitKindIncrementalByTimeRange(c *KindIncrementalByTimeRangeContext)

	// ExitKindIncrementalByUniqueKey is called when exiting the kindIncrementalByUniqueKey production.
	ExitKindIncrementalByUniqueKey(c *KindIncrementalByUniqueKeyContext)

	// ExitKindIncrementalByPartition is called when exiting the kindIncrementalByPartition production.
	ExitKindIncrementalByPartition(c *KindIncrementalByPartitionContext)

	// ExitKindView is called when exiting the kindView production.
	ExitKindView(c *KindViewContext)

	// ExitKindSeed is called when exiting the kindSeed production.
	ExitKindSeed(c *KindSeedContext)

	// ExitKindExternal is called when exiting the kindExternal production.
	ExitKindExternal(c *KindExternalContext)

	// ExitKindEmbedded is called when exiting the kindEmbedded production.
	ExitKindEmbedded(c *KindEmbeddedContext)

	// ExitKindScdType2 is called when exiting the kindScdType2 production.
	ExitKindScdType2(c *KindScdType2Context)

	// ExitAuditDef is called when exiting the auditDef production.
	ExitAuditDef(c *AuditDefContext)

	// ExitQueryProp is called when exiting the queryProp production.
	ExitQueryProp(c *QueryPropContext)

	// ExitAuditDialectProp is called when exiting the auditDialectProp production.
	ExitAuditDialectProp(c *AuditDialectPropContext)

	// ExitAuditArray is called when exiting the auditArray production.
	ExitAuditArray(c *AuditArrayContext)

	// ExitAuditEntry is called when exiting the auditEntry production.
	ExitAuditEntry(c *AuditEntryContext)

	// ExitAuditEntryNameProp is called when exiting the auditEntryNameProp production.
	ExitAuditEntryNameProp(c *AuditEntryNamePropContext)

	// ExitAuditEntryQueryProp is called when exiting the auditEntryQueryProp production.
	ExitAuditEntryQueryProp(c *AuditEntryQueryPropContext)

	// ExitPropValue is called when exiting the propValue production.
	ExitPropValue(c *PropValueContext)

	// ExitStringLit is called when exiting the stringLit production.
	ExitStringLit(c *StringLitContext)

	// ExitIntLit is called when exiting the intLit production.
	ExitIntLit(c *IntLitContext)

	// ExitFloatLit is called when exiting the floatLit production.
	ExitFloatLit(c *FloatLitContext)

	// ExitTrueLit is called when exiting the trueLit production.
	ExitTrueLit(c *TrueLitContext)

	// ExitFalseLit is called when exiting the falseLit production.
	ExitFalseLit(c *FalseLitContext)

	// ExitQualifiedIdentLit is called when exiting the qualifiedIdentLit production.
	ExitQualifiedIdentLit(c *QualifiedIdentLitContext)

	// ExitIdentLit is called when exiting the identLit production.
	ExitIdentLit(c *IdentLitContext)

	// ExitQualifiedIdent is called when exiting the qualifiedIdent production.
	ExitQualifiedIdent(c *QualifiedIdentContext)

	// ExitArrayValue is called when exiting the arrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitDictValue is called when exiting the dictValue production.
	ExitDictValue(c *DictValueContext)

	// ExitDictEntry is called when exiting the dictEntry production.
	ExitDictEntry(c *DictEntryContext)

	// ExitMacroBracedRef is called when exiting the macroBracedRef production.
	ExitMacroBracedRef(c *MacroBracedRefContext)

	// ExitMacroIdentRef is called when exiting the macroIdentRef production.
	ExitMacroIdentRef(c *MacroIdentRefContext)

	// ExitSqlBody is called when exiting the sqlBody production.
	ExitSqlBody(c *SqlBodyContext)
}
