package paystackapi

import (
	"context"
)

type contextKey string

const (
	idempotencyKey   contextKey = "idempotency_key"
	customHeadersKey contextKey = "custom_headers"
)

// WithIdempotencyKey returns a copy of parent context with the idempotency key set.
func WithIdempotencyKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, idempotencyKey, key)
}

// WithCustomHeader returns a copy of parent context with a custom header set.
// Note: This adds to any existing custom headers in the context.
func WithCustomHeader(ctx context.Context, key, value string) context.Context {
	headers, _ := ctx.Value(customHeadersKey).(map[string]string)
	newHeaders := make(map[string]string)
	for k, v := range headers {
		newHeaders[k] = v
	}
	newHeaders[key] = value
	return context.WithValue(ctx, customHeadersKey, newHeaders)
}

// GetIdempotencyKey retrieves the idempotency key from the context.
func GetIdempotencyKey(ctx context.Context) string {
	val, _ := ctx.Value(idempotencyKey).(string)
	return val
}

// GetCustomHeaders retrieves custom headers from the context.
func GetCustomHeaders(ctx context.Context) map[string]string {
	val, _ := ctx.Value(customHeadersKey).(map[string]string)
	return val
}
