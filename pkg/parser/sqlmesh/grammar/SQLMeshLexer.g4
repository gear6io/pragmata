lexer grammar SQLMeshLexer;

// Tracks nesting depth inside MODEL/AUDIT parens.
// When depth returns to 0 on RPAREN, the lexer pushes SQL_BODY_MODE so the
// SQL query that follows is captured as opaque SQL_LINE tokens.
@lexer::members {
    int _depth = 0;
}

// ── Comments & whitespace ──────────────────────────────────────────────────────

BLOCK_COMMENT : '/*' .*? '*/' -> skip ;
LINE_COMMENT  : '--' ~[\r\n]* -> skip ;
WS            : [ \t\r\n]+    -> skip ;

// ── Top-level block keywords ───────────────────────────────────────────────────

MODEL : 'MODEL' ;
AUDIT : 'AUDIT' ;

// ── Model kind values ──────────────────────────────────────────────────────────
// Must be listed before IDENTIFIER so the longer-match rule wins for
// identifiers that are exact kind names.

KIND_INCREMENTAL_BY_TIME_RANGE : 'INCREMENTAL_BY_TIME_RANGE' ;
KIND_INCREMENTAL_BY_UNIQUE_KEY : 'INCREMENTAL_BY_UNIQUE_KEY' ;
KIND_INCREMENTAL_BY_PARTITION  : 'INCREMENTAL_BY_PARTITION' ;
KIND_SCD_TYPE_2                : 'SCD_TYPE_2' ;
KIND_EXTERNAL                  : 'EXTERNAL' ;
KIND_EMBEDDED                  : 'EMBEDDED' ;
KIND_FULL                      : 'FULL' ;
KIND_VIEW                      : 'VIEW' ;
KIND_SEED                      : 'SEED' ;

// ── Boolean literals ───────────────────────────────────────────────────────────

TRUE  : 'TRUE'  | 'true'  ;
FALSE : 'FALSE' | 'false' ;

// ── Property name keywords ─────────────────────────────────────────────────────
// Longer names first: ANTLR resolves ambiguity by rule order when two rules
// could match the same prefix.

PROP_COLUMN_DESCRIPTIONS     : 'column_descriptions' ;
PROP_PHYSICAL_PROPERTIES     : 'physical_properties' ;
PROP_ON_VIRTUAL_UPDATE       : 'on_virtual_update' ;
PROP_INVALIDATE_HARD_DELETES : 'invalidate_hard_deletes' ;
PROP_DISABLE_RESTATEMENT     : 'disable_restatement' ;
PROP_PRE_STATEMENTS          : 'pre_statements' ;
PROP_POST_STATEMENTS         : 'post_statements' ;
PROP_PARTITIONED_BY          : 'partitioned_by' ;
PROP_CLUSTERED_BY            : 'clustered_by' ;
PROP_STORAGE_FORMAT          : 'storage_format' ;
PROP_TABLE_FORMAT            : 'table_format' ;
PROP_CSV_SETTINGS            : 'csv_settings' ;
PROP_UNIQUE_KEY              : 'unique_key' ;
PROP_TIME_COLUMN             : 'time_column' ;
PROP_DESCRIPTION             : 'description' ;
PROP_RETENTION               : 'retention' ;
PROP_DIALECT                 : 'dialect' ;
PROP_COLUMNS                 : 'columns' ;
PROP_AUDITS                  : 'audits' ;
PROP_GRAIN                   : 'grain' ;
PROP_OWNER                   : 'owner' ;
PROP_STAMP                   : 'stamp' ;
PROP_TAGS                    : 'tags' ;
PROP_CRON                    : 'cron' ;
PROP_NAME                    : 'name' ;
PROP_KIND                    : 'kind' ;
PROP_PATH                    : 'path' ;
PROP_QUERY                   : 'query' ;

// ── Punctuation ────────────────────────────────────────────────────────────────
// LPAREN/RPAREN maintain the depth counter.  When RPAREN brings depth to 0 the
// MODEL/AUDIT block has closed and the SQL body follows.

LPAREN : '(' { p._depth++; } ;
RPAREN : ')' {
    p._depth--;
    if p._depth == 0 {
        p.PushMode(SQLMeshLexerSQL_BODY_MODE)
    }
} ;
LBRACKET : '[' ;
RBRACKET : ']' ;
LBRACE   : '{' ;
RBRACE   : '}' ;
COMMA    : ',' ;
EQUALS   : '=' ;
COLON    : ':' ;
DOT      : '.' ;
SEMI     : ';' ;

// ── Macro references ───────────────────────────────────────────────────────────
// @{ident} must be tried before @ident so the braced form wins on maximal munch.

MACRO_BRACED : '@' '{' [a-zA-Z_] [a-zA-Z_0-9]* '}' ;
MACRO_IDENT  : '@'     [a-zA-Z_] [a-zA-Z_0-9]*     ;

// ── Numeric literals ───────────────────────────────────────────────────────────
// FLOAT before INT: '1.5' must not be lexed as INT '1', DOT, INT '5'.

FLOAT_LITERAL : [0-9]* '.' [0-9]+ ;
INT_LITERAL   : [0-9]+              ;

// ── String literals ────────────────────────────────────────────────────────────
// Triple-quoted variants must precede single-character-delimited variants so
// that '"""' beats '"' on maximal-munch.

STRING_LITERAL
    : '"""'    .*?    '"""'
    | '\'\'\'' .*? '\'\'\''
    | '\''  (~['\\] | '\\' .)* '\''
    | '"'   (~["\\] | '\\' .)* '"'
    ;

// ── General identifier ─────────────────────────────────────────────────────────
// Listed after all keyword rules; any unrecognised word falls through here.

IDENTIFIER : [a-zA-Z_] [a-zA-Z_0-9]* ;

// Sentinel: declares the SQL_LINE token type so SQL_BODY_MODE can re-type
// SQL_NL as SQL_LINE.  The null byte never appears in real input, so this
// rule is never matched in DEFAULT_MODE.
SQL_LINE : '\x00' ;

// ── SQL_BODY_MODE ──────────────────────────────────────────────────────────────
//
// Activated when the closing ')' of the MODEL/AUDIT block is consumed (depth
// returns to 0).  Every non-newline character is accumulated via `-> more`
// without emitting a token.  When a newline is reached, SQL_NL emits a
// SQL_LINE token containing the entire buffered line text including the newline.
//
// The optional ';' that immediately follows the closing ')' is discarded.
// Blank lines produce empty SQL_LINE tokens; the converter filters them.

mode SQL_BODY_MODE;

SQL_BODY_SEMI : ';'     -> skip ;
SQL_CHAR      : ~[\r\n] -> more ;
SQL_NL        : [\r\n]  -> type(SQL_LINE) ;
