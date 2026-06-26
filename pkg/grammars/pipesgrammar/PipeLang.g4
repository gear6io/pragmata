parser grammar PipeLang;

options { tokenVocab=PipeLangLexer; }

// ── Top-level ──────────────────────────────────────────────────────────────────

pipeFile : directive* EOF ;

// ── Directives ─────────────────────────────────────────────────────────────────

directive
    : TYPE          VALUE            # type
    | NAME          VALUE            # name
    | DESCRIPTION   VALUE            # description
    | DESCRIPTION_ML SECTION_LINE*   # descriptionML
    | TAGS          VALUE            # tags
    | OWNER         VALUE            # owner
    | DESTINATION   VALUE            # destination
    | SCHEDULE      VALUE            # schedule
    | UNIQUE_KEY    VALUE            # uniqueKey
    | SOURCES       source*          # sourcesClause
    | PARAMS        param*           # paramsClause
    | PIPELINE      SECTION_LINE*    # pipelineClause
    ;

// ── Sources block ──────────────────────────────────────────────────────────────
// DASH is a real token (not skipped) because sources content lives in DEFAULT mode.

source
    : DASH alias=IDENTIFIER COLON_TOK table=IDENTIFIER  # aliasedSource
    | DASH name=IDENTIFIER                               # simpleSource
    ;

// ── Params block ───────────────────────────────────────────────────────────────
// 'type' and 'default' keys match IDENTIFIER at non-col-0; no KW_TYPE / KW_DEFAULT needed.

param
    : pname=IDENTIFIER COLON_TOK LBRACE
      IDENTIFIER COLON_TOK dtype=IDENTIFIER
      (COMMA IDENTIFIER COLON_TOK paramValue)?
      RBRACE
    ;

paramValue : STRING_LIT | NUMBER | IDENTIFIER ;
