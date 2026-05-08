// Package pipeparser parses .pipe file content using an ANTLR4-generated grammar.
// Run scripts/generate-pipe-grammar.sh to regenerate grammar/ after editing PipeLang.g4.
package pipeparser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gear6io/pragmata/pkg/parser/pipeparser/grammar"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// Parse converts raw .pipe file content into a Pipe. name is set as Pipe.Name
// (caller supplies it, typically derived from the filename).
func Parse(name, content string) (*pipetypes.Pipe, error) {
	// Ensure a trailing newline so the final SQL_CHAR accumulation is always
	// flushed by SQL_NL before EOF.
	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}

	lexer := grammar.NewPipeLangLexer(antlr.NewInputStream(content))
	lexErr := &errListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErr)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := grammar.NewPipeLang(stream)
	parseErr := &errListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErr)

	tree := p.PipeFile()

	if lexErr.msg != "" {
		return nil, fmt.Errorf("lex error in %q: %s", name, lexErr.msg)
	}
	if parseErr.msg != "" {
		return nil, fmt.Errorf("parse error in %q: %s", name, parseErr.msg)
	}

	c := &converter{pipe: &pipetypes.Pipe{Name: name}}
	c.Visit(tree)
	return c.result()
}

// errListener collects the first syntax error produced by the lexer or parser.
type errListener struct {
	*antlr.DefaultErrorListener
	msg string
}

func (e *errListener) SyntaxError(
	_ antlr.Recognizer,
	_ interface{},
	line, col int,
	msg string,
	_ antlr.RecognitionException,
) {
	if e.msg == "" {
		e.msg = fmt.Sprintf("line %d:%d %s", line, col, msg)
	}
}
