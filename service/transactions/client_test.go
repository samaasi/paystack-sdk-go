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
	page := 1
	perPage := 10
	params := &ListTransactionParams{PerPage: &perPage, Page: &page}
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

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction/1" {
			t.Errorf("Expected path /transaction/1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Transaction retrieved","data":{"id":1,"reference":"ref_123","status":"success"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Fetch(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.ID != 1 {
		t.Errorf("Expected id 1, got %d", resp.Data.ID)
	}
}

func TestChargeAuthorization(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction/charge_authorization" {
			t.Errorf("Expected path /transaction/charge_authorization, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Charge attempted","data":{"id":2,"reference":"ref_charge","status":"success"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &ChargeAuthorizationRequest{
		Amount:            "10000",
		Email:             "customer@email.com",
		AuthorizationCode: "AUTH_abc123",
	}
	resp, err := client.ChargeAuthorization(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Reference != "ref_charge" {
		t.Errorf("Expected reference ref_charge, got %s", resp.Data.Reference)
	}
}

func TestGetTimeline(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction/timeline/ref_123" {
			t.Errorf("Expected path /transaction/timeline/ref_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Timeline retrieved","data":{"time_spent":9,"attempts":1,"success":true,"mobile":false,"errors":0,"channel":"card","history":[]}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.GetTimeline(context.Background(), "ref_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Data.Success {
		t.Errorf("Expected success true, got false")
	}
}

func TestTotals(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction/totals" {
			t.Errorf("Expected path /transaction/totals, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Transaction totals","data":{"total_transactions":10,"total_volume":50000,"total_volume_by_currency":[{"currency":"NGN","amount":50000}],"pending_transfers":0,"pending_transfers_by_currency":[]}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Totals(context.Background(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.TotalTransactions != 10 {
		t.Errorf("Expected 10 total transactions, got %d", resp.Data.TotalTransactions)
	}
}

func TestExport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction/export" {
			t.Errorf("Expected path /transaction/export, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Export successful","data":{"path":"https://files.paystack.co/exports/transactions.csv"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Export(context.Background(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Path == "" {
		t.Errorf("Expected export path, got empty string")
	}
}

func TestPartialDebit(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transaction/partial_debit" {
			t.Errorf("Expected path /transaction/partial_debit, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Charge attempted","data":{"id":3,"reference":"ref_partial","status":"success","amount":5000}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &PartialDebitRequest{
		AuthorizationCode: "AUTH_abc123",
		Currency:          paystackapi.CurrencyNGN,
		Amount:            "5000",
		Email:             "customer@email.com",
	}
	resp, err := client.PartialDebit(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Reference != "ref_partial" {
		t.Errorf("Expected reference ref_partial, got %s", resp.Data.Reference)
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

func TestInitialize_APIError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Validation failed",
			"errors":  map[string]interface{}{"amount": []string{"is required"}},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	_, err := client.Initialize(context.Background(), &InitializeRequest{Email: "a@b.com"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var apiErr *paystackapi.APIError
	if !isAPIError(err, &apiErr) {
		t.Fatalf("expected *paystackapi.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", apiErr.StatusCode)
	}
	if len(apiErr.Errors["amount"]) == 0 {
		t.Errorf("expected Errors[amount] populated, got %v", apiErr.Errors)
	}
}

func TestVerify_NotFoundError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Transaction reference not found",
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	_, err := client.Verify(context.Background(), "nonexistent_ref")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var apiErr *paystackapi.APIError
	if !isAPIError(err, &apiErr) {
		t.Fatalf("expected *paystackapi.APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404, got %d", apiErr.StatusCode)
	}
}

func TestInitialize_ServerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	_, err := client.Initialize(context.Background(), &InitializeRequest{Email: "a@b.com", Amount: "1000"})
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
}

func isAPIError(err error, target **paystackapi.APIError) bool {
	apiErr, ok := err.(*paystackapi.APIError)
	if ok {
		*target = apiErr
	}
	return ok
}
