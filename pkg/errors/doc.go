// Package errors is the single error-handling layer for pragmata.
// It is adapted from SigNoz's pkg/errors and must be used everywhere
// instead of stdlib errors or fmt.Errorf.
//
// # Choosing New vs Wrap
//
// Use [New] / [Newf] / type-specific shortcuts (e.g. [NewNotFoundf]) when
// there is no underlying error to preserve — you are originating an error:
//
//	// validation: no external error exists
//	if name == "" {
//	    return errors.NewInvalidInputf(errors.CodeInvalidInput, "name is required")
//	}
//
//	// explicit sentinel: you checked a condition, not an error return
//	if src == nil {
//	    return errors.NewNotFoundf(errors.CodeNotFound, "source %q not found", name)
//	}
//
// Use [Wrap] / [Wrapf] / type-specific shortcuts (e.g. [WrapInternalf]) ONLY
// at system boundaries — when you received an untyped error from an external
// library or stdlib and need to assign a type, code, and message:
//
//	row, err := db.QueryRow(ctx, q)
//	if err != nil {
//	    return errors.WrapInternalf(err, errors.CodeInternal, "query sources")
//	}
//
//	if errors.Is(err, sql.ErrNoRows) {
//	    return errors.WrapNotFoundf(err, errors.CodeNotFound, "pipe %q not found", name)
//	}
//
// # Never re-wrap an already-typed error
//
// If the error comes from another internal function (one that already uses this
// package), do NOT wrap it again — that overrides its type and code. Use
// [WithAdditionalf] to attach context without altering type, code, or message:
//
//	// BAD — overrides the type/code set by an internal function
//	exec, err := pipevisitor.Visit(content, opts)
//	if err != nil {
//	    return errors.WrapInternalf(err, errors.CodeInternal, "parse pipe %q", name)
//	}
//
//	// GOOD — preserves type/code, appends caller context
//	exec, err := pipevisitor.Visit(content, opts)
//	if err != nil {
//	    return errors.WithAdditionalf(err, "parse pipe %q", name)
//	}
//
// # Never use fmt.Errorf or stdlib errors.New
//
// Both produce untyped errors: no HTTP status, no code, no stack trace, no
// suggestions, no retry metadata. The HTTP layer cannot infer a status code
// and will default to 500.
//
//	// BAD — loses all structure
//	return fmt.Errorf("query sources: %w", err)
//	return fmt.Errorf("name is required")
//
//	// GOOD
//	return errors.WrapInternalf(err, errors.CodeInternal, "query sources")
//	return errors.NewInvalidInputf(errors.CodeInvalidInput, "name is required")
//
// # HTTP layer
//
// Handlers pass errors directly to [render.ErrorFrom]. The HTTP status code
// is derived from the error type — never set it manually:
//
//	result, err := h.module.Execute(ctx, name)
//	if err != nil {
//	    render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "execute %q", name))
//	    return
//	}
//
// # Enriching errors
//
// Attach user-facing suggestions when the caller can act on them:
//
//	return errors.NewInvalidInputf(errors.CodeInvalidInput, "unknown field %q", field).
//	    WithSuggestions(errors.NewSuggestionsOnLevenshteinDistance(field, errors.NounFields, validFields)...)
//
// Attach retry delay for rate-limit and transient errors:
//
//	return errors.NewTooManyRequestsf(errors.CodeTooManyRequests, "rate limit exceeded").
//	    WithRetryAfter(30 * time.Second)
//
// # Inspecting errors
//
// Use [Ast] and [Asc] to branch on error type or code in middleware or tests:
//
//	if errors.Ast(err, errors.TypeNotFound) { ... }
//	if errors.Asc(err, errors.CodeNotFound) { ... }
package errors
