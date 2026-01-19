package integration

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestFetchPaymentSessionTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/integration/payment_session_timeout" {
			t.Errorf("Expected path /integration/payment_session_timeout, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment session timeout retrieved", "data": {"payment_session_timeout": 30}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.FetchPaymentSessionTimeout(context.Background())
	if err != nil {
		t.Fatalf("FetchPaymentSessionTimeout failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.PaymentSessionTimeout != 30 {
		t.Errorf("Expected timeout 30, got %d", resp.Data.PaymentSessionTimeout)
	}
}

func TestUpdatePaymentSessionTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/integration/payment_session_timeout" {
			t.Errorf("Expected path /integration/payment_session_timeout, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment session timeout updated", "data": {"payment_session_timeout": 60}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.UpdatePaymentSessionTimeout(context.Background(), 60)
	if err != nil {
		t.Fatalf("UpdatePaymentSessionTimeout failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.PaymentSessionTimeout != 60 {
		t.Errorf("Expected timeout 60, got %d", resp.Data.PaymentSessionTimeout)
	}
}
