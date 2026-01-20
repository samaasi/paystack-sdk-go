package misc

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestListBanks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/bank" {
			t.Errorf("Expected /bank, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Banks retrieved",
			"data": []map[string]interface{}{
				{"name": "Test Bank", "code": "000"},
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.ListBanks(context.Background(), "NG", 10, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 bank, got %d", len(resp.Data))
	}
}

func TestListCountries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/country" {
			t.Errorf("Expected /country, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Countries retrieved",
			"data": []map[string]interface{}{
				{"name": "Nigeria", "iso_code": "NG"},
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.ListCountries(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 country, got %d", len(resp.Data))
	}
}

func TestListStates(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/address_verification/states" {
			t.Errorf("Expected /address_verification/states, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "States retrieved",
			"data": []map[string]interface{}{
				{"name": "Lagos", "slug": "lagos"},
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.ListStates(context.Background(), "NG")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 state, got %d", len(resp.Data))
	}
}

func TestResolveCardBIN(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/decision/bin/123456" {
			t.Errorf("Expected /decision/bin/123456, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "BIN resolved",
			"data": map[string]interface{}{
				"bin": "123456",
				"brand": "Visa",
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.ResolveCardBIN(context.Background(), "123456")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if resp.Data.Bin != "123456" {
		t.Errorf("Expected BIN 123456, got %s", resp.Data.Bin)
	}
}

func TestResolveAccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/bank/resolve" {
			t.Errorf("Expected /bank/resolve, got %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("account_number") != "1234567890" {
			t.Errorf("Expected account_number 1234567890, got %s", q.Get("account_number"))
		}
		if q.Get("bank_code") != "011" {
			t.Errorf("Expected bank_code 011, got %s", q.Get("bank_code"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Account resolved",
			"data": map[string]interface{}{
				"account_number": "1234567890",
				"account_name": "Test Account",
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.ResolveAccount(context.Background(), "1234567890", "011")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if resp.Data.AccountNumber != "1234567890" {
		t.Errorf("Expected account number 1234567890, got %s", resp.Data.AccountNumber)
	}
}
