#!/usr/bin/env bash
# Regenerate pkg/grammars/filtergrammar/*.go from the split grammars:
#   FilterQueryLexer.g4  — lexer grammar
#   FilterQuery.g4       — parser grammar (tokenVocab=FilterQueryLexer)
#
# Requires ANTLR 4.13.x and Java.
# Quick install (macOS):
#   brew install antlr
# Or via JAR:
#   curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar
#   alias antlr='java -jar "$(pwd)/antlr-4.13.1-complete.jar"'
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GRAMMAR_DIR="$REPO_ROOT/pkg/grammars/filtergrammar"

echo "Generating FilterQuery Go parser in $GRAMMAR_DIR ..."

rm -f "$GRAMMAR_DIR"/filterlang_lexer.go \
      "$GRAMMAR_DIR"/filterlang_parser.go \
      "$GRAMMAR_DIR"/filterlang_visitor.go \
      "$GRAMMAR_DIR"/filterlang_base_visitor.go \
      "$GRAMMAR_DIR"/filterlang_listener.go \
      "$GRAMMAR_DIR"/filterlang_base_listener.go \
      "$GRAMMAR_DIR"/doc.go

(
  cd "$GRAMMAR_DIR"

  antlr \
    -Dlanguage=Go \
    -package filtergrammar \
    FilterQueryLexer.g4

  antlr \
    -visitor \
    -Dlanguage=Go \
    -package filtergrammar \
    FilterQuery.g4
)

echo "Done. Generated files in $GRAMMAR_DIR"
