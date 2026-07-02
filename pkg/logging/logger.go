package logging

import (
	"context"
	"log/slog"
	"os"
	"runtime"

	"go.opentelemetry.io/otel/trace"

	"github.com/gear6io/pragmata/pkg/errors"
)

// NewLogger builds a structured JSON logger writing to stderr and sets it as
// the slog default. The handler chain enriches every record with:
//   - code.filepath / code.function / code.lineno (call site)
//   - trace_id / span_id (from OTel context when available)
//   - exception.* fields (when errors.Attr(err) is present)
func NewLogger(level slog.Level) *slog.Logger {
	base := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
			}
			return a
		},
	})
	logger := slog.New(&handler{base: base})
	slog.SetDefault(logger)
	return logger
}

type handler struct{ base slog.Handler }

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.base.Enabled(ctx, level)
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handler{base: h.base.WithAttrs(attrs)}
}

func (h *handler) WithGroup(name string) slog.Handler {
	return &handler{base: h.base.WithGroup(name)}
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	// source
	if r.PC != 0 {
		frame, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()
		r.AddAttrs(
			slog.String("code.filepath", frame.File),
			slog.String("code.function", frame.Function),
			slog.Int("code.lineno", frame.Line),
		)
	}

	// correlation
	if span := trace.SpanFromContext(ctx); span != nil && span.IsRecording() {
		sc := span.SpanContext()
		if sc.HasTraceID() {
			r.AddAttrs(slog.String("trace_id", sc.TraceID().String()))
		}
		if sc.HasSpanID() {
			r.AddAttrs(slog.String("span_id", sc.SpanID().String()))
		}
	}

	// exception expansion: find errors.Attr(err) and replace with structured fields
	var foundErr error
	newRecord := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	r.Attrs(func(a slog.Attr) bool {
		if a.Key == "exception" {
			if err, ok := a.Value.Any().(error); ok {
				foundErr = err
				return true
			}
		}
		newRecord.AddAttrs(a)
		return true
	})
	if foundErr == nil {
		return h.base.Handle(ctx, r)
	}

	t, c, m, _, _, _ := errors.Unwrapb(foundErr)
	newRecord.AddAttrs(
		slog.String("exception.type", t.String()),
		slog.String("exception.code", c.String()),
		slog.String("exception.message", m),
	)
	type stacktracer interface{ Stacktrace() string }
	if st, ok := foundErr.(stacktracer); ok && st.Stacktrace() != "" {
		newRecord.AddAttrs(slog.String("exception.stacktrace", st.Stacktrace()))
	}
	return h.base.Handle(ctx, newRecord)
}
