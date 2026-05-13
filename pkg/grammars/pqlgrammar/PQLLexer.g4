lexer grammar PQLLexer;

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
KW_FILTER     : F I L T E R   [ \t]+ -> pushMode(FILTER_BODY_MODE) ;
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

// ── FILTER_BODY_MODE ───────────────────────────────────────────────────────────
//
// Activated immediately after KW_FILTER.  The filter expression may span
// multiple lines; indented continuation lines (starting with whitespace) are
// part of the same filter.  A non-indented PQL keyword at column 0 ends the
// filter body and switches back to default mode.
//
// Each input line (including its newline) is emitted as a single FILTER_LINE
// token.  The PQL compiler joins all FILTER_LINE values and passes the result
// to the FilterLang parser as a single string.
//
// Keyword detection uses GetCharPositionInLine() == 0 (the ANTLR4 Go runtime
// exposes this on the lexer as p.GetCharPositionInLine()).  On keyword match
// we switch to default mode (not pop, because we pushed from default mode and
// the stack is now [DEFAULT, FILTER_BODY_MODE] — switch leaves [DEFAULT]).

mode FILTER_BODY_MODE;

FILTER_KW_FROM
    : { p.GetCharPositionInLine() == 0 }? F R O M [ \t]+
      -> type(KW_FROM), mode(DEFAULT_MODE) ;

FILTER_KW_FILTER
    : { p.GetCharPositionInLine() == 0 }? F I L T E R [ \t]+
      -> type(KW_FILTER), mode(FILTER_BODY_MODE) ;

FILTER_KW_DERIVE
    : { p.GetCharPositionInLine() == 0 }? D E R I V E [ \t]*
      -> type(KW_DERIVE), mode(DEFAULT_MODE) ;

FILTER_KW_SELECT
    : { p.GetCharPositionInLine() == 0 }? S E L E C T [ \t]*
      -> type(KW_SELECT), mode(DEFAULT_MODE) ;

FILTER_KW_GROUP
    : { p.GetCharPositionInLine() == 0 }? G R O U P [ \t]*
      -> type(KW_GROUP), mode(DEFAULT_MODE) ;

FILTER_KW_JOIN
    : { p.GetCharPositionInLine() == 0 }? J O I N [ \t]+
      -> type(KW_JOIN), mode(DEFAULT_MODE) ;

FILTER_KW_ARRAY_JOIN
    : { p.GetCharPositionInLine() == 0 }? A R R A Y '_' J O I N [ \t]+
      -> type(KW_ARRAY_JOIN), mode(DEFAULT_MODE) ;

FILTER_KW_SORT
    : { p.GetCharPositionInLine() == 0 }? S O R T [ \t]*
      -> type(KW_SORT), mode(DEFAULT_MODE) ;

FILTER_KW_TAKE
    : { p.GetCharPositionInLine() == 0 }? T A K E [ \t]+
      -> type(KW_TAKE), mode(DEFAULT_MODE) ;

FILTER_KW_SKIP
    : { p.GetCharPositionInLine() == 0 }? S K I P [ \t]+
      -> type(KW_SKIP), mode(DEFAULT_MODE) ;

FILTER_KW_WINDOW
    : { p.GetCharPositionInLine() == 0 }? W I N D O W [ \t]*
      -> type(KW_WINDOW), mode(DEFAULT_MODE) ;

// One FILTER_LINE per input line (content + newline).
// The lexer accumulates characters with -> more and emits on the newline.
FILTER_CHAR : ~[\r\n] -> more ;
FILTER_NL   : [\r\n]  -> type(FILTER_LINE) ;

// Sentinel: reserves the FILTER_LINE token type so FILTER_NL can retype itself.
// The space byte never appears here; this rule is never matched in DEFAULT_MODE.
FILTER_LINE : ' ' ;

// ── Fragment helpers ────────────────────────────────────────────────────────────

fragment DIGIT : [0-9] ;

fragment A : [Aa] ; fragment B : [Bb] ; fragment C : [Cc] ; fragment D : [Dd] ;
fragment E : [Ee] ; fragment F : [Ff] ; fragment G : [Gg] ; fragment H : [Hh] ;
fragment I : [Ii] ; fragment J : [Jj] ; fragment K : [Kk] ; fragment L : [Ll] ;
fragment M : [Mm] ; fragment N : [Nn] ; fragment O : [Oo] ; fragment P : [Pp] ;
fragment Q : [Qq] ; fragment R : [Rr] ; fragment S : [Ss] ; fragment T : [Tt] ;
fragment U : [Uu] ; fragment V : [Vv] ; fragment W : [Ww] ; fragment X : [Xx] ;
fragment Y : [Yy] ; fragment Z : [Zz] ;
