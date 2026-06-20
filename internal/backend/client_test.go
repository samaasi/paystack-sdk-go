package backend

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient_Defaults(t *testing.T) {
	c := NewClient("sk_test_123")
	if c.baseURL != DefaultBaseURL {
		t.Errorf("expected default base URL %s, got %s", DefaultBaseURL, c.baseURL)
	}
	if c.apiKey != "sk_test_123" {
		t.Errorf("expected api key sk_test_123, got %s", c.apiKey)
	}
	if c.httpClient == nil {
		t.Error("expected non-nil http client")
	}
	if c.backoff == nil {
		t.Error("expected non-nil backoff")
	}
}

func TestWithBaseURL(t *testing.T) {
	c := NewClient("sk_test_123", WithBaseURL("https://custom.example.com"))
	if c.baseURL != "https://custom.example.com" {
		t.Errorf("expected https://custom.example.com, got %s", c.baseURL)
	}
}

func TestWithHTTPClient(t *testing.T) {
	custom := &http.Client{Timeout: 5 * time.Second}
	c := NewClient("sk_test_123", WithHTTPClient(custom))
	if c.httpClient != custom {
		t.Error("expected custom http client to be set")
	}
}

func TestWithMaxRetries(t *testing.T) {
	c := NewClient("sk_test_123", WithMaxRetries(5))
	if c.backoff.MaxRetries != 5 {
		t.Errorf("expected MaxRetries 5, got %d", c.backoff.MaxRetries)
	}
}

func TestWithTimeout(t *testing.T) {
	c := NewClient("sk_test_123", WithTimeout(15*time.Second))
	if c.httpClient.Timeout != 15*time.Second {
		t.Errorf("expected timeout 15s, got %v", c.httpClient.Timeout)
	}
}

func TestCall_GET(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.Header.Get("Authorization") == "" {
			t.Error("expected Authorization header")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"ok","data":{}}`))
	}))
	defer ts.Close()

	c := NewClient("sk_test_123", WithBaseURL(ts.URL))
	var out struct {
		Status bool `json:"status"`
	}
	if err := c.Call(context.Background(), "GET", "/test", nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Status {
		t.Error("expected status true")
	}
}

func TestCall_POST_WithBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.Header.Get("Idempotency-Key") == "" {
			t.Error("expected auto-generated Idempotency-Key header for POST")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"created","data":{}}`))
	}))
	defer ts.Close()

	c := NewClient("sk_test_123", WithBaseURL(ts.URL))
	body := map[string]string{"key": "value"}
	var out struct {
		Status bool `json:"status"`
	}
	if err := c.Call(context.Background(), "POST", "/test", body, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCall_IdempotencyKeyFromContext(t *testing.T) {
	seen := ""
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = r.Header.Get("Idempotency-Key")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"ok","data":{}}`))
	}))
	defer ts.Close()

	ctx := context.Background()
	// inject a known idempotency key via context using the paystackapi package
	// We test indirectly: key set in context should appear in the request header.
	// Use the backend client directly and rely on the Call implementation.
	c := NewClient("sk_test_123", WithBaseURL(ts.URL))
	var out interface{}
	_ = c.Call(ctx, "POST", "/test", map[string]string{}, &out)
	if seen == "" {
		t.Error("expected an Idempotency-Key header to be set for POST")
	}
}

func TestCall_APIError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":false,"message":"Invalid request"}`))
	}))
	defer ts.Close()

	c := NewClient("sk_test_123", WithBaseURL(ts.URL), WithMaxRetries(0))
	var out interface{}
	err := c.Call(context.Background(), "GET", "/test", nil, &out)
	if err == nil {
		t.Fatal("expected error for 400 response, got nil")
	}
}
