// Example: iterate over all transactions using the built-in Iterator.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	paystack "github.com/samaasi/paystack-sdk-go"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
	"github.com/samaasi/paystack-sdk-go/service/transactions"
)

func main() {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable not set")
	}

	client := paystack.NewClient(secretKey)
	ctx := context.Background()

	it := paystackapi.NewIterator(func(ctx context.Context, page, perPage int) (paystackapi.Response[[]transactions.VerifyData], error) {
		resp, err := client.Transactions.List(ctx, &transactions.ListTransactionParams{})
		if err != nil {
			return paystackapi.Response[[]transactions.VerifyData]{}, err
		}
		return paystackapi.Response[[]transactions.VerifyData]{
			Data: resp.Data,
			Meta: resp.Meta,
		}, nil
	})

	count := 0
	for it.Next(ctx) {
		tx := it.Value()
		fmt.Printf("  [%d] ref=%s status=%s amount=%d\n", tx.ID, tx.Reference, tx.Status, tx.Amount)
		count++
	}
	if err := it.Err(); err != nil {
		log.Fatalf("iterator error: %v", err)
	}
	fmt.Printf("Total transactions fetched: %d\n", count)
}
