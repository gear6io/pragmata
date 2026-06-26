#!/usr/bin/env bash
# Regenerate pkg/grammars/pipesgrammar/*.go from:
#   PipeLangLexer.g4  — lexer grammar
#   PipeLang.g4       — parser grammar (tokenVocab=PipeLangLexer)
#
# Requires ANTLR 4.13.x and Java.
# Quick install (macOS): brew install antlr
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GRAMMAR_DIR="$REPO_ROOT/pkg/grammars/pipesgrammar"

echo "Generating PipeLang Go parser in $GRAMMAR_DIR ..."

rm -f "$GRAMMAR_DIR"/pipelang_lexer.go \
      "$GRAMMAR_DIR"/pipelang_parser.go \
      "$GRAMMAR_DIR"/pipelang_visitor.go \
      "$GRAMMAR_DIR"/pipelang_base_visitor.go \
      "$GRAMMAR_DIR"/pipelang_listener.go \
      "$GRAMMAR_DIR"/pipelang_base_listener.go

# Run from grammar dir so PipeLangLexer.tokens is written next to the .g4 files
# and the parser grammar finds it via tokenVocab. Lexer must come first.
(
  cd "$GRAMMAR_DIR"

  antlr \
    -Dlanguage=Go \
    -package grammar \
    PipeLangLexer.g4

  antlr \
    -visitor \
    -Dlanguage=Go \
    -package grammar \
    PipeLang.g4
)

echo "Done. Generated files in $GRAMMAR_DIR"
