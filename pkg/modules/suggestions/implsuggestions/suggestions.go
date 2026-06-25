package implsuggestions

import (
	"context"
	"regexp"
	"strings"

	"github.com/gear6io/pragmata/pkg/modules/suggestions"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/types/suggestiontypes"
)

type module struct {
	store sqlstore.SQLStore
}

// NewModule returns a Module backed by the given store.
func NewModule(store sqlstore.SQLStore) suggestions.Module {
	return &module{store: store}
}

func (m *module) GetSuggestions(ctx context.Context, req suggestiontypes.SuggestionRequest) (*suggestiontypes.SuggestionResponse, error) {
	switch req.ContextType {
	case suggestiontypes.ContextTypeSource:
		return m.sourceSuggestions(ctx, req)
	case suggestiontypes.ContextTypeField:
		return m.fieldSuggestions(ctx, req)
	default:
		return &suggestiontypes.SuggestionResponse{Complete: true, Suggestions: []suggestiontypes.Suggestion{}}, nil
	}
}

// sourceSuggestions returns pipe names as source candidates.
// Raw database table introspection is a stub — will be filled once the
// datasource/schema layer is implemented.
func (m *module) sourceSuggestions(ctx context.Context, req suggestiontypes.SuggestionRequest) (*suggestiontypes.SuggestionResponse, error) {
	pipes, err := m.store.ListPipes(ctx)
	if err != nil {
		return nil, err
	}

	var out []suggestiontypes.Suggestion
	for _, p := range pipes {
		if !matchesText(p.Name, req.SearchText, req.MatchingType) {
			continue
		}
		out = append(out, suggestiontypes.Suggestion{
			Value:  p.Name,
			Label:  p.Name,
			Kind:   suggestiontypes.ContextTypeSource,
			Detail: "pipe",
		})
	}

	return &suggestiontypes.SuggestionResponse{Complete: true, Suggestions: coalesce(out)}, nil
}

// fieldSuggestions resolves columns for NodeRef in this priority order:
//  1. A node defined in the current pipe content (same-pipe node)
//  2. Another pipe whose name matches NodeRef (cross-pipe: last-node columns)
//  3. A raw source table — stub, returns empty until datasource introspection lands
func (m *module) fieldSuggestions(ctx context.Context, req suggestiontypes.SuggestionRequest) (*suggestiontypes.SuggestionResponse, error) {
	nodeRef := strings.TrimPrefix(req.NodeRef, "@")

	// 1. Same-pipe node resolution.
	if req.PipeContent != "" {
		pipe, err := pipevisitor.Visit("", req.PipeContent)
		if err == nil {
			for _, node := range pipe.Nodes {
				if strings.EqualFold(node.Name, nodeRef) {
					cols := extractSelectColumns(node.SQL)
					return buildFieldResponse(cols, "node "+node.Name, req.SearchText, req.MatchingType), nil
				}
			}
		}
	}

	// 2. Cross-pipe resolution: find a stored pipe whose name matches NodeRef.
	if nodeRef != "" {
		stored, err := m.store.GetPipe(ctx, nodeRef)
		if err == nil && stored != nil {
			cols := lastNodeColumns(stored)
			return buildFieldResponse(cols, "pipe "+stored.Name, req.SearchText, req.MatchingType), nil
		}
	}

	// 3. Raw source table — no introspection yet.
	return &suggestiontypes.SuggestionResponse{Complete: true, Suggestions: []suggestiontypes.Suggestion{}}, nil
}

// buildFieldResponse filters cols by searchText/matchingType and wraps them as Suggestions.
func buildFieldResponse(cols []string, detail, searchText string, mt suggestiontypes.MatchingType) *suggestiontypes.SuggestionResponse {
	var out []suggestiontypes.Suggestion
	for _, col := range cols {
		if !matchesText(col, searchText, mt) {
			continue
		}
		out = append(out, suggestiontypes.Suggestion{
			Value:  col,
			Label:  col,
			Kind:   suggestiontypes.ContextTypeField,
			Detail: detail,
		})
	}
	return &suggestiontypes.SuggestionResponse{Complete: true, Suggestions: coalesce(out)}
}

// lastNodeColumns returns the SELECT column aliases of the last node in a pipe.
func lastNodeColumns(pipe *pipetypes.Pipe) []string {
	if len(pipe.Nodes) == 0 {
		return nil
	}
	return extractSelectColumns(pipe.Nodes[len(pipe.Nodes)-1].SQL)
}

// selectListRE captures everything between SELECT and the first FROM/WHERE/GROUP/HAVING/ORDER/LIMIT.
var selectListRE = regexp.MustCompile(`(?is)SELECT\s+(DISTINCT\s+)?(.+?)\s+FROM\b`)

// columnTokenRE splits a SELECT list on commas that are not inside parentheses.
// We walk the string manually to handle nested function calls.
var aliasRE = regexp.MustCompile(`(?i)\bAS\s+(\w+)\s*$`)
var bareColRE = regexp.MustCompile(`(?:\w+\.)?(\w+)\s*$`)

// extractSelectColumns parses a SQL string and returns the output column names/aliases.
// It handles the common cases: bare column, table.column, expr AS alias, *.
// Wildcards (*) are omitted since we can't enumerate them statically.
func extractSelectColumns(sql string) []string {
	m := selectListRE.FindStringSubmatch(sql)
	if m == nil {
		return nil
	}
	list := m[2]
	items := splitSelectList(list)

	var cols []string
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" || item == "*" {
			continue
		}
		// AS alias takes precedence.
		if a := aliasRE.FindStringSubmatch(item); a != nil {
			cols = append(cols, a[1])
			continue
		}
		// table.column or plain column — strip table prefix.
		if b := bareColRE.FindStringSubmatch(item); b != nil {
			cols = append(cols, b[1])
		}
	}
	return cols
}

// splitSelectList splits a SELECT list on top-level commas (not inside parentheses).
func splitSelectList(s string) []string {
	var parts []string
	depth := 0
	start := 0
	for i, ch := range s {
		switch ch {
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, s[start:i])
				start = i + 1
			}
		}
	}
	parts = append(parts, s[start:])
	return parts
}

// matchesText returns true when candidate satisfies the search text under the given matching type.
func matchesText(candidate, searchText string, mt suggestiontypes.MatchingType) bool {
	if searchText == "" {
		return true
	}
	switch mt {
	case suggestiontypes.MatchingTypeExact:
		return strings.EqualFold(candidate, searchText)
	default: // fuzzy = case-insensitive substring
		return strings.Contains(strings.ToLower(candidate), strings.ToLower(searchText))
	}
}

// coalesce returns an empty slice instead of nil so JSON marshals as [].
func coalesce(s []suggestiontypes.Suggestion) []suggestiontypes.Suggestion {
	if s == nil {
		return []suggestiontypes.Suggestion{}
	}
	return s
}
