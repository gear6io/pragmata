lexer grammar PRQLLexer;

// ── Whitespace & comments ──────────────────────────────────────────────────────

COMMENT : '#' ~[\r\n]* -> skip ;
WS      : [ \t]+ -> skip ;

// ── Transform keywords ─────────────────────────────────────────────────────────
// Each keyword pushes DIRECTIVE_VALUE_MODE (for the filter transform) or stays
// in default mode.  Keywords that introduce a filter body push FILTER_BODY_MODE
// so that multi-line FilterLang expressions are captured as opaque FILTER_LINE
// tokens without the FilterLang keywords colliding with PQL keywords.
//
// All other keywords simply emit their token type; the parser handles structure.

KW_FROM       : F R O M       [ \t]+ ;
KW_FILTER     : F I L T E R   [ \t]+ ;
KW_DERIVE     : D E R I V E   [ \t]* ;
KW_SELECT     : S E L E C T   [ \t]* ;
KW_GROUP      : G R O U P     [ \t]* ;
KW_AGGREGATE  : A G G R E G A T E [ \t]* ;
KW_JOIN       : J O I N       [ \t]+ ;
KW_SORT       : S O R T       [ \t]* ;
KW_TAKE       : T A K E       [ \t]+ ;
KW_SKIP       : S K I P       [ \t]+ ;
KW_WINDOW     : W I N D O W   [ \t]* ;
KW_ARRAY_JOIN : A R R A Y '_' J O I N [ \t]+ ;

// ── Join sub-keywords ──────────────────────────────────────────────────────────

KW_SIDE  : S I D E ;
KW_INNER : I N N E R ;
KW_LEFT  : L E F T ;
KW_RIGHT : R I G H T ;
KW_FULL  : F U L L ;
KW_AS    : A S ;
KW_FINAL : F I N A L ;

// ── Punctuation ────────────────────────────────────────────────────────────────

LBRACE   : '{' ;
RBRACE   : '}' ;
LPAREN   : '(' ;
RPAREN   : ')' ;
COMMA    : ',' ;
COLON    : ':' ;
DOT      : '.' ;
RANGE    : '..' ;
MINUS    : '-' ;
PLUS     : '+' ;
STAR     : '*' ;
SLASH    : '/' ;
PERCENT  : '%' ;
PIPE     : '|' ;
CAST_OP  : '::' ;
EQ       : '=' ;
NEQ      : '!=' ;
LTE      : '<=' ;
GTE      : '>=' ;
LT       : '<' ;
GT       : '>' ;

// ── Literals ───────────────────────────────────────────────────────────────────
// FLOAT before INTEGER so '1.5' is not split into INT DOT INT.

FLOAT
    : DIGIT* '.' DIGIT+
    ;

INTEGER
    : DIGIT+
    ;

STRING
    : '\'' ( ~['\\\r\n] | '\\' . )* '\''
    | '"'  ( ~["\\\r\n] | '\\' . )* '"'
    ;

// ── Identifier ─────────────────────────────────────────────────────────────────
// Listed last so all keyword rules take precedence.

IDENT : [a-zA-Z_][a-zA-Z0-9_]* ;

// ── Newlines (significant in PQL — transforms are line-oriented) ───────────────

NEWLINE : [\r\n]+ -> skip ;


// ── Fragment helpers ────────────────────────────────────────────────────────────

fragment DIGIT : [0-9] ;

fragment A : [Aa] ; fragment B : [Bb] ; fragment C : [Cc] ; fragment D : [Dd] ;
fragment E : [Ee] ; fragment F : [Ff] ; fragment G : [Gg] ; fragment H : [Hh] ;
fragment I : [Ii] ; fragment J : [Jj] ; fragment K : [Kk] ; fragment L : [Ll] ;
fragment M : [Mm] ; fragment N : [Nn] ; fragment O : [Oo] ; fragment P : [Pp] ;
fragment Q : [Qq] ; fragment R : [Rr] ; fragment S : [Ss] ; fragment T : [Tt] ;
fragment U : [Uu] ; fragment V : [Vv] ; fragment W : [Ww] ; fragment X : [Xx] ;
fragment Y : [Yy] ; fragment Z : [Zz] ;
