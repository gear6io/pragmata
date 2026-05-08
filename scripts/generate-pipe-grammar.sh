#!/usr/bin/env bash
# Regenerate pkg/parser/pipeparser/grammar/*.go from the split grammars:
#   PipeLangLexer.g4  — lexer grammar (handles SQL_BODY_MODE)
#   PipeLang.g4       — parser grammar (tokenVocab=PipeLangLexer)
#
# Requires ANTLR 4.13.x and Java.
# Quick install (macOS):
#   brew install antlr
# Or via JAR:
#   curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar
#   alias antlr='java -jar "$(pwd)/antlr-4.13.1-complete.jar"'
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GRAMMAR_DIR="$REPO_ROOT/pkg/grammars/pipesgrammar"

echo "Generating PipeLang Go parser in $GRAMMAR_DIR ..."

# Remove hand-written stubs; ANTLR will generate the real implementations.
rm -f "$GRAMMAR_DIR"/pipelang_lexer.go \
      "$GRAMMAR_DIR"/pipelang_parser.go \
      "$GRAMMAR_DIR"/pipelang_visitor.go \
      "$GRAMMAR_DIR"/pipelang_base_visitor.go \
      "$GRAMMAR_DIR"/pipelang_listener.go \
      "$GRAMMAR_DIR"/pipelang_base_listener.go \
      "$GRAMMAR_DIR"/doc.go

# Run ANTLR from inside the grammar directory so that:
#   1. PipeLangLexer.tokens is written next to the grammar files
#   2. The parser grammar finds PipeLangLexer.tokens via tokenVocab
# The lexer must be generated before the parser.
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
