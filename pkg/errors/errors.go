package errors

import (
	stderrors "errors"
	"fmt"
	"time"
)

// base is the fundamental struct that implements the error interface.
type base struct {
	t           typ
	c           Code
	m           string
	e           error
	u           string
	a           []additional
	s           fmt.Stringer
	r           *retry
	suggestions []string
}

// additional is a single supplementary error detail with optional suggestions.
type additional struct {
	message     string
	suggestions []string
}

// Stacktrace returns the stacktrace captured at error creation time.
func (b *base) Stacktrace() string {
	if b.s == nil {
		return ""
	}
	return b.s.String()
}

// WithStacktrace replaces the auto-captured stacktrace with a pre-formatted string.
func (b *base) WithStacktrace(s string) *base {
	return &base{t: b.t, c: b.c, m: b.m, e: b.e, u: b.u, a: b.a, s: rawStacktrace(s), r: b.r, suggestions: b.suggestions}
}

func (b *base) Error() string {
	if b.e != nil {
		return b.e.Error()
	}
	return b.m
}

// Unwrap exposes the wrapped cause so stdlib errors.Is / errors.As can walk the chain.
func (b *base) Unwrap() error {
	return b.e
}

func New(t typ, code Code, message string) *base {
	return &base{t: t, c: code, m: message, a: []additional{}, s: newStackTrace()}
}

func Newf(t typ, code Code, format string, args ...any) *base {
	return &base{t: t, c: code, m: fmt.Sprintf(format, args...), s: newStackTrace()}
}

func Wrap(cause error, t typ, code Code, message string) *base {
	b := &base{t: t, c: code, m: message, e: cause, s: newStackTrace()}
	propagateHints(b, cause)
	return b
}

func Wrapf(cause error, t typ, code Code, format string, args ...any) *base {
	b := &base{t: t, c: code, m: fmt.Sprintf(format, args...), e: cause, s: newStackTrace()}
	propagateHints(b, cause)
	return b
}

// propagateHints copies user-facing hints from the inner error so they survive re-wrapping.
func propagateHints(b *base, cause error) {
	if inner, ok := cause.(*base); ok {
		b.r = inner.r
		b.a = inner.a
		b.suggestions = inner.suggestions
	}
}

func WithAdditionalf(cause error, format string, args ...any) *base {
	if b, ok := cause.(*base); ok {
		return b.WithAdditional(fmt.Sprintf(format, args...))
	}
	t, c, m, e, u, a := Unwrapb(cause)
	b := &base{t: t, c: c, m: m, e: e, u: u, a: a, s: newStackTrace(), r: retryOf(cause)}
	return b.WithAdditional(fmt.Sprintf(format, args...))
}

func WithSuggestiveAdditionalf(cause error, suggestions []string, format string, args ...any) *base {
	if b, ok := cause.(*base); ok {
		return b.WithSuggestiveAdditional(fmt.Sprintf(format, args...), suggestions...)
	}
	t, c, m, e, u, a := Unwrapb(cause)
	b := &base{t: t, c: c, m: m, e: e, u: u, a: a, s: newStackTrace(), r: retryOf(cause)}
	return b.WithSuggestiveAdditional(fmt.Sprintf(format, args...), suggestions...)
}

func (b *base) WithUrl(u string) *base {
	return &base{t: b.t, c: b.c, m: b.m, e: b.e, u: u, a: b.a, s: b.s, r: b.r, suggestions: b.suggestions}
}

func (b *base) WithAdditional(messages ...string) *base {
	extra := make([]additional, len(messages))
	for i, m := range messages {
		extra[i] = additional{message: m}
	}
	return b.WithAdditionals(extra...)
}

func (b *base) WithAdditionals(additionals ...additional) *base {
	nb := *b
	nb.a = append(append([]additional{}, b.a...), additionals...)
	return &nb
}

func (b *base) withRetry(r retry) *base {
	return &base{t: b.t, c: b.c, m: b.m, e: b.e, u: b.u, a: b.a, s: b.s, r: &r, suggestions: b.suggestions}
}

func (b *base) WithSuggestions(suggestions ...string) *base {
	return &base{t: b.t, c: b.c, m: b.m, e: b.e, u: b.u, a: b.a, s: b.s, r: b.r, suggestions: suggestions}
}

func (b *base) WithSuggestiveAdditional(message string, suggestions ...string) *base {
	return b.WithAdditionals(additional{message: message, suggestions: suggestions})
}

func (b *base) WithRetryAfter(delay time.Duration) *base {
	return b.withRetry(newRetryAfter(delay))
}

// Unwrapb extracts the individual fields from the first *base in the chain.
// Falls back to TypeInternal + CodeUnknown for plain errors.
//
//nolint:staticcheck
func Unwrapb(cause error) (typ, Code, string, error, string, []additional) {
	if b, ok := cause.(*base); ok {
		return b.t, b.c, b.m, b.e, b.u, b.a
	}
	return TypeInternal, CodeUnknown, cause.Error(), cause, "", []additional{}
}

// Ast checks if the error matches the specified type.
func Ast(cause error, t typ) bool {
	got, _, _, _, _, _ := Unwrapb(cause)
	return got == t
}

// Asc checks if the error matches the specified code.
func Asc(cause error, code Code) bool {
	_, got, _, _, _, _ := Unwrapb(cause)
	return got.s == code.s
}

func Join(errs ...error) error          { return stderrors.Join(errs...) }
func As(err error, target any) bool     { return stderrors.As(err, target) }
func Is(err, target error) bool         { return stderrors.Is(err, target) }

// RetryDelayOf returns the explicit retry delay or zero if none is set.
func RetryDelayOf(err error) time.Duration {
	b, ok := err.(*base)
	if !ok || b.r == nil {
		return 0
	}
	return b.r.delay
}

func retryOf(err error) *retry {
	if b, ok := err.(*base); ok {
		return b.r
	}
	return nil
}

func suggestionsOf(err error) []string {
	if b, ok := err.(*base); ok {
		return b.suggestions
	}
	return nil
}

// Type-specific constructors (New + Wrap pairs).

func NewInvalidInputf(code Code, format string, args ...any) *base {
	return Newf(TypeInvalidInput, code, format, args...)
}
func WrapInvalidInputf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeInvalidInput, code, format, args...)
}

func NewInternalf(code Code, format string, args ...any) *base {
	return Newf(TypeInternal, code, format, args...)
}
func WrapInternalf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeInternal, code, format, args...)
}

func NewNotFoundf(code Code, format string, args ...any) *base {
	return Newf(TypeNotFound, code, format, args...)
}
func WrapNotFoundf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeNotFound, code, format, args...)
}

func NewAlreadyExistsf(code Code, format string, args ...any) *base {
	return Newf(TypeAlreadyExists, code, format, args...)
}
func WrapAlreadyExistsf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeAlreadyExists, code, format, args...)
}

func NewUnauthenticatedf(code Code, format string, args ...any) *base {
	return Newf(TypeUnauthenticated, code, format, args...)
}
func WrapUnauthenticatedf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeUnauthenticated, code, format, args...)
}

func NewForbiddenf(code Code, format string, args ...any) *base {
	return Newf(TypeForbidden, code, format, args...)
}
func WrapForbiddenf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeForbidden, code, format, args...)
}

func NewUnsupportedf(code Code, format string, args ...any) *base {
	return Newf(TypeUnsupported, code, format, args...)
}
func WrapUnsupportedf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeUnsupported, code, format, args...)
}

func NewTimeoutf(code Code, format string, args ...any) *base {
	return Newf(TypeTimeout, code, format, args...)
}
func WrapTimeoutf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeTimeout, code, format, args...)
}

func NewCanceledf(code Code, format string, args ...any) *base {
	return Newf(TypeCanceled, code, format, args...)
}
func WrapCanceledf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeCanceled, code, format, args...)
}

func NewMethodNotAllowedf(code Code, format string, args ...any) *base {
	return Newf(TypeMethodNotAllowed, code, format, args...)
}

func NewTooManyRequestsf(code Code, format string, args ...any) *base {
	return Newf(TypeTooManyRequests, code, format, args...)
}
func WrapTooManyRequestsf(cause error, code Code, format string, args ...any) *base {
	return Wrapf(cause, TypeTooManyRequests, code, format, args...)
}
