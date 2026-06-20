package backend

import (
	"strings"
	"testing"
)

func TestNewRequest_Headers(t *testing.T) {
	req, err := NewRequest("GET", "https://example.com", "sk_test_abc", nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got := req.Header.Get("Authorization"); got != "Bearer sk_test_abc" {
		t.Errorf("Authorization = %q, want %q", got, "Bearer sk_test_abc")
	}
	if got := req.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	if ua := req.Header.Get("User-Agent"); !strings.Contains(ua, "paystack-sdk-go") {
		t.Errorf("User-Agent %q missing sdk identifier", ua)
	}
}

func TestNewRequest_CustomHeaders(t *testing.T) {
	opts := &RequestOptions{Headers: map[string]string{"Idempotency-Key": "idem-123"}}
	req, err := NewRequest("POST", "https://example.com", "sk_test_abc", nil, opts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got := req.Header.Get("Idempotency-Key"); got != "idem-123" {
		t.Errorf("Idempotency-Key = %q, want %q", got, "idem-123")
	}
}

func TestNewRequest_WithBody(t *testing.T) {
	body := []byte(`{"amount":1000}`)
	req, err := NewRequest("POST", "https://example.com", "sk_test", body, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if req.Body == nil {
		t.Error("expected non-nil body")
	}
}

func TestEncodeQueryParams_BasicTypes(t *testing.T) {
	type Params struct {
		Name    string `query:"name"`
		Count   int    `query:"count"`
		Active  bool   `query:"active"`
		Amount  float64 `query:"amount"`
	}
	q, err := EncodeQueryParams(&Params{Name: "test", Count: 5, Active: true, Amount: 1.5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(q, "name=test") {
		t.Errorf("missing name=test in %q", q)
	}
	if !strings.Contains(q, "count=5") {
		t.Errorf("missing count=5 in %q", q)
	}
	if !strings.Contains(q, "active=true") {
		t.Errorf("missing active=true in %q", q)
	}
	if !strings.Contains(q, "amount=1.5") {
		t.Errorf("missing amount=1.5 in %q", q)
	}
}

func TestEncodeQueryParams_PointerFields(t *testing.T) {
	type Params struct {
		Name    *string `query:"name,omitempty"`
		Missing *string `query:"missing,omitempty"`
	}
	name := "hello"
	q, err := EncodeQueryParams(&Params{Name: &name, Missing: nil})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(q, "name=hello") {
		t.Errorf("missing name=hello in %q", q)
	}
	if strings.Contains(q, "missing") {
		t.Errorf("nil pointer field should be omitted, got %q", q)
	}
}

func TestEncodeQueryParams_OmitemptyZeroValues(t *testing.T) {
	type Params struct {
		Page    int    `query:"page,omitempty"`
		Filter  string `query:"filter,omitempty"`
	}
	q, err := EncodeQueryParams(&Params{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q != "" {
		t.Errorf("expected empty query string for zero values with omitempty, got %q", q)
	}
}

func TestEncodeQueryParams_SliceFields(t *testing.T) {
	type Params struct {
		Channels []string `query:"channel"`
	}
	q, err := EncodeQueryParams(&Params{Channels: []string{"card", "bank"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(q, "channel=card") {
		t.Errorf("missing channel=card in %q", q)
	}
	if !strings.Contains(q, "channel=bank") {
		t.Errorf("missing channel=bank in %q", q)
	}
}

func TestEncodeQueryParams_NilInput(t *testing.T) {
	q, err := EncodeQueryParams(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q != "" {
		t.Errorf("expected empty string for nil input, got %q", q)
	}
}

func TestEncodeQueryParams_NonStructReturnsError(t *testing.T) {
	_, err := EncodeQueryParams("not-a-struct")
	if err == nil {
		t.Error("expected error for non-struct input")
	}
}

func TestEncodeQueryParams_DashTagSkipped(t *testing.T) {
	type Params struct {
		Secret string `query:"-"`
		Name   string `query:"name"`
	}
	q, err := EncodeQueryParams(&Params{Secret: "hidden", Name: "visible"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.Contains(q, "hidden") || strings.Contains(q, "Secret") {
		t.Errorf("dash-tagged field should be skipped, got %q", q)
	}
	if !strings.Contains(q, "name=visible") {
		t.Errorf("missing name=visible in %q", q)
	}
}
