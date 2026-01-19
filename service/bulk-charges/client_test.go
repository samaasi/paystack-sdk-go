package bulkcharges

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestInitiate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/bulkcharge" {
			t.Errorf("Expected path /bulkcharge, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Bulk charge initiated", "data": {"batch_code": "BCH_123456", "reference": "REF_123", "id": 1, "status": "active"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := InitiateBulkChargeRequest{
		{Authorization: "AUTH_123", Amount: 1000, Reference: "ref_1"},
		{Authorization: "AUTH_456", Amount: 2000, Reference: "ref_2"},
	}

	resp, err := client.Initiate(context.Background(), req)
	if err != nil {
		t.Fatalf("Initiate failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.BatchCode != "BCH_123456" {
		t.Errorf("Expected batch code BCH_123456, got %s", resp.Data.BatchCode)
	}
}

func TestList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/bulkcharge" {
			t.Errorf("Expected path /bulkcharge, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Bulk charges retrieved", "data": [{"batch_code": "BCH_123456"}]}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.List(context.Background(), &ListBulkChargesParams{PerPage: 10})
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if len(resp.Data) == 0 {
		t.Errorf("Expected data, got empty")
	}
}

func TestFetch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/bulkcharge/BCH_123456" {
			t.Errorf("Expected path /bulkcharge/BCH_123456, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Bulk charge retrieved", "data": {"batch_code": "BCH_123456", "status": "active"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "BCH_123456")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.BatchCode != "BCH_123456" {
		t.Errorf("Expected batch code BCH_123456, got %s", resp.Data.BatchCode)
	}
}

func TestFetchCharges(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/bulkcharge/BCH_123456/charges" {
			t.Errorf("Expected path /bulkcharge/BCH_123456/charges, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Charges retrieved", "data": [{"reference": "ref_1", "status": "success"}]}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.FetchCharges(context.Background(), "BCH_123456", nil)
	if err != nil {
		t.Fatalf("FetchCharges failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if len(resp.Data) == 0 {
		t.Errorf("Expected data, got empty")
	}
}
