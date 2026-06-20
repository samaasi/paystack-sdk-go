// Example: initialize a transaction and redirect the user to the checkout URL.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	paystack "github.com/samaasi/paystack-sdk-go/v2"
	"github.com/samaasi/paystack-sdk-go/v2/service/transactions"
)

func main() {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable not set")
	}

	client := paystack.NewClient(secretKey)

	req := &transactions.InitializeRequest{
		Email:  "customer@example.com",
		Amount: "50000", // amount in kobo (NGN 500.00)
	}

	resp, err := client.Transactions.Initialize(context.Background(), req)
	if err != nil {
		log.Fatalf("Initialize failed: %v", err)
	}

	fmt.Println("Authorization URL:", resp.Data.AuthorizationURL)
	fmt.Println("Reference:", resp.Data.Reference)
	// Redirect the user to resp.Data.AuthorizationURL to complete payment.
}
