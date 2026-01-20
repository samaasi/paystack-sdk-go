package transactions

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestInitialize(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction/initialize" {
			t.Errorf("Expected path /transaction/initialize, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Authorization URL created","data":{"authorization_url":"https://checkout.paystack.com/access_code","access_code":"access_code","reference":"reference"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &InitializeRequest{
		Amount: "10000",
		Email:  "customer@email.com",
	}
	resp, err := client.Initialize(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.AccessCode != "access_code" {
		t.Errorf("Expected access code access_code, got %s", resp.Data.AccessCode)
	}
}

func TestVerify(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction/verify/ref_123" {
			t.Errorf("Expected path /transaction/verify/ref_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Verification successful","data":{"id":1,"status":"success","reference":"ref_123","amount":10000}}`))
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
	if resp.Data.Status != "success" {
		t.Errorf("Expected status success, got %s", resp.Data.Status)
	}
}

func TestList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction" {
			t.Errorf("Expected path /transaction, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Transactions retrieved","data":[{"id":1,"reference":"ref_123"}]}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	params := &ListTransactionParams{PerPage: 10, Page: 1}
	resp, err := client.List(context.Background(), params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(resp.Data))
	}
}
