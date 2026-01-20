package transferControl

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCheckBalance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/balance" {
			t.Errorf("Expected /balance, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "Balances retrieved",
			"data": []map[string]interface{}{
				{"currency": "NGN", "balance": 1000},
			},
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.CheckBalance(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 balance, got %d", len(resp.Data))
	}
}

func TestResendOTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transfer/resend_otp" {
			t.Errorf("Expected /transfer/resend_otp, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "OTP sent",
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.ResendOTP(context.Background(), &ResendOTPRequest{TransferCode: "TRF_xyz", Reason: "resend"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
}

func TestEnableOTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/transfer/enable_otp" {
			t.Errorf("Expected /transfer/enable_otp, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  true,
			"message": "OTP enabled",
		})
	}))
	defer ts.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(ts.URL)))
	resp, err := client.EnableOTP(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resp.Status {
		t.Errorf("Expected status true, got %v", resp.Status)
	}
}
