lexer grammar FilterLangLexer;

// ── Whitespace ─────────────────────────────────────────────────────────────────

WS : [ \t\r\n]+ -> skip ;

// ── Operators ──────────────────────────────────────────────────────────────────
// Two-char tokens must precede their single-char prefixes to win on maximal munch.

OP_LTE  : '<=' ;
OP_GTE  : '>=' ;
OP_NEQ  : '!=' ;
OP_LT   : '<' ;
OP_GT   : '>' ;
OP_EQ   : '=' ;
CAST_OP : '::' ;

// ── Punctuation ────────────────────────────────────────────────────────────────

DOT    : '.' ;
COMMA  : ',' ;
LPAREN : '(' ;
RPAREN : ')' ;

// ── Keywords (case-insensitive) ────────────────────────────────────────────────
// Multi-word tokens first so HAS_ANY / HAS_ALL beat HAS on maximal munch.

HAS_ANY  : H A S '_' A N Y ;
HAS_ALL  : H A S '_' A L L ;
AND      : A N D ;
NOT      : N O T ;
OR       : O R ;
IN       : I N ;
HAS      : H A S ;
CONTAINS : C O N T A I N S ;
ILIKE    : I L I K E ;
LIKE     : L I K E ;
MATCHES  : M A T C H E S ;
EXISTS   : E X I S T S ;
IS       : I S ;
NULL     : N U L L ;
TRUE     : T R U E ;
FALSE    : F A L S E ;

// ── Literals ───────────────────────────────────────────────────────────────────
// FLOAT before INTEGER: '1.5' must not split into INT '1', DOT, INT '5'.
// Optional leading '-' is part of the token; FilterLang has no subtraction.

FLOAT
    : '-'? DIGIT* '.' DIGIT+
    ;

INTEGER
    : '-'? DIGIT+
    ;

STRING
    : '\'' ( ~['\\\r\n] | '\\' . )* '\''
    | '"'  ( ~["\\\r\n] | '\\' . )* '"'
    ;

// ── Identifier ─────────────────────────────────────────────────────────────────
// Listed after all keyword rules; any unrecognised word falls through to IDENT.
// Keywords cannot be used as bare-word scalars or path components — quote them.

IDENT : [a-zA-Z_][a-zA-Z0-9_]* ;

// ── Fragments ──────────────────────────────────────────────────────────────────

fragment DIGIT : [0-9] ;

fragment A : [Aa] ; fragment B : [Bb] ; fragment C : [Cc] ; fragment D : [Dd] ;
fragment E : [Ee] ; fragment F : [Ff] ; fragment G : [Gg] ; fragment H : [Hh] ;
fragment I : [Ii] ; fragment J : [Jj] ; fragment K : [Kk] ; fragment L : [Ll] ;
fragment M : [Mm] ; fragment N : [Nn] ; fragment O : [Oo] ; fragment P : [Pp] ;
fragment Q : [Qq] ; fragment R : [Rr] ; fragment S : [Ss] ; fragment T : [Tt] ;
fragment U : [Uu] ; fragment V : [Vv] ; fragment W : [Ww] ; fragment X : [Xx] ;
fragment Y : [Yy] ; fragment Z : [Zz] ;
