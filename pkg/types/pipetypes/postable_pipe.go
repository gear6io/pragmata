package pipetypes

// PostablePipe is the create/update request body for pipe endpoints.
// The raw PipeLang content is parsed server-side to populate all pipe fields.
type PostablePipe struct {
	Content string `json:"content"`
}
