parser grammar PipeLang;

options { tokenVocab=PipeLangLexer; }

// ── Parser rules ──────────────────────────────────────────────────────────────

pipeFile
    : statement* EOF
    ;

statement
    : directive
    | nodeBlock
    ;

directive
    : KW_DESCRIPTION REST_OF_LINE    # descriptionDir
    | KW_TAGS        REST_OF_LINE    # tagsDir
    | KW_TYPE        REST_OF_LINE    # typeDir
    | KW_DATASOURCE  REST_OF_LINE    # datasourceDir
    | KW_TARGET_DATASOURCE REST_OF_LINE # targetDatasourceDir
    | KW_COPY_SCHEDULE REST_OF_LINE  # copyScheduleDir
    ;

nodeBlock
    : KW_NODE REST_OF_LINE sqlBlock
    ;

sqlBlock
    : KW_SQL_ARROW sqlBody
    ;

// SQL content is opaque: one SQL_LINE token per input line, emitted by the
// lexer's SQL_BODY_MODE.  The converter inspects the first line for the
// '%' template marker and applies dedent to the remainder.
sqlBody
    : SQL_LINE*
    ;
