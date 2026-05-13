parser grammar FilterQuery;

options { tokenVocab=FilterQueryLexer; }

/*
 * Parser Rules
 */

query
    : expression
    EOF
    ;

// Expression with standard boolean precedence:
//    - parentheses > NOT > AND > OR
//    - consecutive expressions with no AND/OR => implicit AND
expression
    : orExpression
    ;

// OR expressions
orExpression
    : andExpression ( OR andExpression )*
    ;

// AND expressions + optional chaining with implicit AND if no OR is present
andExpression
    : unaryExpression ( AND unaryExpression | unaryExpression )*
    ;

// A unary expression handles optional NOT
unaryExpression
    : NOT? primary
    ;

// Primary constructs: grouped expressions, a comparison (key op value),
// a function call, or a full-text string
primary
    : LPAREN orExpression RPAREN
    | comparison
    | functionCall
    | fullText
    | key
    | value
    ;

/*
 * Comparison-like filters
 *
 * Includes all operators: =, !=, <>, <, <=, >, >=, [NOT] LIKE, [NOT] ILIKE,
 * [NOT] BETWEEN, [NOT] IN, [NOT] EXISTS, [NOT] REGEXP, [NOT] CONTAINS, etc.
 */
comparison
    : key EQUALS value
    | key (NOT_EQUALS | NEQ) value
    | key LT value
    | key LE value
    | key GT value
    | key GE value

    | key (LIKE | ILIKE) value
    | key NOT (LIKE | ILIKE) value

    | key BETWEEN value AND value
    | key NOT BETWEEN value AND value

    | key inClause
    | key notInClause

    | key EXISTS
    | key NOT EXISTS

    | key REGEXP value
    | key NOT REGEXP value

    | key CONTAINS value
    | key NOT CONTAINS value
    ;

// in(...) or in[...]
inClause
    : IN LPAREN valueList RPAREN
    | IN LBRACK valueList RBRACK
    | IN value
    ;

notInClause
    : NOT IN LPAREN valueList RPAREN
    | NOT IN LBRACK valueList RBRACK
    | NOT IN value
    ;

// List of values for in(...) or in[...]
valueList
    : value ( COMMA value )*
    ;

// Full-text search: a standalone quoted string is allowed as a "primary"
// e.g. `"Waiting for response" http.status_code=200`
fullText
    : QUOTED_TEXT
    | FREETEXT
    ;

/*
 * Function calls like:
 *    has(payload.user_ids, 123)
 *    hasAny(payload.user_ids, [123, 456])
 *    ...
 */
functionCall
    : (HASTOKEN | HAS | HASANY | HASALL) LPAREN functionParamList RPAREN
    ;

// Function parameters can be keys, single scalar values, or arrays
functionParamList
    : functionParam ( COMMA functionParam )*
    ;

functionParam
    : key
    | value
    | array
    ;

// An array: [ item1, item2, item3 ]
array
    : LBRACK valueList RBRACK
    ;

/*
 * A 'value' can be a string literal (double or single-quoted),
//  a numeric literal, boolean, or a "bare" token as needed.
 */
value
    : QUOTED_TEXT
    | NUMBER
    | BOOL
    | KEY
    ;

/*
 * A key can include letters, digits, underscores, dots, brackets
 * E.g. service.name, query_log.query_duration_ms, proto.user_objects[].name
 */
key
    : KEY
    ;
