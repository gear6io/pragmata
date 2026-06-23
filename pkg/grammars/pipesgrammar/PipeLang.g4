parser grammar PipeLang;

options { tokenVocab=PipeLangLexer; }

// ── Top-level rule ─────────────────────────────────────────────────────────────

pipeFile
    : directive* EOF
    ;

// ── Directives ─────────────────────────────────────────────────────────────────

directive
    : TYPE          REST_OF_LINE    # typeDir
    | NAME          REST_OF_LINE    # nameDir
    | DESCRIPTION   REST_OF_LINE    # descriptionDir
    | DESCRIPTION_ML BLOCK_LINE*    # descriptionMLDir
    | TAGS          REST_OF_LINE    # tagsDir
    | OWNER         REST_OF_LINE    # ownerDir
    | DESTINATION   REST_OF_LINE    # destinationDir
    | SCHEDULE      REST_OF_LINE    # scheduleDir
    | UNIQUE_KEY    REST_OF_LINE    # uniqueKeyDir
    | SOURCES       BLOCK_LINE*     # sourcesDir
    | PARAMS        BLOCK_LINE*     # paramsDir
    | PIPELINE      pipelineBlock   # pipelineDir
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
