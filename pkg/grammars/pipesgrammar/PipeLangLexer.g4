lexer grammar PipeLangLexer;

// ── Shared fragments ──────────────────────────────────────────────────────────

// COLON abstracts the separator so ':' is never repeated literally in keyword rules.
fragment COLON : ':' [ \t]* ;

// REST captures the remainder of a line including its newline terminator.
fragment REST  : ~[\r\n]* [\r\n] ;

// ID is a standard identifier used in node names.
fragment ID    : [a-zA-Z_][a-zA-Z0-9_]* ;

// ── DEFAULT mode ──────────────────────────────────────────────────────────────

COMMENT  : '#' ~[\r\n]* -> skip ;
WS_BLANK : [ \t]* [\r\n]+ -> skip ;
WS       : [ \t]+ -> skip ;

// Simple value directives: keyword at column 0, value on the same line.
// Each pushes VALUE_MODE; REST_OF_LINE there captures the value and pops back.
TYPE        : 'type'        COLON -> pushMode(VALUE_MODE) ;
NAME        : 'name'        COLON -> pushMode(VALUE_MODE) ;
TAGS        : 'tags'        COLON -> pushMode(VALUE_MODE) ;
OWNER       : 'owner'       COLON -> pushMode(VALUE_MODE) ;
DESTINATION : 'destination' COLON -> pushMode(VALUE_MODE) ;
SCHEDULE    : 'schedule'    COLON -> pushMode(VALUE_MODE) ;

// Description has two forms.  DESCRIPTION_ML must be declared first so that
// ANTLR's maximal-munch rule picks it when '|' follows the colon.
DESCRIPTION_ML : 'description' COLON '|' REST -> pushMode(BLOCK_MODE) ;
DESCRIPTION    : 'description' COLON          -> pushMode(VALUE_MODE) ;

// Block sections: the entire header line (including any comment) is consumed
// before the mode is pushed, so the block mode starts at column 0 of the
// first indented content line.
SOURCES  : 'sources'  COLON REST -> pushMode(BLOCK_MODE) ;
PARAMS   : 'params'   COLON REST -> pushMode(BLOCK_MODE) ;
PIPELINE : 'pipeline' COLON REST -> pushMode(PIPELINE_MODE) ;

// ── VALUE_MODE ────────────────────────────────────────────────────────────────
//
// Active after any simple-value keyword.  REST_OF_LINE captures everything
// through the newline and pops back to DEFAULT (or, when called from BLOCK_MODE
// via mode(VALUE_MODE), back to DEFAULT via the mode stack).

mode VALUE_MODE;

REST_OF_LINE : ~[\r\n]* [\r\n] -> popMode ;

// ── BLOCK_MODE ────────────────────────────────────────────────────────────────
//
// Used for sources:, params:, and multi-line description: | blocks.
//
// Col-0 keyword re-emitters follow the same pattern as the existing SQL_KW_*
// rules: when a top-level keyword is recognised at column 0, its token type is
// reassigned and the mode is switched so VALUE_MODE can recapture the value
// (for simple directives) or BLOCK_MODE / PIPELINE_MODE can start fresh.
//
// mode(X) replaces the current mode stack entry, keeping DEFAULT beneath it, so
// that REST_OF_LINE -> popMode returns correctly to DEFAULT.

mode BLOCK_MODE;

BLOCK_TYPE
    : { p.GetCharPositionInLine() == 0 }? 'type'        COLON
      -> type(TYPE), mode(VALUE_MODE) ;

BLOCK_NAME
    : { p.GetCharPositionInLine() == 0 }? 'name'        COLON
      -> type(NAME), mode(VALUE_MODE) ;

BLOCK_DESC_ML
    : { p.GetCharPositionInLine() == 0 }? 'description' COLON '|' REST
      -> type(DESCRIPTION_ML), mode(BLOCK_MODE) ;

BLOCK_DESC
    : { p.GetCharPositionInLine() == 0 }? 'description' COLON
      -> type(DESCRIPTION), mode(VALUE_MODE) ;

BLOCK_TAGS
    : { p.GetCharPositionInLine() == 0 }? 'tags'        COLON
      -> type(TAGS), mode(VALUE_MODE) ;

BLOCK_OWNER
    : { p.GetCharPositionInLine() == 0 }? 'owner'       COLON
      -> type(OWNER), mode(VALUE_MODE) ;

BLOCK_DESTINATION
    : { p.GetCharPositionInLine() == 0 }? 'destination' COLON
      -> type(DESTINATION), mode(VALUE_MODE) ;

BLOCK_SCHEDULE
    : { p.GetCharPositionInLine() == 0 }? 'schedule'    COLON
      -> type(SCHEDULE), mode(VALUE_MODE) ;

BLOCK_SOURCES
    : { p.GetCharPositionInLine() == 0 }? 'sources'     COLON REST
      -> type(SOURCES), mode(BLOCK_MODE) ;

BLOCK_PARAMS
    : { p.GetCharPositionInLine() == 0 }? 'params'      COLON REST
      -> type(PARAMS), mode(BLOCK_MODE) ;

BLOCK_PIPELINE
    : { p.GetCharPositionInLine() == 0 }? 'pipeline'    COLON REST
      -> type(PIPELINE), mode(PIPELINE_MODE) ;

// Indented content lines (the actual block data — sources list items, param
// definitions, or description continuation lines).
BLOCK_LINE : [ \t]+ ~[\r\n]* [\r\n] ;

// Blank lines and column-0 comments are skipped within blocks.
BLOCK_BLANK   : [ \t]* [\r\n] -> skip ;
BLOCK_COMMENT : { p.GetCharPositionInLine() == 0 }? '#' ~[\r\n]* [\r\n]? -> skip ;

// ── PIPELINE_MODE ─────────────────────────────────────────────────────────────
//
// Used for the pipeline: section.  Like BLOCK_MODE, col-0 keywords end the
// section and re-emit with the correct type.  Within the section, @name: lines
// are NODE_HEADER tokens and all other indented lines are PRQL_LINE tokens.
// NODE_HEADER is declared before PRQL_LINE so it wins when both could match.

mode PIPELINE_MODE;

PIPE_TYPE
    : { p.GetCharPositionInLine() == 0 }? 'type'        COLON
      -> type(TYPE), mode(VALUE_MODE) ;

PIPE_NAME
    : { p.GetCharPositionInLine() == 0 }? 'name'        COLON
      -> type(NAME), mode(VALUE_MODE) ;

PIPE_DESC_ML
    : { p.GetCharPositionInLine() == 0 }? 'description' COLON '|' REST
      -> type(DESCRIPTION_ML), mode(BLOCK_MODE) ;

PIPE_DESC
    : { p.GetCharPositionInLine() == 0 }? 'description' COLON
      -> type(DESCRIPTION), mode(VALUE_MODE) ;

PIPE_TAGS
    : { p.GetCharPositionInLine() == 0 }? 'tags'        COLON
      -> type(TAGS), mode(VALUE_MODE) ;

PIPE_OWNER
    : { p.GetCharPositionInLine() == 0 }? 'owner'       COLON
      -> type(OWNER), mode(VALUE_MODE) ;

PIPE_DESTINATION
    : { p.GetCharPositionInLine() == 0 }? 'destination' COLON
      -> type(DESTINATION), mode(VALUE_MODE) ;

PIPE_SCHEDULE
    : { p.GetCharPositionInLine() == 0 }? 'schedule'    COLON
      -> type(SCHEDULE), mode(VALUE_MODE) ;

PIPE_SOURCES
    : { p.GetCharPositionInLine() == 0 }? 'sources'     COLON REST
      -> type(SOURCES), mode(BLOCK_MODE) ;

PIPE_PARAMS
    : { p.GetCharPositionInLine() == 0 }? 'params'      COLON REST
      -> type(PARAMS), mode(BLOCK_MODE) ;

PIPE_PIPELINE
    : { p.GetCharPositionInLine() == 0 }? 'pipeline'    COLON REST
      -> type(PIPELINE), mode(PIPELINE_MODE) ;

// @name: node header — must come before PRQL_LINE (both start with [ \t]+).
NODE_HEADER : [ \t]+ '@' ID COLON ~[\r\n]* [\r\n]? ;

// PRQL body lines: any other indented content.
PRQL_LINE : [ \t]+ ~[\r\n]+ [\r\n]? ;

PIPE_BLANK   : [ \t]* [\r\n] -> skip ;
PIPE_COMMENT : { p.GetCharPositionInLine() == 0 }? '#' ~[\r\n]* [\r\n]? -> skip ;
