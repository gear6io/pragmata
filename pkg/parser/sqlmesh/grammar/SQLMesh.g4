parser grammar SQLMesh;

options { tokenVocab=SQLMeshLexer; }

// ── Top-level ──────────────────────────────────────────────────────────────────

sqlmeshFile
    : topStatement* EOF
    ;

topStatement
    : modelDef
    | auditDef
    ;

// ── MODEL block ────────────────────────────────────────────────────────────────

modelDef
    : MODEL LPAREN (modelProp (COMMA modelProp)* COMMA?)? RPAREN sqlBody
    ;

// Each property is a labeled alternative so ANTLR generates a distinct
// VisitXxxProp visitor method for every property name.
modelProp
    : PROP_NAME                    EQUALS propValue    # nameProp
    | PROP_KIND                    EQUALS kindValue    # kindProp
    | PROP_DIALECT                 EQUALS propValue    # dialectProp
    | PROP_OWNER                   EQUALS propValue    # ownerProp
    | PROP_CRON                    EQUALS propValue    # cronProp
    | PROP_GRAIN                   EQUALS arrayValue   # grainProp
    | PROP_PARTITIONED_BY          EQUALS propValue    # partitionedByProp
    | PROP_CLUSTERED_BY            EQUALS arrayValue   # clusteredByProp
    | PROP_STORAGE_FORMAT          EQUALS propValue    # storageFormatProp
    | PROP_TABLE_FORMAT            EQUALS propValue    # tableFormatProp
    | PROP_RETENTION               EQUALS propValue    # retentionProp
    | PROP_TAGS                    EQUALS arrayValue   # tagsProp
    | PROP_DESCRIPTION             EQUALS propValue    # descriptionProp
    | PROP_STAMP                   EQUALS propValue    # stampProp
    | PROP_COLUMNS                 EQUALS dictValue    # columnsProp
    | PROP_COLUMN_DESCRIPTIONS     EQUALS dictValue    # columnDescriptionsProp
    | PROP_PHYSICAL_PROPERTIES     EQUALS dictValue    # physicalPropertiesProp
    | PROP_PRE_STATEMENTS          EQUALS arrayValue   # preStatementsProp
    | PROP_POST_STATEMENTS         EQUALS arrayValue   # postStatementsProp
    | PROP_ON_VIRTUAL_UPDATE       EQUALS arrayValue   # onVirtualUpdateProp
    | PROP_AUDITS                  EQUALS auditArray   # auditsProp
    | PROP_TIME_COLUMN             EQUALS propValue    # timeColumnProp
    | PROP_UNIQUE_KEY              EQUALS propValue    # uniqueKeyProp
    | PROP_INVALIDATE_HARD_DELETES EQUALS propValue    # invalidateHardDeletesProp
    | PROP_DISABLE_RESTATEMENT     EQUALS propValue    # disableRestatementProp
    | PROP_PATH                    EQUALS propValue    # pathProp
    | PROP_CSV_SETTINGS            EQUALS dictValue    # csvSettingsProp
    ;

// ── Model kind values ──────────────────────────────────────────────────────────

kindValue
    : KIND_FULL                       # kindFull
    | KIND_INCREMENTAL_BY_TIME_RANGE  # kindIncrementalByTimeRange
    | KIND_INCREMENTAL_BY_UNIQUE_KEY  # kindIncrementalByUniqueKey
    | KIND_INCREMENTAL_BY_PARTITION   # kindIncrementalByPartition
    | KIND_VIEW                       # kindView
    | KIND_SEED                       # kindSeed
    | KIND_EXTERNAL                   # kindExternal
    | KIND_EMBEDDED                   # kindEmbedded
    | KIND_SCD_TYPE_2                 # kindScdType2
    ;

// ── AUDIT block ────────────────────────────────────────────────────────────────
// Standalone audit definitions live in audits/ directory files.
// IDENTIFIER? allows anonymous audits (no name) as well as named ones.

auditDef
    : AUDIT IDENTIFIER? LPAREN (auditProp (COMMA auditProp)* COMMA?)? RPAREN sqlBody
    ;

auditProp
    : PROP_QUERY   EQUALS propValue   # queryProp
    | PROP_DIALECT EQUALS propValue   # auditDialectProp
    ;

// ── Inline audit array ─────────────────────────────────────────────────────────
// Used inside MODEL blocks: audits = [{name = "...", query = "..."}]

auditArray
    : LBRACKET (auditEntry (COMMA auditEntry)* COMMA?)? RBRACKET
    ;

auditEntry
    : LBRACE (auditEntryProp (COMMA auditEntryProp)* COMMA?)? RBRACE
    ;

auditEntryProp
    : PROP_NAME  EQUALS propValue   # auditEntryNameProp
    | PROP_QUERY EQUALS propValue   # auditEntryQueryProp
    ;

// ── Value types ────────────────────────────────────────────────────────────────

propValue
    : literal
    | arrayValue
    | dictValue
    | macroRef
    ;

// qualifiedIdent must be tried before bare IDENTIFIER so "schema.table" is not
// split into three tokens by the literal rule.
literal
    : STRING_LITERAL   # stringLit
    | INT_LITERAL      # intLit
    | FLOAT_LITERAL    # floatLit
    | TRUE             # trueLit
    | FALSE            # falseLit
    | qualifiedIdent   # qualifiedIdentLit
    | IDENTIFIER       # identLit
    ;

// Dotted names like "db.schema.table" used in the name property value.
qualifiedIdent
    : IDENTIFIER (DOT IDENTIFIER)+
    ;

arrayValue
    : LBRACKET (propValue (COMMA propValue)* COMMA?)? RBRACKET
    ;

dictValue
    : LBRACE (dictEntry (COMMA dictEntry)* COMMA?)? RBRACE
    ;

dictEntry
    : (STRING_LITERAL | IDENTIFIER) COLON propValue
    ;

macroRef
    : MACRO_BRACED   # macroBracedRef
    | MACRO_IDENT    # macroIdentRef
    ;

// ── SQL body ───────────────────────────────────────────────────────────────────
// SQL_LINE tokens are emitted by SQL_BODY_MODE: one token per input line.
// The converter joins lines and strips blank ones.

sqlBody
    : SQL_LINE*
    ;
