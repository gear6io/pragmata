// Package parser reads .pipe files into pipetypes.Pipe structs.
//
// .pipe file grammar (simplified):
//
//	DESCRIPTION <text>
//	TAGS <tag1>, <tag2>
//
//	NODE <name>
//	SQL >
//	  %
//	  <sql body>
//
//	TYPE ENDPOINT | MATERIALIZED | COPY
//	DATASOURCE <name>          (MATERIALIZED only)
//	TARGET_DATASOURCE <name>   (COPY only)
//	COPY_SCHEDULE <cron>       (COPY only)
//
// The % on the first line of a SQL block marks it as template-enabled.
// Directives (TYPE, DATASOURCE, etc.) may appear before or after NODE blocks.
package parser

import (
	"fmt"
	"strings"

	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
)

// Parse converts raw .pipe file content into a Pipe. name is set as the pipe name.
func Parse(name, content string) (*pipetypes.Pipe, error) {
	p := &pipetypes.Pipe{Name: name}
	lines := strings.Split(content, "\n")

	type parserState int
	const (
		stateHeader parserState = iota
		stateNode
		stateSQL
	)

	state := stateHeader
	var currentNode string
	var sqlLines []string
	sqlStarted := false

	flush := func() {
		if currentNode == "" {
			return
		}
		raw := strings.Join(sqlLines, "\n")
		isTemplated := false

		// Check for leading % line (template marker)
		trimmedRaw := strings.TrimLeft(raw, "\n\r ")
		if strings.HasPrefix(trimmedRaw, "%") {
			afterPercent := strings.TrimPrefix(trimmedRaw, "%")
			// Strip the % line
			if strings.HasPrefix(afterPercent, "\n") {
				raw = afterPercent[1:]
			} else if afterPercent == "" {
				raw = ""
			} else {
				raw = afterPercent
			}
			isTemplated = true
		}

		p.Nodes = append(p.Nodes, pipetypes.Node{
			Name:        currentNode,
			SQL:         strings.TrimSpace(dedent(raw)),
			IsTemplated: isTemplated,
		})
		currentNode = ""
		sqlLines = nil
		sqlStarted = false
	}

	for lineIdx, rawLine := range lines {
		lineNum := lineIdx + 1
		trimmed := strings.TrimSpace(rawLine)

		switch state {
		case stateSQL:
			if isTopLevelDirective(trimmed) {
				flush()
				state = stateHeader
				if err := applyDirective(p, trimmed); err != nil {
					return nil, fmt.Errorf("line %d: %w", lineNum, err)
				}
			} else if strings.HasPrefix(trimmed, "NODE ") {
				flush()
				currentNode = strings.TrimSpace(strings.TrimPrefix(trimmed, "NODE "))
				if currentNode == "" {
					return nil, fmt.Errorf("line %d: NODE requires a name", lineNum)
				}
				state = stateNode
			} else {
				if sqlStarted || trimmed != "" {
					sqlStarted = true
					sqlLines = append(sqlLines, rawLine)
				}
			}

		case stateNode:
			if trimmed == "SQL >" || trimmed == "SQL>" {
				state = stateSQL
			} else if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
				return nil, fmt.Errorf("line %d: expected 'SQL >' after NODE declaration, got %q", lineNum, trimmed)
			}

		case stateHeader:
			if trimmed == "" || strings.HasPrefix(trimmed, "#") {
				continue
			}
			if strings.HasPrefix(trimmed, "NODE ") {
				currentNode = strings.TrimSpace(strings.TrimPrefix(trimmed, "NODE "))
				if currentNode == "" {
					return nil, fmt.Errorf("line %d: NODE requires a name", lineNum)
				}
				state = stateNode
			} else {
				if err := applyDirective(p, trimmed); err != nil {
					return nil, fmt.Errorf("line %d: %w", lineNum, err)
				}
			}
		}
	}

	if state == stateSQL {
		flush()
	}

	if len(p.Nodes) == 0 {
		return nil, fmt.Errorf("pipe %q has no NODE blocks", name)
	}
	if p.Type == pipetypes.PipeTypeUndefined {
		return nil, fmt.Errorf("pipe %q has no TYPE declaration", name)
	}
	if p.Type == pipetypes.PipeTypeMaterialized && p.Datasource == "" {
		return nil, fmt.Errorf("pipe %q: TYPE MATERIALIZED requires DATASOURCE", name)
	}
	if p.Type == pipetypes.PipeTypeCopy && p.TargetDatasource == "" {
		return nil, fmt.Errorf("pipe %q: TYPE COPY requires TARGET_DATASOURCE", name)
	}

	return p, nil
}

var topLevelDirectives = []string{
	"DESCRIPTION ", "TAGS ", "TYPE ", "DATASOURCE ",
	"TARGET_DATASOURCE ", "COPY_SCHEDULE ",
}

func isTopLevelDirective(line string) bool {
	for _, d := range topLevelDirectives {
		if strings.HasPrefix(line, d) {
			return true
		}
	}
	return false
}

func applyDirective(p *pipetypes.Pipe, line string) error {
	switch {
	case strings.HasPrefix(line, "DESCRIPTION "):
		p.Description = strings.TrimSpace(strings.TrimPrefix(line, "DESCRIPTION "))
	case strings.HasPrefix(line, "TAGS "):
		raw := strings.TrimPrefix(line, "TAGS ")
		for _, t := range strings.Split(raw, ",") {
			if tag := strings.TrimSpace(t); tag != "" {
				p.Tags = append(p.Tags, tag)
			}
		}
	case strings.HasPrefix(line, "TYPE "):
		p.Type = pipetypes.PipeType{valuer.NewString(strings.TrimSpace(strings.TrimPrefix(line, "TYPE ")))}
	case strings.HasPrefix(line, "DATASOURCE "):
		p.Datasource = strings.TrimSpace(strings.TrimPrefix(line, "DATASOURCE "))
	case strings.HasPrefix(line, "TARGET_DATASOURCE "):
		p.TargetDatasource = strings.TrimSpace(strings.TrimPrefix(line, "TARGET_DATASOURCE "))
	case strings.HasPrefix(line, "COPY_SCHEDULE "):
		p.CopySchedule = strings.TrimSpace(strings.TrimPrefix(line, "COPY_SCHEDULE "))
	}
	return nil
}

// dedent removes the common leading whitespace from all non-empty lines.
func dedent(s string) string {
	lines := strings.Split(s, "\n")
	minIndent := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		n := len(line) - len(strings.TrimLeft(line, " \t"))
		if minIndent < 0 || n < minIndent {
			minIndent = n
		}
	}
	if minIndent <= 0 {
		return s
	}
	out := make([]string, len(lines))
	for i, line := range lines {
		if len(line) >= minIndent {
			out[i] = line[minIndent:]
		}
	}
	return strings.Join(out, "\n")
}
