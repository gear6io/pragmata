// Package template renders {{ Type(name, default) }} placeholders in SQL strings.
//
// Supported types: String, Int32, Int64, Float32, Float64, DateTime.
//
// Each placeholder is replaced by the resolved parameter value. For String,
// the value is single-quoted and SQL-escaped. Numeric types are validated then
// inserted verbatim so ClickHouse evaluates in-expression arithmetic:
//
//	OFFSET {{ Int32(page, 0) * Int32(page_size, 100) }}
//
// with page=2, page_size=10 becomes:
//
//	OFFSET 2 * 10
package template

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gear6io/pragmata/pkg/errors"
)

var (
	// blockRe matches {{ ... }} blocks, non-greedy.
	blockRe = regexp.MustCompile(`\{\{([^}]*)\}\}`)
	// callRe matches a single Type(name, default) call inside an expression.
	callRe = regexp.MustCompile(`(\w+)\(([^)]+)\)`)

	CodeInvalidParam = errors.MustNewCode("invalid_param")
)

// Render substitutes all {{ Type(name, default) }} tokens in sql using params.
// URL query params take precedence over the declared defaults.
func Render(sql string, params map[string]string) (string, error) {
	var renderErr error
	result := blockRe.ReplaceAllStringFunc(sql, func(block string) string {
		if renderErr != nil {
			return block
		}
		inner := block[2 : len(block)-2] // strip {{ and }}
		resolved, err := resolveBlock(inner, params)
		if err != nil {
			renderErr = errors.WithAdditionalf(err, "block %q", block)
			return block
		}
		return resolved
	})
	return result, renderErr
}

// resolveBlock replaces every Type(name, default) call inside a {{ }} block.
func resolveBlock(expr string, params map[string]string) (string, error) {
	var resolveErr error
	result := callRe.ReplaceAllStringFunc(expr, func(call string) string {
		if resolveErr != nil {
			return call
		}
		m := callRe.FindStringSubmatch(call)
		if m == nil {
			return call
		}
		typeName := m[1]
		args := splitArgs(m[2])
		if len(args) == 0 {
			resolveErr = errors.NewInvalidInputf(errors.CodeInvalidInput, "%q: missing name argument", call)
			return call
		}
		paramName := strings.TrimSpace(args[0])
		var defaultVal string
		if len(args) > 1 {
			defaultVal = strings.TrimSpace(args[1])
		}
		val, provided := params[paramName]
		if !provided {
			val = defaultVal
		}
		rendered, err := formatTyped(typeName, val, defaultVal, provided)
		if err != nil {
			resolveErr = errors.WithAdditionalf(err, "param %q", paramName)
			return call
		}
		return rendered
	})
	return result, resolveErr
}

// formatTyped validates and formats a parameter value for its SQL type.
func formatTyped(typeName, val, defaultVal string, provided bool) (string, error) {
	switch typeName {
	case "String":
		s := val
		if !provided || s == "" {
			s = strings.Trim(defaultVal, "'\"")
		}
		return "'" + escapeStr(s) + "'", nil

	case "Int32", "Int64":
		if val == "" {
			return defaultVal, nil
		}
		if _, err := strconv.ParseInt(val, 10, 64); err != nil {
			return "", errors.NewInvalidInputf(CodeInvalidParam, "invalid %s %q", typeName, val)
		}
		return val, nil

	case "Float32", "Float64":
		if val == "" {
			return defaultVal, nil
		}
		if _, err := strconv.ParseFloat(val, 64); err != nil {
			return "", errors.NewInvalidInputf(CodeInvalidParam, "invalid %s %q", typeName, val)
		}
		return val, nil

	case "DateTime":
		if val == "" {
			return defaultVal, nil
		}
		return "'" + escapeStr(val) + "'", nil

	default:
		if val == "" {
			return defaultVal, nil
		}
		return val, nil
	}
}

// splitArgs splits a comma-separated argument list, respecting single-quoted strings.
func splitArgs(s string) []string {
	var args []string
	var cur strings.Builder
	inQuote := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '\'' && !inQuote:
			inQuote = true
			cur.WriteByte(c)
		case c == '\'' && inQuote:
			inQuote = false
			cur.WriteByte(c)
		case c == ',' && !inQuote:
			args = append(args, cur.String())
			cur.Reset()
		default:
			cur.WriteByte(c)
		}
	}
	if cur.Len() > 0 {
		args = append(args, cur.String())
	}
	return args
}

func escapeStr(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
