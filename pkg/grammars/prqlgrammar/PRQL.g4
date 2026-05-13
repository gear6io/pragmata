parser grammar PRQL;

options {
	tokenVocab = PRQLLexer;
}

// ── Top-level ──────────────────────────────────────────────────────────────────

query: clause+ EOF?;

// ── Clauses ────────────────────────────────────────────────────────────────────

clause:
	fromClause
	| filterClause
	| deriveClause
	| selectClause
	| groupClause
	| joinClause
	| arrayJoinClause
	| sortClause
	| takeClause
	| skipClause
	| windowClause;

// ── from ─────────────────────────────────────────────────────────────────────── from <table>
// [final] KW_FINAL maps to ClickHouse FINAL modifier (ReplacingMergeTree — reads latest version of
// each row by deduplicating on the primary key).

fromClause: KW_FROM IDENT KW_FINAL?;

// ── filter ───────────────────────────────────────────────────────────────────── filter
// <FilterLang expression>
// 
// The lexer captures filter content as one FILTER_LINE token per input line (including continuation
// lines that start with whitespace). The PQL compiler joins all FILTER_LINE values and delegates
// parsing to the FilterLang grammar. Multi-line filters use explicit AND / OR across lines — there
// is no implicit operator between continuation lines.
// 
// Example: filter status = 200 AND env = prod AND observed_timestamp >= {{ start |
// default('2024-01-01') }}

filterClause: KW_FILTER filterBody;

filterBody: FILTER_LINE+;

// ── derive ───────────────────────────────────────────────────────────────────── derive { name =
// <ch-expr> [, name = <ch-expr>]* } Adds computed columns to every row without grouping.

deriveClause: KW_DERIVE LBRACE assignmentList RBRACE;

// ── select ───────────────────────────────────────────────────────────────────── select { col [,
// col]* } select { alias = <ch-expr> [, alias = <ch-expr>]* } Optional projection; default is
// SELECT *.

selectClause: KW_SELECT LBRACE selectionList RBRACE;

// ── group ────────────────────────────────────────────────────────────────────── group { key [,
// key]* } ( aggregate { name = <ch-expr> [, name = <ch-expr>]* } ) GROUP BY + aggregation — always
// appear together.

groupClause:
	KW_GROUP LBRACE keyList RBRACE LPAREN KW_AGGREGATE LBRACE assignmentList RBRACE RPAREN;

// ── join ─────────────────────────────────────────────────────────────────────── join
// [side:<left|right|inner|full>] <table> (<condition>) condition is either: ==col_name shorthand
// for this.col_name = that.col_name <FilterLang expr> full condition referencing both sides

joinClause:
	KW_JOIN (KW_SIDE COLON joinSide)? IDENT LPAREN joinCond RPAREN;

joinSide: KW_LEFT | KW_RIGHT | KW_INNER | KW_FULL;

joinCond:
	EQ EQ IDENT		# SelfJoinCond
	| joinCondExpr	# ExplicitJoinCond;

// join conditions are FilterLang expressions that may reference both sides; captured as raw
// expression tokens and passed to the FilterLang compiler.
joinCondExpr: joinCondToken+;

joinCondToken:
	IDENT
	| DOT
	| INTEGER
	| FLOAT
	| STRING
	| EQ
	| NEQ
	| LT
	| GT
	| LTE
	| GTE
	| CAST_OP
	| LPAREN joinCondInner* RPAREN;

joinCondInner: joinCondToken | COMMA;

// ── array_join ───────────────────────────────────────────────────────────────── array_join <col>
// [as <alias>] ClickHouse ARRAY JOIN — unnests an array column into rows.

arrayJoinClause: KW_ARRAY_JOIN IDENT (KW_AS IDENT)?;

// ── sort ─────────────────────────────────────────────────────────────────────── sort { [-|+]col
// [, [-|+]col]* } '-' prefix → DESC. '+' or no prefix → ASC.

sortClause: KW_SORT LBRACE sortList RBRACE;

// ── take ─────────────────────────────────────────────────────────────────────── take n → LIMIT n
// take n..m → LIMIT (m-n) OFFSET n (0-indexed, consistent with PRQL)

takeClause: KW_TAKE INTEGER (RANGE INTEGER)?;

// ── skip ─────────────────────────────────────────────────────────────────────── skip n → OFFSET n
// Used together with take n when the range shorthand is not sufficient.

skipClause: KW_SKIP INTEGER;

// ── window ───────────────────────────────────────────────────────────────────── window { name =
// <ch-window-expr> [, name = <ch-window-expr>]* } ClickHouse window functions; expressions are
// opaque and passed verbatim.

windowClause: KW_WINDOW LBRACE assignmentList RBRACE;

// ── Sub-rules ──────────────────────────────────────────────────────────────────

assignmentList: assignment (COMMA assignment)*;

// name = <ch-expr> where ch-expr is an opaque ClickHouse expression. The compiler reconstructs
// original text via GetTextFromInterval on the opaqueExpr context node — whitespace in the
// expression is preserved that way even though the lexer skips WS between structural tokens.
assignment: IDENT EQ opaqueExpr;

selectionList: selectionItem (COMMA selectionItem)*;

selectionItem:
	IDENT EQ opaqueExpr	# AliasedSelection
	| opaqueExpr		# BareSelection;

keyList: keyItem (COMMA keyItem)*;

keyItem: IDENT EQ opaqueExpr # ComputedKey | IDENT # ColumnKey;

sortList: sortItem (COMMA sortItem)*;

sortItem:
	MINUS IDENT		# DescSort
	| PLUS IDENT	# AscSortExplicit
	| IDENT			# AscSort;

// ── Opaque ClickHouse expression ───────────────────────────────────────────────
// 
// opaqueExpr captures an arbitrary ClickHouse SQL expression appearing inside a PQL {} block. The
// parser rule approach (rather than a single lexer token) lets ANTLR4 handle balanced parentheses
// structurally: COMMA is only a list separator at the top level; inside LPAREN...RPAREN it is part
// of the expression (e.g. multiIf(x > 1, 'high', 'low') or quantile(0.95)(col)).
// 
// The compiler visitor uses GetTextFromInterval to reconstruct the original text including
// whitespace from the token stream's hidden channel.

opaqueExpr: opaqueToken+;

opaqueToken:
	IDENT
	| INTEGER
	| FLOAT
	| STRING
	| STAR
	| PLUS
	| MINUS
	| SLASH
	| PERCENT
	| PIPE
	| EQ
	| NEQ
	| LT
	| GT
	| LTE
	| GTE
	| CAST_OP
	| DOT
	// clause keywords are valid ClickHouse function / column names in expression position
	| KW_AS
	| KW_FINAL
	| KW_INNER
	| KW_LEFT
	| KW_RIGHT
	| KW_FULL
	| KW_SIDE
	// balanced parentheses — COMMA is permitted inside (function arguments)
	| LPAREN opaqueInner* RPAREN;

// Inside a parenthesised group, commas separate function arguments — they are part of the
// expression, not the assignmentList separator.
opaqueInner: opaqueToken | COMMA;
