package paystackapi

import (
	"context"
	"testing"
)

func TestWithIdempotencyKey_RoundTrip(t *testing.T) {
	ctx := WithIdempotencyKey(context.Background(), "idem-abc")
	if got := GetIdempotencyKey(ctx); got != "idem-abc" {
		t.Errorf("expected idem-abc, got %q", got)
	}
}

func TestGetIdempotencyKey_Missing(t *testing.T) {
	got := GetIdempotencyKey(context.Background())
	if got != "" {
		t.Errorf("expected empty string for missing key, got %q", got)
	}
}

func TestWithCustomHeader_RoundTrip(t *testing.T) {
	ctx := WithCustomHeader(context.Background(), "X-Trace-ID", "trace-001")
	headers := GetCustomHeaders(ctx)
	if headers["X-Trace-ID"] != "trace-001" {
		t.Errorf("expected X-Trace-ID = trace-001, got %q", headers["X-Trace-ID"])
	}
}

func TestWithCustomHeader_Multiple(t *testing.T) {
	ctx := context.Background()
	ctx = WithCustomHeader(ctx, "X-A", "val-a")
	ctx = WithCustomHeader(ctx, "X-B", "val-b")
	headers := GetCustomHeaders(ctx)
	if headers["X-A"] != "val-a" {
		t.Errorf("expected X-A = val-a, got %q", headers["X-A"])
	}
	if headers["X-B"] != "val-b" {
		t.Errorf("expected X-B = val-b, got %q", headers["X-B"])
	}
}

func TestWithCustomHeader_DoesNotMutateParent(t *testing.T) {
	parent := context.Background()
	child := WithCustomHeader(parent, "X-Only-Child", "yes")

	parentHeaders := GetCustomHeaders(parent)
	if _, ok := parentHeaders["X-Only-Child"]; ok {
		t.Error("WithCustomHeader mutated the parent context")
	}
	childHeaders := GetCustomHeaders(child)
	if childHeaders["X-Only-Child"] != "yes" {
		t.Error("child context missing expected header")
	}
}

func TestGetCustomHeaders_Missing(t *testing.T) {
	headers := GetCustomHeaders(context.Background())
	if headers != nil {
		t.Errorf("expected nil for missing headers, got %v", headers)
	}
}
