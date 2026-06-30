package errors_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gear6io/pragmata/pkg/errors"
)

func TestHTTPStatus(t *testing.T) {
	cases := []struct {
		err  error
		want int
	}{
		{errors.NewNotFoundf(errors.CodeNotFound, "x"), http.StatusNotFound},
		{errors.NewInvalidInputf(errors.CodeInvalidInput, "x"), http.StatusBadRequest},
		{errors.NewAlreadyExistsf(errors.CodeAlreadyExists, "x"), http.StatusConflict},
		{errors.NewUnauthenticatedf(errors.CodeUnauthenticated, "x"), http.StatusUnauthorized},
		{errors.NewForbiddenf(errors.CodeForbidden, "x"), http.StatusForbidden},
		{errors.NewInternalf(errors.CodeInternal, "x"), http.StatusInternalServerError},
		{errors.NewTimeoutf(errors.CodeTimeout, "x"), http.StatusRequestTimeout},
		{errors.NewTooManyRequestsf(errors.CodeTooManyRequests, "x"), http.StatusTooManyRequests},
		{errors.NewCanceledf(errors.CodeCanceled, "x"), 499},
	}
	for _, c := range cases {
		if got := errors.HTTPStatus(c.err); got != c.want {
			t.Errorf("HTTPStatus(%v) = %d, want %d", c.err, got, c.want)
		}
	}
}

func TestUnwrapChain(t *testing.T) {
	inner := errors.NewInternalf(errors.CodeInternal, "db error")
	outer := errors.WrapInternalf(inner, errors.CodeInternal, "handler error")
	if !errors.Is(outer, inner) {
		t.Error("Is: outer should unwrap to inner")
	}
}

func TestHintPropagation(t *testing.T) {
	inner := errors.NewInvalidInputf(errors.CodeInvalidInput, "bad field").
		WithSuggestions("did you mean: `foo`")
	outer := errors.WrapInvalidInputf(inner, errors.CodeInvalidInput, "validation failed")
	j := errors.AsJSON(outer)
	if len(j.Suggestions) == 0 {
		t.Error("suggestions should propagate through Wrap")
	}
}

func TestAsJSONNullSafety(t *testing.T) {
	j := errors.AsJSON(errors.NewInternalf(errors.CodeInternal, "oops"))
	if j.Errors == nil {
		t.Error("Errors must be non-nil slice")
	}
	if j.Suggestions == nil {
		t.Error("Suggestions must be non-nil slice")
	}
}

func TestRetryDelayOf(t *testing.T) {
	err := errors.NewTooManyRequestsf(errors.CodeTooManyRequests, "slow down").
		WithRetryAfter(5 * time.Second)
	if d := errors.RetryDelayOf(err); d != 5*time.Second {
		t.Errorf("RetryDelayOf = %v, want 5s", d)
	}
	j := errors.AsJSON(err)
	if j.Retry == nil || j.Retry.Delay != 5000 {
		t.Errorf("JSON retry delay = %v, want 5000ms", j.Retry)
	}
}

func TestContextSentinels(t *testing.T) {
	wrapped := errors.WrapCanceledf(context.Canceled, errors.CodeCanceled, "request canceled")
	if !errors.Is(wrapped, context.Canceled) {
		t.Error("context.Canceled should survive WrapCanceledf")
	}
}

func TestLevenshtein(t *testing.T) {
	match, ok := errors.ClosestLevenshteinMatch("fieled", []string{"field", "filter", "format"})
	if !ok || match != "field" {
		t.Errorf("expected 'field', got %q (ok=%v)", match, ok)
	}
	_, ok = errors.ClosestLevenshteinMatch("xyz", []string{"field", "filter"})
	if ok {
		t.Error("expected no match for 'xyz'")
	}
}

func TestStacktrace(t *testing.T) {
	err := errors.NewInternalf(errors.CodeInternal, "boom")
	if err.Stacktrace() == "" {
		t.Error("stacktrace should be captured at creation")
	}
	overridden := err.WithStacktrace("custom\n\tfile.go:1\n")
	if overridden.Stacktrace() != "custom\n\tfile.go:1\n" {
		t.Error("WithStacktrace should override")
	}
}

func TestAstAsc(t *testing.T) {
	err := errors.NewNotFoundf(errors.CodeNotFound, "missing")
	if !errors.Ast(err, errors.TypeNotFound) {
		t.Error("Ast should match TypeNotFound")
	}
	if !errors.Asc(err, errors.CodeNotFound) {
		t.Error("Asc should match CodeNotFound")
	}
}
