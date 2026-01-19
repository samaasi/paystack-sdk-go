package disputes

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/dispute" {
			t.Errorf("Expected path /dispute, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Disputes retrieved", "data": [{"id": 123}]}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.List(context.Background(), &ListDisputesParams{PerPage: 10})
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
		if r.URL.Path != "/dispute/123" {
			t.Errorf("Expected path /dispute/123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Dispute retrieved", "data": {"id": 123}}`)
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
		t.Errorf("Expected id 123, got %d", resp.Data.ID)
	}
}

func TestListTransactionDisputes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/dispute/transaction/123" {
			t.Errorf("Expected path /dispute/transaction/123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Disputes retrieved", "data": [{"id": 123}]}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.ListTransactionDisputes(context.Background(), "123")
	if err != nil {
		t.Fatalf("ListTransactionDisputes failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestUpdate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/dispute/123" {
			t.Errorf("Expected path /dispute/123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Dispute updated", "data": {"id": 123}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &UpdateDisputeRequest{
		RefundAmount: "1000",
	}

	resp, err := client.Update(context.Background(), "123", req)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestAddEvidence(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/dispute/123/evidence" {
			t.Errorf("Expected path /dispute/123/evidence, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Evidence added", "data": {"id": 123}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &AddEvidenceRequest{
		CustomerEmail: "test@example.com",
	}

	resp, err := client.AddEvidence(context.Background(), "123", req)
	if err != nil {
		t.Fatalf("AddEvidence failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestGetUploadURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/dispute/123/upload_url" {
			t.Errorf("Expected path /dispute/123/upload_url, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Upload URL retrieved", "data": {"signedUrl": "http://example.com"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.GetUploadURL(context.Background(), "123", "test.jpg")
	if err != nil {
		t.Fatalf("GetUploadURL failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestResolve(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/dispute/123/resolve" {
			t.Errorf("Expected path /dispute/123/resolve, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Dispute resolved", "data": {"id": 123}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &ResolveDisputeRequest{
		Resolution: "merchant-accepted",
	}

	resp, err := client.Resolve(context.Background(), "123", req)
	if err != nil {
		t.Fatalf("Resolve failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestExport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/dispute/export" {
			t.Errorf("Expected path /dispute/export, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Disputes exported", "data": {"path": "http://example.com/export"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Export(context.Background(), nil)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}
