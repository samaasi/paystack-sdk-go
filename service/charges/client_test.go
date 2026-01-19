package charges

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
		if r.URL.Path != "/charge" {
			t.Errorf("Expected path /charge, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Charge initiated", "data": {"reference": "ref_123", "status": "pending"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &CreateChargeRequest{
		Email:  "test@example.com",
		Amount: "1000",
		Bank: &BankSource{
			Code:          "057",
			AccountNumber: "0000000000",
		},
	}

	resp, err := client.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.Reference != "ref_123" {
		t.Errorf("Expected reference ref_123, got %s", resp.Data.Reference)
	}
}

func TestSubmitPIN(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/charge/submit_pin" {
			t.Errorf("Expected path /charge/submit_pin, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "PIN submitted", "data": {"reference": "ref_123", "status": "success"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &SubmitPINRequest{
		Pin:       "1234",
		Reference: "ref_123",
	}

	resp, err := client.SubmitPIN(context.Background(), req)
	if err != nil {
		t.Fatalf("SubmitPIN failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestSubmitOTP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/charge/submit_otp" {
			t.Errorf("Expected path /charge/submit_otp, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "OTP submitted", "data": {"reference": "ref_123", "status": "success"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	req := &SubmitOTPRequest{
		OTP:       "123456",
		Reference: "ref_123",
	}

	resp, err := client.SubmitOTP(context.Background(), req)
	if err != nil {
		t.Fatalf("SubmitOTP failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestCheckPending(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/charge/ref_123" {
			t.Errorf("Expected path /charge/ref_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Charge checked", "data": {"reference": "ref_123", "status": "success"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.CheckPending(context.Background(), "ref_123")
	if err != nil {
		t.Fatalf("CheckPending failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}
