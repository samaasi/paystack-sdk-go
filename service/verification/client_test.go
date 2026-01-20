package verification

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestResolveAccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/bank/resolve" {
			t.Errorf("Expected path /bank/resolve, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Account resolved","data":{"account_number":"1234567890","account_name":"Test Account"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.ResolveAccount(context.Background(), "1234567890", "058")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.AccountNumber != "1234567890" {
		t.Errorf("Expected account number 1234567890, got %s", resp.Data.AccountNumber)
	}
}

func TestValidateAccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/bank/validate" {
			t.Errorf("Expected path /bank/validate, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Account validated","data":{"verified":true}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &ValidateAccountRequest{AccountNumber: "1234567890"}
	resp, err := client.ValidateAccount(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Data.Verified {
		t.Errorf("Expected verified true, got false")
	}
}

func TestResolveCardBIN(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/decision/bin/539983" {
			t.Errorf("Expected path /decision/bin/539983, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"BIN resolved","data":{"bin":"539983","brand":"Mastercard"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.ResolveCardBIN(context.Background(), "539983")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Brand != "Mastercard" {
		t.Errorf("Expected brand Mastercard, got %s", resp.Data.Brand)
	}
}
