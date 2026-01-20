package subaccounts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreateSubaccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/subaccount" {
			t.Errorf("Expected path /subaccount, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":true,"message":"Subaccount created","data":{"subaccount_code":"ACCT_45678"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &CreateSubaccountRequest{
		BusinessName:     "Cheese Sticks",
		SettlementBank:   "044",
		AccountNumber:    "0193274682",
		PercentageCharge: 18.2,
	}
	resp, err := client.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.SubaccountCode != "ACCT_45678" {
		t.Errorf("Expected subaccount code ACCT_45678, got %s", resp.Data.SubaccountCode)
	}
}

func TestListSubaccounts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/subaccount" {
			t.Errorf("Expected path /subaccount, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Subaccounts retrieved","data":[{"subaccount_code":"ACCT_123"}]}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.List(context.Background(), 50, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 subaccount, got %d", len(resp.Data))
	}
}

func TestFetchSubaccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/subaccount/ACCT_123" {
			t.Errorf("Expected path /subaccount/ACCT_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Subaccount retrieved","data":{"subaccount_code":"ACCT_123"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Fetch(context.Background(), "ACCT_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.SubaccountCode != "ACCT_123" {
		t.Errorf("Expected subaccount code ACCT_123, got %s", resp.Data.SubaccountCode)
	}
}

func TestUpdateSubaccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/subaccount/ACCT_123" {
			t.Errorf("Expected path /subaccount/ACCT_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Subaccount updated","data":{"subaccount_code":"ACCT_123"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &UpdateSubaccountRequest{
		BusinessName: "Cheese Balls",
	}
	resp, err := client.Update(context.Background(), "ACCT_123", req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.SubaccountCode != "ACCT_123" {
		t.Errorf("Expected subaccount code ACCT_123, got %s", resp.Data.SubaccountCode)
	}
}
