parser grammar PipeLang;

options { tokenVocab=PipeLangLexer; }

// ── Top-level rule ─────────────────────────────────────────────────────────────

pipeFile
    : directive* EOF
    ;

// ── Directives ─────────────────────────────────────────────────────────────────

directive
    : TYPE          REST_OF_LINE    # type
    | NAME          REST_OF_LINE    # name
    | DESCRIPTION   REST_OF_LINE    # description
    | DESCRIPTION_ML BLOCK_LINE*    # descriptionML
    | TAGS          REST_OF_LINE    # tags
    | OWNER         REST_OF_LINE    # owner
    | DESTINATION   REST_OF_LINE    # destination
    | SCHEDULE      REST_OF_LINE    # schedule
    | UNIQUE_KEY    REST_OF_LINE    # uniqueKey
    | SOURCES       BLOCK_LINE*     # sourcesClause
    | PARAMS        BLOCK_LINE*     # paramsClause
    | PIPELINE      pipelineBlock   # pipelineClause
    ;

// ── Pipeline block ─────────────────────────────────────────────────────────────

// A pipeline block is a sequence of named PRQL nodes.
// Each node starts with a NODE_HEADER (@name:) and is followed by PRQL_LINE
// tokens that form the PRQL query body.

pipelineBlock
    : pipelineNode*
    ;

pipelineNode
    : NODE_HEADER PRQL_LINE*
    ;
