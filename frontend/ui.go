// Package frontend embeds the compiled SPA into the Go binary.
package frontend

import "embed"

//go:embed all:dist
var FS embed.FS
