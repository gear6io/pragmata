parser grammar Jinja;

options { tokenVocab=JinjaLexer; }

template : element* EOF ;

element
    : TEXT                                          # rawText
    | EXPR_START expr EXPR_END                      # exprTag
    | ifBlock                                       # ifTag
    | forBlock                                      # forTag
    | macroBlock                                    # macroTag
    | setStmt                                       # setTag
    | COMMENT_START COMMENT_BODY* COMMENT_END       # commentTag
    ;

// ── Expressions — {{ expr }} ──────────────────────────────────────────────────
//
// Precedence (lowest to highest): filter chain > function call > field access > primary.
// ANTLR resolves left-recursive alternatives via operator precedence climbing.

expr
    : primary                                                      # primaryExpr
    | expr E_DOT E_NAME                                            # fieldAccess
    | expr E_LPAREN argList? E_RPAREN                              # funcCall
    | expr E_PIPE E_NAME (E_LPAREN argList? E_RPAREN)?             # filterExpr
    ;

primary
    : E_NAME                     # nameExpr
    | E_STRING                   # stringExpr
    | E_NUMBER                   # numberExpr
    | E_LPAREN expr E_RPAREN     # parenExpr
    ;

argList : arg (E_COMMA arg)* ;

arg
    : E_NAME E_ASSIGN expr   # kwarg
    | expr                   # posarg
    ;

// ── Control flow — {% if %} ───────────────────────────────────────────────────

ifBlock
    : BLOCK_START KW_IF blockExpr BLOCK_END
      element*
      elifClause*
      elseClause?
      BLOCK_START KW_ENDIF BLOCK_END
    ;

elifClause : BLOCK_START KW_ELIF blockExpr BLOCK_END element* ;
elseClause : BLOCK_START KW_ELSE BLOCK_END element* ;

// Minimal boolean expression covering the common SQLMesh patterns:
//   {% if var %}, {% if var == 'val' %}, {% if var == other_var %}
blockExpr
    : B_NAME
    | B_NAME B_ASSIGN B_STRING
    | B_NAME B_ASSIGN B_NAME
    ;

// ── Control flow — {% for %} ──────────────────────────────────────────────────

forBlock
    : BLOCK_START KW_FOR B_NAME KW_IN B_NAME BLOCK_END
      element*
      BLOCK_START KW_ENDFOR BLOCK_END
    ;

// ── Macro definition — {% macro %} ───────────────────────────────────────────

macroBlock
    : BLOCK_START KW_MACRO B_NAME B_LPAREN paramList? B_RPAREN BLOCK_END
      element*
      BLOCK_START KW_ENDMACRO BLOCK_END
    ;

paramList : param (B_COMMA param)* ;

param : B_NAME (B_ASSIGN B_STRING)?   # namedParam ;

// ── Variable assignment — {% set %} ──────────────────────────────────────────

setStmt : BLOCK_START KW_SET B_NAME B_ASSIGN (B_NAME | B_STRING) BLOCK_END ;
