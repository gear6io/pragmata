package suggestiontypes

// MatchingType controls how candidate strings are compared against the search text.
type MatchingType string

const (
	MatchingTypeExact MatchingType = "exact"
	MatchingTypeFuzzy MatchingType = "fuzzy"
)

// ContextType identifies which kind of completion is being requested.
type ContextType string

const (
	// ContextTypeSource requests suggestions for items in a sources: directive
	// (pipe names or database tables).
	ContextTypeSource ContextType = "source"
	// ContextTypeField requests column/field name suggestions for use inside SQL
	// pipeline node bodies.
	ContextTypeField ContextType = "field"
)

// Suggestion is a single autocomplete candidate.
type Suggestion struct {
	Value  string `json:"value"`
	Label  string `json:"label"`
	Kind   string `json:"kind"`             // "source" | "field"
	Detail string `json:"detail,omitempty"` // human-readable provenance, e.g. "from step1"
}

// SuggestionRequest is the payload sent by the editor for every completion trigger.
type SuggestionRequest struct {
	ContextType  ContextType  `json:"contextType"`
	MatchingType MatchingType `json:"matchingType"`
	SearchText   string       `json:"searchText"`
	// NodeRef is the alias or node name whose output columns are requested
	// (FieldContext only).
	NodeRef string `json:"nodeRef,omitempty"`
	// PipeContent is the full current pipe DSL text so the server can parse
	// same-pipe nodes without a round-trip (FieldContext only).
	PipeContent string `json:"pipeContent,omitempty"`
}

// SuggestionResponse is returned by the suggestions endpoint.
type SuggestionResponse struct {
	// Complete is false when results were truncated (future pagination hook).
	Complete    bool         `json:"complete"`
	Suggestions []Suggestion `json:"suggestions"`
}
