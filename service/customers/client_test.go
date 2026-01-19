package customers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/customer" {
			t.Errorf("Expected path /customer, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Customer created", "data": {"email": "test@example.com", "customer_code": "CUS_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &CreateCustomerRequest{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
	}

	resp, err := client.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", resp.Data.Email)
	}
}

func TestList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/customer" {
			t.Errorf("Expected path /customer, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Customers retrieved", "data": [{"email": "test@example.com"}]}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.List(context.Background(), &ListCustomersParams{PerPage: 10})
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
		if r.URL.Path != "/customer/CUS_123" {
			t.Errorf("Expected path /customer/CUS_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Customer retrieved", "data": {"email": "test@example.com", "customer_code": "CUS_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "CUS_123")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.CustomerCode != "CUS_123" {
		t.Errorf("Expected customer code CUS_123, got %s", resp.Data.CustomerCode)
	}
}

func TestUpdate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/customer/CUS_123" {
			t.Errorf("Expected path /customer/CUS_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Customer updated", "data": {"email": "test@example.com", "first_name": "Jane"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &UpdateCustomerRequest{
		FirstName: "Jane",
	}

	resp, err := client.Update(context.Background(), "CUS_123", req)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.FirstName != "Jane" {
		t.Errorf("Expected first name Jane, got %s", resp.Data.FirstName)
	}
}

func TestValidate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/customer/CUS_123/identification" {
			t.Errorf("Expected path /customer/CUS_123/identification, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Customer validated", "data": true}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &ValidateCustomerRequest{
		Country: "NG",
		Type:    "bvn",
		Value:   "12345678901",
	}

	resp, err := client.Validate(context.Background(), "CUS_123", req)
	if err != nil {
		t.Fatalf("Validate failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestWhitelist(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/customer/set_risk_action" {
			t.Errorf("Expected path /customer/set_risk_action, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Customer whitelisted", "data": {"customer_code": "CUS_123", "risk_action": "allow"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Whitelist(context.Background(), "CUS_123")
	if err != nil {
		t.Fatalf("Whitelist failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.RiskAction != "allow" {
		t.Errorf("Expected risk action allow, got %s", resp.Data.RiskAction)
	}
}
