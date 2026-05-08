#!/usr/bin/env bash
# Regenerate pkg/parser/sqlmesh/grammar/*.go from the split grammars:
#   SQLMeshLexer.g4  — lexer grammar (handles SQL_BODY_MODE)
#   SQLMesh.g4       — parser grammar (tokenVocab=SQLMeshLexer)
#
# Requires ANTLR 4.13.x and Java.
# Quick install (macOS):
#   brew install antlr
# Or via JAR:
#   curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar
#   alias antlr='java -jar "$(pwd)/antlr-4.13.1-complete.jar"'
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GRAMMAR_DIR="$REPO_ROOT/pkg/grammars/sqlmeshgrammar"

echo "Generating SQLMesh Go parser in $GRAMMAR_DIR ..."

# Remove hand-written stubs; ANTLR will generate the real implementations.
rm -f "$GRAMMAR_DIR"/sqlmesh_lexer.go \
      "$GRAMMAR_DIR"/sqlmesh_parser.go \
      "$GRAMMAR_DIR"/sqlmesh_visitor.go \
      "$GRAMMAR_DIR"/sqlmesh_base_visitor.go \
      "$GRAMMAR_DIR"/sqlmesh_listener.go \
      "$GRAMMAR_DIR"/sqlmesh_base_listener.go

# Run ANTLR from inside the grammar directory so that:
#   1. SQLMeshLexer.tokens is written next to the grammar files
#   2. The parser grammar finds SQLMeshLexer.tokens via tokenVocab
# The lexer must be generated before the parser.
(
  cd "$GRAMMAR_DIR"

  antlr \
    -Dlanguage=Go \
    -package grammar \
    SQLMeshLexer.g4

  antlr \
    -visitor \
    -Dlanguage=Go \
    -package grammar \
    SQLMesh.g4
)

echo "Done. Generated files in $GRAMMAR_DIR"
