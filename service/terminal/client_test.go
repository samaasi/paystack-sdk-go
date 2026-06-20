package terminal

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestSendEvent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/terminal/TERM_123/event" {
			t.Errorf("Expected path /terminal/TERM_123/event, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Event sent","data":{"id":"EVT_123","status":"sent"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &SendEventRequest{
		Type:   "invoice",
		Action: "process",
		Data:   map[string]interface{}{"amount": 1000},
	}
	resp, err := client.SendEvent(context.Background(), "TERM_123", req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.ID != "EVT_123" {
		t.Errorf("Expected event ID EVT_123, got %s", resp.Data.ID)
	}
}

func TestFetchEventStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/terminal/TERM_123/event/EVT_123" {
			t.Errorf("Expected path /terminal/TERM_123/event/EVT_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Event status retrieved","data":{"id":"EVT_123","status":"delivered"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.FetchEventStatus(context.Background(), "TERM_123", "EVT_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Status != "delivered" {
		t.Errorf("Expected status delivered, got %s", resp.Data.Status)
	}
}

func TestListTerminals(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/terminal" {
			t.Errorf("Expected path /terminal, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Terminals retrieved","data":[{"terminal_id":"TERM_123"}]}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.List(context.Background(), 50, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 terminal, got %d", len(resp.Data))
	}
}

func TestFetchTerminal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/terminal/TERM_123" {
			t.Errorf("Expected path /terminal/TERM_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Terminal retrieved","data":{"terminal_id":"TERM_123","name":"Store Terminal","status":"active"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Fetch(context.Background(), "TERM_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.TerminalID != "TERM_123" {
		t.Errorf("Expected terminal ID TERM_123, got %s", resp.Data.TerminalID)
	}
	if resp.Data.Status != "active" {
		t.Errorf("Expected status active, got %s", resp.Data.Status)
	}
}

func TestUpdateTerminal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/terminal/TERM_123" {
			t.Errorf("Expected path /terminal/TERM_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Terminal updated","data":{"terminal_id":"TERM_123","name":"Updated Terminal"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Update(context.Background(), "TERM_123", &UpdateTerminalRequest{Name: "Updated Terminal"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Name != "Updated Terminal" {
		t.Errorf("Expected name Updated Terminal, got %s", resp.Data.Name)
	}
}

func TestCommission(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/terminal/commission_device" {
			t.Errorf("Expected path /terminal/commission_device, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Device commissioned","data":{"terminal_id":"TERM_123","serial":"SN_12345"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Commission(context.Background(), "SN_12345")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestDecommission(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/terminal/decommission_device" {
			t.Errorf("Expected path /terminal/decommission_device, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Device decommissioned","data":{"terminal_id":"TERM_123","serial":"SN_12345"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Decommission(context.Background(), "SN_12345")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestFetchTerminalPresence(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/terminal/TERM_123/presence" {
			t.Errorf("Expected path /terminal/TERM_123/presence, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Terminal presence","data":{"online":true,"available":true}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.FetchPresence(context.Background(), "TERM_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Data.Online {
		t.Errorf("Expected terminal to be online")
	}
}
