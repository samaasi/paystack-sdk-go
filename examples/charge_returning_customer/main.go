// Example: charge a returning customer using a saved authorization code.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	paystack "github.com/samaasi/paystack-sdk-go/v2"
	"github.com/samaasi/paystack-sdk-go/v2/paystackapi"
	"github.com/samaasi/paystack-sdk-go/v2/service/transactions"
)

func main() {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable not set")
	}

	client := paystack.NewClient(secretKey)

	req := &transactions.ChargeAuthorizationRequest{
		AuthorizationCode: "AUTH_xxxxxxxxxx",
		Email:             "customer@example.com",
		Amount:            "10000", // NGN 100.00 in kobo
		Currency:          paystackapi.CurrencyNGN,
	}

	resp, err := client.Transactions.ChargeAuthorization(context.Background(), req)
	if err != nil {
		var apiErr *paystackapi.APIError
		if errors.As(err, &apiErr) {
			log.Fatalf("API error %d: %s (code: %s)", apiErr.StatusCode, apiErr.Message, apiErr.Code)
		}
		log.Fatalf("request error: %v", err)
	}

	fmt.Printf("Charge status: %s\n", resp.Data.Status)
	fmt.Printf("Reference:     %s\n", resp.Data.Reference)
}
