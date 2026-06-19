package paymentrequests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreatePaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest" {
			t.Errorf("Expected path /paymentrequest, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request created", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Create(context.Background(), &CreatePaymentRequestRequest{Customer: "CUS_123", Amount: 1000})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.RequestCode != "PRQ_123" {
		t.Errorf("Expected request code PRQ_123, got %s", resp.Data.RequestCode)
	}
}

func TestListPaymentRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest" {
			t.Errorf("Expected path /paymentrequest, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment requests retrieved", "data": [{"request_code": "PRQ_123"}]}`)
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
		t.Errorf("Expected 1 request, got %d", len(resp.Data))
	}
}

func TestFetchPaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request retrieved", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "PRQ_123")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.RequestCode != "PRQ_123" {
		t.Errorf("Expected request code PRQ_123, got %s", resp.Data.RequestCode)
	}
}

func TestVerifyPaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/verify/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/verify/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request verified", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Verify(context.Background(), "PRQ_123")
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestSendNotification(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/notify/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/notify/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Notification sent", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.SendNotification(context.Background(), "PRQ_123")
	if err != nil {
		t.Fatalf("SendNotification failed: %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestTotal(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/total" {
			t.Errorf("Expected path /paymentrequest/total, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request totals", "data": {"total": [{"currency": "NGN", "amount": 50000}]}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Total(context.Background())
	if err != nil {
		t.Fatalf("Total failed: %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if len(resp.Data.Total) != 1 {
		t.Errorf("Expected 1 total entry, got %d", len(resp.Data.Total))
	}
	if resp.Data.Total[0].Currency != "NGN" {
		t.Errorf("Expected currency NGN, got %s", resp.Data.Total[0].Currency)
	}
}

func TestUpdatePaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request updated", "data": {"request_code": "PRQ_123", "amount": 2000}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Update(context.Background(), "PRQ_123", &UpdatePaymentRequestRequest{Amount: 2000})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.RequestCode != "PRQ_123" {
		t.Errorf("Expected request code PRQ_123, got %s", resp.Data.RequestCode)
	}
}

func TestArchivePaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/archive/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/archive/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request archived", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Archive(context.Background(), "PRQ_123")
	if err != nil {
		t.Fatalf("Archive failed: %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestFinalizePaymentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/paymentrequest/finalize/PRQ_123" {
			t.Errorf("Expected path /paymentrequest/finalize/PRQ_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Payment request finalized", "data": {"request_code": "PRQ_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Finalize(context.Background(), "PRQ_123")
	if err != nil {
		t.Fatalf("Finalize failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}
