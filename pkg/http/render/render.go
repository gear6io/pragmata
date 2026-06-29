package render

import (
	"encoding/json"
	"net/http"

	"github.com/gear6io/pragmata/pkg/errors"
)

// SuccessResponse is the standard JSON envelope for successful responses.
type SuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

// ErrorResponse is the shape written by ErrorFrom. Used for OpenAPI schema registration.
type ErrorResponse struct {
	Status      string                       `json:"status"`
	Type        string                       `json:"type"`
	Code        string                       `json:"code"`
	Message     string                       `json:"message"`
	Url         string                       `json:"url,omitempty"`
	Errors      []errors.ErrorAdditional  `json:"errors"`
	Retry       *errors.RetryJSON         `json:"retry,omitempty"`
	Suggestions []string                     `json:"suggestions"`
}

// Success writes a 2xx JSON response wrapped in SuccessResponse.
func Success[T any](w http.ResponseWriter, status int, data T) {
	writeJSON(w, status, SuccessResponse{Status: "success", Data: data})
}

// ErrorFrom writes a structured error JSON response derived from err.
// The HTTP status code and response body are both driven by the error's type.
func ErrorFrom(w http.ResponseWriter, err error) {
	writeJSON(w, errors.HTTPStatus(err), struct {
		Status string `json:"status"`
		*errors.JSON
	}{"error", errors.AsJSON(err)})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
