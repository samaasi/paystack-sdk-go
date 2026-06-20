package misc

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/v2/internal/backend"
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
	country, perPage, page := "NG", 10, 1
	resp, err := client.ListBanks(context.Background(), &ListBanksParams{
		Country: &country,
		PerPage: &perPage,
		Page:    &page,
	})
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

func TestListBanksError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Invalid key",
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("bad_key", backend.WithBaseURL(ts.URL)))
	_, err := client.ListBanks(context.Background(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
