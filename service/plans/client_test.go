package plans

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreatePlan(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/plan" {
			t.Errorf("Expected path /plan, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Plan created", "data": {"name": "Test Plan", "plan_code": "PLN_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Create(context.Background(), &CreatePlanRequest{Name: "Test Plan", Amount: 5000, Interval: "monthly"})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.PlanCode != "PLN_123" {
		t.Errorf("Expected plan code PLN_123, got %s", resp.Data.PlanCode)
	}
}

func TestListPlans(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/plan" {
			t.Errorf("Expected path /plan, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Plans retrieved", "data": [{"plan_code": "PLN_123"}]}`)
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
		t.Errorf("Expected 1 plan, got %d", len(resp.Data))
	}
}

func TestFetchPlan(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/plan/PLN_123" {
			t.Errorf("Expected path /plan/PLN_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Plan retrieved", "data": {"plan_code": "PLN_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "PLN_123")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.PlanCode != "PLN_123" {
		t.Errorf("Expected plan code PLN_123, got %s", resp.Data.PlanCode)
	}
}

func TestUpdatePlan(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/plan/PLN_123" {
			t.Errorf("Expected path /plan/PLN_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Plan updated", "data": {"plan_code": "PLN_123", "name": "Updated Plan"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Update(context.Background(), "PLN_123", &UpdatePlanRequest{Name: "Updated Plan"})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.Name != "Updated Plan" {
		t.Errorf("Expected name Updated Plan, got %s", resp.Data.Name)
	}
}
