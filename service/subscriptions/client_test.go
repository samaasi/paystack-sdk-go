package subscriptions

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreateSubscription(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/subscription" {
			t.Errorf("Expected path /subscription, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":true,"message":"Subscription created","data":{"subscription_code":"SUB_vsyqdmlzble3uii"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &CreateSubscriptionRequest{
		Customer: "cus_12345",
		Plan:     "PLN_12345",
	}
	resp, err := client.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.SubscriptionCode != "SUB_vsyqdmlzble3uii" {
		t.Errorf("Expected subscription code SUB_vsyqdmlzble3uii, got %s", resp.Data.SubscriptionCode)
	}
}

func TestListSubscriptions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/subscription" {
			t.Errorf("Expected path /subscription, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Subscriptions retrieved","data":[{"subscription_code":"SUB_123"}]}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.List(context.Background(), 50, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 subscription, got %d", len(resp.Data))
	}
}

func TestFetchSubscription(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/subscription/SUB_123" {
			t.Errorf("Expected path /subscription/SUB_123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Subscription retrieved","data":{"subscription_code":"SUB_123"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.Fetch(context.Background(), "SUB_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.SubscriptionCode != "SUB_123" {
		t.Errorf("Expected subscription code SUB_123, got %s", resp.Data.SubscriptionCode)
	}
}

func TestDisableSubscription(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/subscription/disable" {
			t.Errorf("Expected path /subscription/disable, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Subscription disabled"}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &EnableDisableSubscriptionRequest{
		Code:  "SUB_123",
		Token: "token_123",
	}
	resp, err := client.Disable(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestGenerateLink(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/subscription/SUB_123/manage/link" {
			t.Errorf("Expected path /subscription/SUB_123/manage/link, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Link generated","data":{"link":"https://paystack.com/manage/SUB_123"}}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	resp, err := client.GenerateLink(context.Background(), "SUB_123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Data.Link == "" {
		t.Errorf("Expected a non-empty management link")
	}
}

func TestEnableSubscription(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/subscription/enable" {
			t.Errorf("Expected path /subscription/enable, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":true,"message":"Subscription enabled"}`))
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("sk_test_123", backend.WithBaseURL(ts.URL)))
	req := &EnableDisableSubscriptionRequest{
		Code:  "SUB_123",
		Token: "token_123",
	}
	_, err := client.Enable(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
