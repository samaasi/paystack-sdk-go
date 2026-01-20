package transfers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestInitiate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transfer" {
			t.Errorf("Expected path /transfer, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Transfer initiated","data":{"reference":"ref_123","transfer_code":"TRF_123","status":"pending"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &InitiateRequest{
		Source:    "balance",
		Amount:    50000,
		Recipient: "RCP_123",
	}
	resp, err := client.Initiate(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.TransferCode != "TRF_123" {
		t.Errorf("Expected transfer code TRF_123, got %s", resp.Data.TransferCode)
	}
}

func TestFinalize(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transfer/finalize_transfer" {
			t.Errorf("Expected path /transfer/finalize_transfer, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Transfer finalized","data":{"reference":"ref_123","transfer_code":"TRF_123","status":"success"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &FinalizeRequest{
		TransferCode: "TRF_123",
		OTP:          "123456",
	}
	resp, err := client.Finalize(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Status != "success" {
		t.Errorf("Expected status success, got %s", resp.Data.Status)
	}
}

func TestList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transfer" {
			t.Errorf("Expected path /transfer, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Transfers retrieved","data":[{"reference":"ref_123","transfer_code":"TRF_123"}]}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	params := &ListTransferParams{PerPage: 10, Page: 1}
	resp, err := client.List(context.Background(), params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 transfer, got %d", len(resp.Data))
	}
}

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transfer/TRF_123" {
			t.Errorf("Expected path /transfer/TRF_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Transfer retrieved","data":{"reference":"ref_123","transfer_code":"TRF_123"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Fetch(context.Background(), "TRF_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.TransferCode != "TRF_123" {
		t.Errorf("Expected transfer code TRF_123, got %s", resp.Data.TransferCode)
	}
}

func TestVerify(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transfer/verify/ref_123" {
			t.Errorf("Expected path /transfer/verify/ref_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Transfer verified","data":{"reference":"ref_123","transfer_code":"TRF_123"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Verify(context.Background(), "ref_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Reference != "ref_123" {
		t.Errorf("Expected reference ref_123, got %s", resp.Data.Reference)
	}
}
