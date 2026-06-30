package errors

import (
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
)

// Attr returns an slog.Attr with a standardized "exception" key for structured logging.
func Attr(err error) slog.Attr {
	return slog.Any("exception", err)
}

// TypeAttr returns an OTel attribute.KeyValue with the "error.type" semconv key.
func TypeAttr(err error) attribute.KeyValue {
	t, _, _, _, _, _ := Unwrapb(err)
	return attribute.String("error.type", t.String())
}
