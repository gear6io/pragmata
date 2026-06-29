package errors

import "net/http"

// HTTPStatus maps an error's type to an HTTP status code.
func HTTPStatus(err error) int {
	t, _, _, _, _, _ := Unwrapb(err)
	switch t {
	case TypeNotFound:
		return http.StatusNotFound
	case TypeInvalidInput, TypeUnsupported:
		return http.StatusBadRequest
	case TypeAlreadyExists:
		return http.StatusConflict
	case TypeUnauthenticated:
		return http.StatusUnauthorized
	case TypeForbidden:
		return http.StatusForbidden
	case TypeMethodNotAllowed:
		return http.StatusMethodNotAllowed
	case TypeTimeout:
		return http.StatusRequestTimeout
	case TypeTooManyRequests:
		return http.StatusTooManyRequests
	case TypeCanceled:
		return 499
	default:
		return http.StatusInternalServerError
	}
}

// JSON is the structured error response body. Arrays are never nil (OpenAPI-safe).
type JSON struct {
	Type        string           `json:"type"`
	Code        string           `json:"code"`
	Message     string           `json:"message"`
	Url         string           `json:"url,omitempty"`
	Errors      []ErrorAdditional `json:"errors"`
	Retry       *RetryJSON       `json:"retry,omitempty"`
	Suggestions []string         `json:"suggestions"`
}

// ErrorAdditional is a supplementary error detail in the JSON response.
type ErrorAdditional struct {
	Message     string   `json:"message"`
	Suggestions []string `json:"suggestions"`
}

// RetryJSON carries retry-after metadata in the JSON response.
type RetryJSON struct {
	Delay int64 `json:"delay_ms"`
}

// AsJSON converts any error to a JSON response struct.
// Plain (non-base) errors are mapped to TypeInternal / CodeUnknown.
func AsJSON(cause error) *JSON {
	t, c, m, _, u, a := Unwrapb(cause)

	ea := make([]ErrorAdditional, len(a))
	for i, v := range a {
		ea[i] = ErrorAdditional{Message: v.message, Suggestions: nonNilStrings(v.suggestions)}
	}

	var retryJSON *RetryJSON
	if r := retryOf(cause); r != nil {
		retryJSON = &RetryJSON{Delay: r.delay.Milliseconds()}
	}

	return &JSON{
		Type:        t.String(),
		Code:        c.String(),
		Message:     m,
		Url:         u,
		Errors:      ea,
		Retry:       retryJSON,
		Suggestions: nonNilStrings(suggestionsOf(cause)),
	}
}

func nonNilStrings(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}
