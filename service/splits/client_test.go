package splits

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreateSplit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/split" {
			t.Errorf("Expected path /split, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Split created", "data": {"name": "Test Split", "split_code": "SPL_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Create(context.Background(), &CreateSplitRequest{Name: "Test Split", Type: "percentage", Currency: "NGN"})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.SplitCode != "SPL_123" {
		t.Errorf("Expected split code SPL_123, got %s", resp.Data.SplitCode)
	}
}

func TestListSplits(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/split" {
			t.Errorf("Expected path /split, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Splits retrieved", "data": [{"split_code": "SPL_123"}]}`)
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
		t.Errorf("Expected 1 split, got %d", len(resp.Data))
	}
}

func TestFetchSplit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/split/SPL_123" {
			t.Errorf("Expected path /split/SPL_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Split retrieved", "data": {"split_code": "SPL_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "SPL_123")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.SplitCode != "SPL_123" {
		t.Errorf("Expected split code SPL_123, got %s", resp.Data.SplitCode)
	}
}

func TestUpdateSplit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/split/SPL_123" {
			t.Errorf("Expected path /split/SPL_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Split updated", "data": {"split_code": "SPL_123", "name": "Updated Split"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Update(context.Background(), "SPL_123", &UpdateSplitRequest{Name: "Updated Split", Active: true})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.Name != "Updated Split" {
		t.Errorf("Expected name Updated Split, got %s", resp.Data.Name)
	}
}

func TestAddSubaccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/split/SPL_123/subaccount/add" {
			t.Errorf("Expected path /split/SPL_123/subaccount/add, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Subaccount added", "data": {"split_code": "SPL_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.AddSubaccount(context.Background(), "SPL_123", &SubaccountRequest{Subaccount: "SUB_123", Share: 20})
	if err != nil {
		t.Fatalf("AddSubaccount failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestRemoveSubaccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/split/SPL_123/subaccount/remove" {
			t.Errorf("Expected path /split/SPL_123/subaccount/remove, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Subaccount removed", "data": {"split_code": "SPL_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.RemoveSubaccount(context.Background(), "SPL_123", &SubaccountRequest{Subaccount: "SUB_123"})
	if err != nil {
		t.Fatalf("RemoveSubaccount failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}
