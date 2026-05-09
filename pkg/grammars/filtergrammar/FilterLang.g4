parser grammar FilterLang;

options { tokenVocab=FilterLangLexer; }

// ── Top-level ──────────────────────────────────────────────────────────────────

filterExpr : filter EOF ;

// ── Boolean expression ─────────────────────────────────────────────────────────
// Operator precedence (lowest → highest): OR < AND < NOT < atom.
// ANTLR4 resolves left-recursive alternatives top-to-bottom, so OR is lowest.

filter
    : filter OR  filter         # OrExpr
    | filter AND filter         # AndExpr
    | NOT filter                # NotExpr
    | LPAREN filter RPAREN      # GroupExpr
    | condition                 # CondExpr
    ;

// ── Conditions ─────────────────────────────────────────────────────────────────

condition
    // explicit cast: path::TypeName OP scalar  (overrides registry resolution)
    : path CAST_OP typeName op scalar                   # CastCompare

    // standard comparison
    | path op scalar                                     # Compare

    // membership
    | path IN     LPAREN scalarList RPAREN               # InList
    | path NOT IN LPAREN scalarList RPAREN               # NotInList

    // array operations (for Array attribute columns)
    | path HAS        scalar                             # HasElement
    | path HAS_ANY    LPAREN scalarList RPAREN           # HasAny
    | path HAS_ALL    LPAREN scalarList RPAREN           # HasAll

    // string operations
    | path CONTAINS     scalar                           # ContainsOp
    | path NOT CONTAINS scalar                           # NotContainsOp
    | path LIKE         scalar                           # LikeOp
    | path ILIKE        scalar                           # ILikeOp
    | path NOT LIKE     scalar                           # NotLikeOp
    | path MATCHES      scalar                           # MatchesOp

    // existence / null checks
    | path EXISTS                                        # ExistsCheck
    | path NOT EXISTS                                    # NotExistsCheck
    | path IS NULL                                       # IsNull
    | path IS NOT NULL                                   # IsNotNull
    ;

// ── Operators ──────────────────────────────────────────────────────────────────

op : OP_EQ | OP_NEQ | OP_LT | OP_GT | OP_LTE | OP_GTE ;

// ── Path ───────────────────────────────────────────────────────────────────────
// a             → static column or attribute shorthand
// a.b           → explicit attribute subpath (e.g. attributes.env)
// a.b.c         → nested attribute path

path : IDENT (DOT IDENT)* ;

// ── Type name ──────────────────────────────────────────────────────────────────
// e.g. Int64, String, Float64, Array(String) — the lexer emits this as IDENT;
// complex types like Array(String) are parsed as IDENT LPAREN IDENT RPAREN but
// the visitor reconstructs the raw text via GetTextFromInterval.

typeName : IDENT (LPAREN IDENT RPAREN)? ;

// ── Scalar values ──────────────────────────────────────────────────────────────

scalarList : scalar (COMMA scalar)* ;

scalar
    : STRING    # ScalarString
    | FLOAT     # ScalarFloat
    | INTEGER   # ScalarInt
    | TRUE      # ScalarTrue
    | FALSE     # ScalarFalse
    | NULL      # ScalarNull
    // bare word: env = prod  (type resolved from schema registry)
    | IDENT     # ScalarBareWord
    ;
