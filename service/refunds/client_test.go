package refunds

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreateRefund(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/refund" {
			t.Errorf("Expected path /refund, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Refund created", "data": {"id": 123}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Create(context.Background(), &CreateRefundRequest{Transaction: "TRX_123"})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.ID != 123 {
		t.Errorf("Expected ID 123, got %d", resp.Data.ID)
	}
}

func TestListRefunds(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/refund" {
			t.Errorf("Expected path /refund, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Refunds retrieved", "data": [{"id": 123}]}`)
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
		t.Errorf("Expected 1 refund, got %d", len(resp.Data))
	}
}

func TestFetchRefund(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/refund/123" {
			t.Errorf("Expected path /refund/123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Refund retrieved", "data": {"id": 123}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "123")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.ID != 123 {
		t.Errorf("Expected ID 123, got %d", resp.Data.ID)
	}
}
