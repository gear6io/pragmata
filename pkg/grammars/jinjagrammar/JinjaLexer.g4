lexer grammar JinjaLexer;

// ── DEFAULT_MODE ──────────────────────────────────────────────────────────────

EXPR_START    : '{{' -> pushMode(EXPR_MODE) ;
BLOCK_START   : '{%' -> pushMode(BLOCK_MODE) ;
COMMENT_START : '{#' -> pushMode(COMMENT_MODE) ;
// Raw SQL/text: any char that isn't the start of a Jinja delimiter.
TEXT          : (~[{] | '{' ~['{%#])+ ;

// ── EXPR_MODE — inside {{ ... }} ─────────────────────────────────────────────

mode EXPR_MODE;
EXPR_END   : '}}' -> popMode ;
E_NAME     : [a-zA-Z_][a-zA-Z0-9_]* ;
E_STRING   : '\'' (~['\\\r\n])* '\'' | '"' (~["\\\r\n])* '"' ;
E_NUMBER   : [0-9]+ ('.' [0-9]+)? ;
E_LPAREN   : '(' ;
E_RPAREN   : ')' ;
E_COMMA    : ',' ;
E_DOT      : '.' ;
E_PIPE     : '|' ;
E_EQ       : '==' ;
E_NEQ      : '!=' ;
E_ASSIGN   : '=' ;
E_WS       : [ \t\r\n]+ -> skip ;

// ── BLOCK_MODE — inside {% ... %} ────────────────────────────────────────────

mode BLOCK_MODE;
BLOCK_END   : '%}' -> popMode ;
KW_IF       : 'if' ;
KW_ELIF     : 'elif' ;
KW_ELSE     : 'else' ;
KW_ENDIF    : 'endif' ;
KW_FOR      : 'for' ;
KW_IN       : 'in' ;
KW_ENDFOR   : 'endfor' ;
KW_MACRO    : 'macro' ;
KW_ENDMACRO : 'endmacro' ;
KW_SET      : 'set' ;
B_NAME     : [a-zA-Z_][a-zA-Z0-9_]* ;
B_STRING   : '\'' (~['\\\r\n])* '\'' | '"' (~["\\\r\n])* '"' ;
B_ASSIGN   : '=' ;
B_LPAREN   : '(' ;
B_RPAREN   : ')' ;
B_COMMA    : ',' ;
B_WS       : [ \t\r\n]+ -> skip ;

// ── COMMENT_MODE — inside {# ... #} ─────────────────────────────────────────

mode COMMENT_MODE;
COMMENT_END  : '#}' -> popMode ;
COMMENT_BODY : (~[#] | '#' ~[}])+ ;
