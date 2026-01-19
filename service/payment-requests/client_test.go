package paymentrequests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreatePaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest" {
			t.Errorf("Expected path /paymentrequest, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request created", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Create(context.Background(), &CreatePaymentRequestRequest{Customer: "CUS_123", Amount: 1000})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.RequestCode != "PRQ_123" {
		t.Errorf("Expected request code PRQ_123, got %s", resp.Data.RequestCode)
	}
}

func TestListPaymentRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest" {
			t.Errorf("Expected path /paymentrequest, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment requests retrieved", "data": [{"request_code": "PRQ_123"}]}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.List(context.Background())
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 request, got %d", len(resp.Data))
	}
}

func TestFetchPaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request retrieved", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "PRQ_123")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.RequestCode != "PRQ_123" {
		t.Errorf("Expected request code PRQ_123, got %s", resp.Data.RequestCode)
	}
}

func TestVerifyPaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/verify/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/verify/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request verified", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Verify(context.Background(), "PRQ_123")
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestFinalizePaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/finalize/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/finalize/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request finalized", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Finalize(context.Background(), "PRQ_123")
	if err != nil {
		t.Fatalf("Finalize failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}
