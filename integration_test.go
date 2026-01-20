package paystacksdkgo_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	paystack "github.com/samaasi/paystack-sdk-go"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
	"github.com/samaasi/paystack-sdk-go/service/misc"
)

func TestConfiguration(t *testing.T) {
	// Test Timeout
	t.Run("Timeout", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
		}))
		defer ts.Close()

		client := paystack.NewClient("key",
			paystack.WithBaseURL(ts.URL),
			paystack.WithTimeout(100*time.Millisecond),
			paystack.WithMaxRetries(0),
		)

		_, err := client.Misc.ListCountries(context.Background())
		if err == nil {
			t.Error("Expected timeout error, got nil")
		}
	})

	// Test MaxRetries
	t.Run("MaxRetries", func(t *testing.T) {
		attempts := 0
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attempts++
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		client := paystack.NewClient("key",
			paystack.WithBaseURL(ts.URL),
			paystack.WithMaxRetries(2),
		)

		client.Misc.ListCountries(context.Background())

		// Initial request + 2 retries = 3 attempts
		if attempts != 3 {
			t.Errorf("Expected 3 attempts, got %d", attempts)
		}
	})
}

func TestIdempotency(t *testing.T) {
	expectedKey := "unique-key-123"

	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val := r.Header.Get("Idempotency-Key"); val != expectedKey {
			t.Errorf("Expected Idempotency-Key '%s', got '%s'", expectedKey, val)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"status": true, "message": "Success"}`)
	}))
	defer ts.Close()

	// Initialize client
	client := paystack.NewClient("secret_key", paystack.WithBaseURL(ts.URL))

	// Create context with idempotency key
	ctx := paystackapi.WithIdempotencyKey(context.Background(), expectedKey)

	// Call any service method (e.g., ListBanks from Misc)
	_, err := client.Misc.ListCountries(ctx)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestMiscServiceIntegration(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": true,
			"message": "Banks retrieved",
			"data": [
				{
					"name": "Abbey Mortgage Bank",
					"slug": "abbey-mortgage-bank",
					"code": "801",
					"longcode": "",
					"gateway": null,
					"pay_with_bank": false,
					"active": true,
					"country": "Nigeria",
					"currency": "NGN",
					"type": "nuban",
					"is_deleted": false,
					"createdAt": "2020-11-24T10:20:43.000Z",
					"updatedAt": "2020-11-24T10:20:43.000Z"
				}
			]
		}`)
	}))
	defer ts.Close()

	// Initialize client with mock server URL
	client := paystack.NewClient("secret_key", paystack.WithBaseURL(ts.URL))

	// Call Misc service
	resp, err := client.Misc.ListBanks(context.Background(), &misc.ListBanksParams{
		Country: "Nigeria",
		PerPage: 10,
		Page:    1,
	})
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if len(resp.Data) == 0 {
		t.Errorf("Expected banks, got none")
	}
	if resp.Data[0].Name != "Abbey Mortgage Bank" {
		t.Errorf("Expected bank name 'Abbey Mortgage Bank', got '%s'", resp.Data[0].Name)
	}
}

func TestAPIErrorStatusCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"status":false, "message":"Rate limit exceeded"}`))
	}))
	defer ts.Close()

	client := paystack.NewClient("key", paystack.WithBaseURL(ts.URL))
	_, err := client.Misc.ListCountries(context.Background())
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	apiErr, ok := err.(*paystackapi.APIError)
	if !ok {
		t.Fatalf("Expected APIError, got %T", err)
	}

	if apiErr.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status code %d, got %d", http.StatusTooManyRequests, apiErr.StatusCode)
	}
}
