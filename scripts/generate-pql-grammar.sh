#!/usr/bin/env bash
# Regenerate pkg/grammars/pqlgrammar/*.go from the split grammars:
#   PQLLexer.g4  — lexer grammar (handles FILTER_BODY_MODE)
#   PQL.g4       — parser grammar (tokenVocab=PQLLexer)
#
# Requires ANTLR 4.13.x and Java.
# Quick install (macOS):
#   brew install antlr
# Or via JAR:
#   curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar
#   alias antlr='java -jar "$(pwd)/antlr-4.13.1-complete.jar"'
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GRAMMAR_DIR="$REPO_ROOT/pkg/grammars/pqlgrammar"

echo "Generating PQL Go parser in $GRAMMAR_DIR ..."

rm -f "$GRAMMAR_DIR"/pql_lexer.go \
      "$GRAMMAR_DIR"/pql_parser.go \
      "$GRAMMAR_DIR"/pql_visitor.go \
      "$GRAMMAR_DIR"/pql_base_visitor.go \
      "$GRAMMAR_DIR"/pql_listener.go \
      "$GRAMMAR_DIR"/pql_base_listener.go \
      "$GRAMMAR_DIR"/doc.go

(
  cd "$GRAMMAR_DIR"

  antlr \
    -Dlanguage=Go \
    -package pqlgrammar \
    PQLLexer.g4

  antlr \
    -visitor \
    -Dlanguage=Go \
    -package pqlgrammar \
    PQL.g4
)

echo "Done. Generated files in $GRAMMAR_DIR"
