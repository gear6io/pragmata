#!/usr/bin/env bash
# Regenerate pkg/grammars/jinjagrammar/*.go from the split grammars:
#   JinjaLexer.g4  — lexer grammar (handles EXPR/BLOCK/COMMENT modes)
#   Jinja.g4       — parser grammar (tokenVocab=JinjaLexer)
#
# Requires ANTLR 4.13.x and Java.
# Quick install (macOS):
#   brew install antlr
# Or via JAR:
#   curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar
#   alias antlr='java -jar "$(pwd)/antlr-4.13.1-complete.jar"'
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GRAMMAR_DIR="$REPO_ROOT/pkg/grammars/jinjagrammar"

echo "Generating Jinja Go parser in $GRAMMAR_DIR ..."

# Remove previously generated files so stale code doesn't linger.
rm -f "$GRAMMAR_DIR"/jinja_lexer.go \
      "$GRAMMAR_DIR"/jinja_parser.go \
      "$GRAMMAR_DIR"/jinja_visitor.go \
      "$GRAMMAR_DIR"/jinja_base_visitor.go \
      "$GRAMMAR_DIR"/jinja_listener.go \
      "$GRAMMAR_DIR"/jinja_base_listener.go

# Run ANTLR from inside the grammar directory so that:
#   1. JinjaLexer.tokens is written next to the grammar files
#   2. The parser grammar finds JinjaLexer.tokens via tokenVocab
# The lexer must be generated before the parser.
(
  cd "$GRAMMAR_DIR"

  antlr \
    -Dlanguage=Go \
    -package jinjagrammar \
    JinjaLexer.g4

  antlr \
    -visitor \
    -Dlanguage=Go \
    -package jinjagrammar \
    Jinja.g4
)

echo "Done. Generated files in $GRAMMAR_DIR"
