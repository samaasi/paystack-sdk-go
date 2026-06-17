package transactions

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
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

func TestInitializeWithMetadata(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var payload map[string]interface{}
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("Failed to parse request body: %v", err)
		}
		meta, ok := payload["metadata"].(map[string]interface{})
		if !ok {
			t.Fatal("Expected metadata in request body")
		}
		if meta["cart_id"] != "398" {
			t.Errorf("Expected cart_id 398, got %v", meta["cart_id"])
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Authorization URL created","data":{"authorization_url":"https://checkout.paystack.com/test","access_code":"test_code","reference":"ref_meta"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &InitializeRequest{
		Amount: "50000",
		Email:  "meta@test.com",
		Metadata: paystackapi.Metadata{
			"cart_id": "398",
			"custom_fields": []map[string]interface{}{
				{
					"display_name":  "Invoice ID",
					"variable_name": "invoice_id",
					"value":         "INV-001",
				},
			},
		},
	}
	resp, err := client.Initialize(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Reference != "ref_meta" {
		t.Errorf("Expected reference ref_meta, got %s", resp.Data.Reference)
	}
}

func TestVerifyWithMetadata(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Verification successful","data":{"id":1,"status":"success","reference":"ref_meta","amount":50000,"metadata":{"cart_id":"398","custom_fields":[{"display_name":"Invoice ID","variable_name":"invoice_id","value":"INV-001"}]}}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Verify(context.Background(), "ref_meta")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Metadata == nil {
		t.Fatal("Expected metadata in response, got nil")
	}
	if resp.Data.Metadata["cart_id"] != "398" {
		t.Errorf("Expected cart_id 398, got %v", resp.Data.Metadata["cart_id"])
	}
}
