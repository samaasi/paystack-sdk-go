package transferRecipients

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transferrecipient" {
			t.Errorf("Expected /transferrecipient, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Recipient created",
			"data": map[string]interface{}{
				"recipient_code": "RCP_xyz",
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.Create(context.Background(), &CreateRequest{Type: "nuban", Name: "John Doe", AccountNumber: "123", BankCode: "011"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if resp.Data.RecipientCode != "RCP_xyz" {
		t.Errorf("Expected recipient code RCP_xyz, got %s", resp.Data.RecipientCode)
	}
}

func TestList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transferrecipient" {
			t.Errorf("Expected /transferrecipient, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Recipients retrieved",
			"data": []map[string]interface{}{
				{"recipient_code": "RCP_xyz"},
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.List(context.Background(), 10, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 recipient, got %d", len(resp.Data))
	}
}

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transferrecipient/RCP_xyz" {
			t.Errorf("Expected /transferrecipient/RCP_xyz, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Recipient retrieved",
			"data": map[string]interface{}{
				"recipient_code": "RCP_xyz",
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.Fetch(context.Background(), "RCP_xyz")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if resp.Data.RecipientCode != "RCP_xyz" {
		t.Errorf("Expected recipient code RCP_xyz, got %s", resp.Data.RecipientCode)
	}
}

func TestUpdate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/transferrecipient/RCP_xyz" {
			t.Errorf("Expected /transferrecipient/RCP_xyz, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Recipient updated",
			"data": map[string]interface{}{
				"recipient_code": "RCP_xyz",
				"name":           "Jane Doe",
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.Update(context.Background(), "RCP_xyz", &UpdateRequest{Name: "Jane Doe"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Name != "Jane Doe" {
		t.Errorf("Expected name Jane Doe, got %s", resp.Data.Name)
	}
}

func TestBulkCreate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transferrecipient/bulk" {
			t.Errorf("Expected /transferrecipient/bulk, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Recipients created",
			"data": []map[string]interface{}{
				{"recipient_code": "RCP_aaa"},
				{"recipient_code": "RCP_bbb"},
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	req := &BatchCreateRequest{
		Batch: []CreateRequest{
			{Type: "nuban", Name: "Alice", AccountNumber: "111", BankCode: "011"},
			{Type: "nuban", Name: "Bob", AccountNumber: "222", BankCode: "011"},
		},
	}
	resp, err := client.BulkCreate(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(resp.Data) != 2 {
		t.Errorf("Expected 2 recipients, got %d", len(resp.Data))
	}
	if resp.Data[0].RecipientCode != "RCP_aaa" {
		t.Errorf("Expected first recipient code RCP_aaa, got %s", resp.Data[0].RecipientCode)
	}
}

func TestDelete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if r.URL.Path != "/transferrecipient/RCP_xyz" {
			t.Errorf("Expected /transferrecipient/RCP_xyz, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Recipient deleted",
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.Delete(context.Background(), "RCP_xyz")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
}
