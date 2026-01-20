package paystacksdkgo_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	paystack "github.com/samaasi/paystack-sdk-go"
)

func TestIntegration_Misc(t *testing.T) {
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
	resp, err := client.Misc.ListBanks(context.Background(), "Nigeria", 10, 1)
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
