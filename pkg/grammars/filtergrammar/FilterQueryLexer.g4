lexer grammar FilterQueryLexer;

/*
 * Lexer Rules
 */

// Common punctuation / symbols
LPAREN : '(' ;
RPAREN : ')' ;
LBRACK : '[' ;
RBRACK : ']' ;
COMMA  : ','  ;

EQUALS      : '=' | '==' ;
NOT_EQUALS  : '!=' ;
NEQ         : '<>' ;       // alternate not-equals operator
LT          : '<'  ;
LE          : '<=' ;
GT          : '>'  ;
GE          : '>='  ;

// Operators that are made of multiple keywords
LIKE        : [Ll][Ii][Kk][Ee] ;
ILIKE       : [Ii][Ll][Ii][Kk][Ee] ;
BETWEEN     : [Bb][Ee][Tt][Ww][Ee][Ee][Nn] ;
EXISTS      : [Ee][Xx][Ii][Ss][Tt][Ss]? ;
REGEXP      : [Rr][Ee][Gg][Ee][Xx][Pp] ;
CONTAINS    : [Cc][Oo][Nn][Tt][Aa][Ii][Nn][Ss]? ;
IN          : [Ii][Nn] ;

// Boolean logic
NOT         : [Nn][Oo][Tt] ;
AND         : [Aa][Nn][Dd] ;
OR          : [Oo][Rr] ;

// For easy referencing in function calls
HASTOKEN    : [Hh][Aa][Ss][Tt][Oo][Kk][Ee][Nn];
HAS         : [Hh][Aa][Ss] ;
HASANY      : [Hh][Aa][Ss][Aa][Nn][Yy] ;
HASALL      : [Hh][Aa][Ss][Aa][Ll][Ll] ;

// Potential boolean constants
BOOL
    : [Tt][Rr][Uu][Ee]
    | [Ff][Aa][Ll][Ss][Ee]
    ;

fragment SIGN : [+-] ;

// Numbers: optional sign, then digits, optional fractional part,
// optional scientific notation (handy for future use)
NUMBER
    : SIGN? DIGIT+ ('.' DIGIT*)? ([eE] SIGN? DIGIT+)?    //  -10.25  42  +3.14  6.02e23
    | SIGN? '.' DIGIT+ ([eE] SIGN? DIGIT+)?              //  -.75    .5    -.5e-3
    ;

// Double/single-quoted text, capturing full text search strings, values, etc.
QUOTED_TEXT
    :  (   '"' ( ~["\\] | '\\' . )* '"'     // double-quoted
        |   '\'' ( ~['\\] | '\\' . )* '\'' // single-quoted
        )
    ;

fragment SEGMENT      : [a-zA-Z$_@{#] [a-zA-Z0-9$_@#{}:\-/]* ;
fragment EMPTY_BRACKS : '[' ']' ;
fragment OLD_JSON_BRACKS: '[' '*' ']';

KEY
    : SEGMENT ( '.' SEGMENT | EMPTY_BRACKS | OLD_JSON_BRACKS | '.' DIGIT+)*
    ;

// Ignore whitespace
WS
    : [ \t\r\n]+ -> skip
    ;

// Digits used by NUMBER
fragment DIGIT
    : [0-9]
    ;

FREETEXT : (~[ \t\r\n=()'"<>!,[\]])+ ;
