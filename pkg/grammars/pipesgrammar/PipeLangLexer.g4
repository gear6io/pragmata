lexer grammar PipeLangLexer;

// cleanValue trims surrounding whitespace from a VALUE token.
// VALUE already stops at '#', so comment stripping is not needed here.
@members {
func (l *PipeLangLexer) cleanValue() {
    t := l.GetText()
    s, e := 0, len(t)
    for s < e && (t[s] == ' ' || t[s] == '\t') { s++ }
    for e > s && (t[e-1] == ' ' || t[e-1] == '\t') { e-- }
    l.SetText(t[s:e])
}
}

// ── Shared fragments ──────────────────────────────────────────────────────────
fragment COLON : ':' [ \t]* ;
fragment REST  : ~[\r\n]* [\r\n] ;
fragment ID    : [a-zA-Z_][a-zA-Z0-9_]* ;

// ── DEFAULT mode ──────────────────────────────────────────────────────────────
// All keyword rules carry a col-0 predicate so they only fire at the start of
// a line. Content tokens (IDENTIFIER, COLON_TOK, etc.) are therefore safe in
// DEFAULT — they cannot conflict with keywords away from column 0.

COMMENT  : '#' ~[\r\n]* -> skip ;
WS_BLANK : [ \t]* [\r\n]+ -> skip ;
WS       : [ \t]+ -> skip ;

TYPE        : { p.GetCharPositionInLine() == 0 }? 'type'        COLON -> pushMode(VALUE_MODE) ;
NAME        : { p.GetCharPositionInLine() == 0 }? 'name'        COLON -> pushMode(VALUE_MODE) ;
TAGS        : { p.GetCharPositionInLine() == 0 }? 'tags'        COLON -> pushMode(VALUE_MODE) ;
OWNER       : { p.GetCharPositionInLine() == 0 }? 'owner'       COLON -> pushMode(VALUE_MODE) ;
DESTINATION : { p.GetCharPositionInLine() == 0 }? 'destination' COLON -> pushMode(VALUE_MODE) ;
SCHEDULE    : { p.GetCharPositionInLine() == 0 }? 'schedule'    COLON -> pushMode(VALUE_MODE) ;
UNIQUE_KEY  : { p.GetCharPositionInLine() == 0 }? 'unique_key'  COLON -> pushMode(VALUE_MODE) ;

// DESCRIPTION_ML must be declared before DESCRIPTION (maximal munch on '|').
DESCRIPTION_ML : { p.GetCharPositionInLine() == 0 }? 'description' COLON '|' REST -> pushMode(SECTION_MODE) ;
DESCRIPTION    : { p.GetCharPositionInLine() == 0 }? 'description' COLON          -> pushMode(VALUE_MODE)   ;

// SOURCES and PARAMS do not push a mode — the parser's LL structure handles
// section boundaries (source*/param* loops stop at the next col-0 keyword).
// PIPELINE pushes SECTION_MODE because its body is opaque indented text.
SOURCES  : { p.GetCharPositionInLine() == 0 }? 'sources'  COLON REST ;
PARAMS   : { p.GetCharPositionInLine() == 0 }? 'params'   COLON REST ;
PIPELINE : { p.GetCharPositionInLine() == 0 }? 'pipeline' COLON REST -> pushMode(SECTION_MODE) ;

// Content tokens — 'type'/'default' at non-col-0 fall through to IDENTIFIER.
IDENTIFIER : [a-zA-Z_][a-zA-Z0-9_./\-]* ;
COLON_TOK  : ':' ;
LBRACE     : '{' ;
RBRACE     : '}' ;
COMMA      : ',' ;
DASH       : '-' ;
STRING_LIT : '"' ~["]* '"' ;
NUMBER     : '-'? [0-9]+ ('.' [0-9]*)? ;

// ── VALUE_MODE ────────────────────────────────────────────────────────────────
// Entered after any simple-value keyword. VALUE stops at '#', so cleanValue
// only needs to trim whitespace.
// ⚠  Do not add tokens here without care — they affect every directive value.

mode VALUE_MODE;

INLINE_CMT : '#' ~[\r\n]* -> skip ;
VALUE      : ~[\r\n#]+ { l.cleanValue() } -> popMode ;
VALUE_NL   : [\r\n] -> skip, popMode ;

// ── SECTION_MODE ──────────────────────────────────────────────────────────────
// Used for pipeline blocks and multi-line description bodies.
// Exit-keyword rules are written ONCE here instead of being repeated across
// multiple modes (the root cause of the old grammar's 44-rule duplication).

mode SECTION_MODE;

SEC_TYPE        : { p.GetCharPositionInLine() == 0 }? 'type'        COLON -> type(TYPE),         mode(VALUE_MODE)    ;
SEC_NAME        : { p.GetCharPositionInLine() == 0 }? 'name'        COLON -> type(NAME),         mode(VALUE_MODE)    ;
SEC_DESC_ML     : { p.GetCharPositionInLine() == 0 }? 'description' COLON '|' REST -> type(DESCRIPTION_ML), mode(SECTION_MODE) ;
SEC_DESC        : { p.GetCharPositionInLine() == 0 }? 'description' COLON          -> type(DESCRIPTION),    mode(VALUE_MODE)   ;
SEC_TAGS        : { p.GetCharPositionInLine() == 0 }? 'tags'        COLON -> type(TAGS),         mode(VALUE_MODE)    ;
SEC_OWNER       : { p.GetCharPositionInLine() == 0 }? 'owner'       COLON -> type(OWNER),        mode(VALUE_MODE)    ;
SEC_DESTINATION : { p.GetCharPositionInLine() == 0 }? 'destination' COLON -> type(DESTINATION),  mode(VALUE_MODE)    ;
SEC_SCHEDULE    : { p.GetCharPositionInLine() == 0 }? 'schedule'    COLON -> type(SCHEDULE),     mode(VALUE_MODE)    ;
SEC_UNIQUE_KEY  : { p.GetCharPositionInLine() == 0 }? 'unique_key'  COLON -> type(UNIQUE_KEY),   mode(VALUE_MODE)    ;
SEC_SOURCES     : { p.GetCharPositionInLine() == 0 }? 'sources'     COLON REST -> type(SOURCES),  popMode ;
SEC_PARAMS      : { p.GetCharPositionInLine() == 0 }? 'params'      COLON REST -> type(PARAMS),   popMode ;
SEC_PIPELINE    : { p.GetCharPositionInLine() == 0 }? 'pipeline'    COLON REST -> type(PIPELINE), mode(SECTION_MODE) ;

SECTION_LINE    : [ \t]+ ~[\r\n]* [\r\n]? ;
SECTION_BLANK   : [ \t]* [\r\n] -> skip ;
SECTION_COMMENT : { p.GetCharPositionInLine() == 0 }? '#' ~[\r\n]* [\r\n]? -> skip ;
