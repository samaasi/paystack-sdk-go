package settlements

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestListSettlements(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/settlement" {
			t.Errorf("Expected path /settlement, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Settlements retrieved", "data": [{"id": 123}]}`)
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
		t.Errorf("Expected 1 settlement, got %d", len(resp.Data))
	}
}

func TestFetchSettlementTransactions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/settlement/123/transactions" {
			t.Errorf("Expected path /settlement/123/transactions, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Settlement transactions retrieved", "data": [{"id": 456}]}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.FetchTransactions(context.Background(), "123")
	if err != nil {
		t.Fatalf("FetchTransactions failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(resp.Data))
	}
}
