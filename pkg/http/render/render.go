// Package render provides consistent JSON response helpers.
// Port of SigNoz's pkg/http/render pattern.
package render

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse is the standard JSON envelope for successful responses.
type SuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

// ErrorResponse is the standard JSON envelope for error responses.
type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Success writes a 2xx JSON response wrapped in SuccessResponse.
func Success[T any](w http.ResponseWriter, status int, data T) {
	writeJSON(w, status, SuccessResponse{Status: "success", Data: data})
}

// Error writes an error JSON response wrapped in ErrorResponse.
func Error(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, ErrorResponse{Status: "error", Error: msg})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
