# Paystack SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/samaasi/paystack-sdk-go.svg)](https://pkg.go.dev/github.com/samaasi/paystack-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/samaasi/paystack-sdk-go)](https://goreportcard.com/report/github.com/samaasi/paystack-sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive and robust Go SDK for the [Paystack API](https://paystack.com/docs/api/). This library provides a clean, idiomatic Go interface for integrating Paystack payments into your applications.

## Features

- **Full API Coverage**: Comprehensive support for all 24 Paystack services (Transactions, Customers, Transfers, etc.).
- **Smart Retries & Rate Limiting**: Built-in exponential backoff that respects Paystack's `429 Retry-After` headers and strictly avoids retrying non-retryable `4xx` errors.
- **Automatic Idempotency**: Automatically generates UUIDv4 `Idempotency-Key` headers for all `POST` and `PUT` requests to prevent duplicate charges on network retries.
- **Generic Pagination Iterators**: Best-in-class `paystackapi.Iterator[T]` for elegantly fetching paginated list endpoints.
- **Strongly Typed & Safe**: Uses strictly typed enums (`paystackapi.Currency`, `paystackapi.Channel`, etc.) rather than raw strings to prevent runtime errors.
- **100% Mockable**: Every service client is exposed as a Go `interface` allowing effortless unit testing via `gomock`.
- **Webhook IP Allowlisting**: Defense-in-depth webhook signature verification and official IP address validation.
- **Zero Dependencies**: Relies solely on the Go standard library for a lightweight footprint and maximum security.

## Installation

```bash
go get github.com/samaasi/paystack-sdk-go
```

## Usage

### Initialization

Initialize the client with your secret key.

```go
package main

import (
	"fmt"
	"os"

	paystack "github.com/samaasi/paystack-sdk-go"
)

func main() {
	apiKey := os.Getenv("PAYSTACK_SECRET_KEY")
	client := paystack.NewClient(apiKey)
	
	// You can also pass options
	// client := paystack.NewClient(apiKey, 
	//     paystack.WithBaseURL("https://api.paystack.co"),
	//     paystack.WithMaxRetries(5),
	//     paystack.WithTimeout(10 * time.Second),
	// )
}
```

### Configuration

You can configure the client with the following options:

- `WithBaseURL(url string)`: Override the default Paystack API base URL.
- `WithMaxRetries(retries int)`: Set the maximum number of retries for failed requests (default: 3).
- `WithTimeout(timeout time.Duration)`: Set the timeout for HTTP requests (default: 30s).
- `WithHTTPClient(client *http.Client)`: Use a custom HTTP client.

### Making Requests

Example: Initializing a transaction.

```go
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
	client := paystack.NewClient(os.Getenv("PAYSTACK_SECRET_KEY"))

	req := &transactions.InitializeRequest{
		Email:    "customer@email.com",
		Amount:   "500000", // in kobo
		Currency: paystackapi.CurrencyNGN,
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

	resp, err := client.Transactions.Initialize(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Authorization URL: %s\n", resp.Data.AuthorizationURL)
}
```

### Pagination with Iterators

The SDK provides a generic Iterator pattern to effortlessly fetch records across multiple pages without manually handling `Meta` cursors or next-page logic.

```go
package main

import (
	"context"
	"fmt"
	"log"

	paystack "github.com/samaasi/paystack-sdk-go"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
	"github.com/samaasi/paystack-sdk-go/service/transactions"
)

func main() {
	client := paystack.NewClient("sk_test_...")
	ctx := context.Background()

	iter := paystackapi.NewIterator(ctx, func(ctx context.Context, page, perPage int) (paystackapi.Response[[]transactions.VerifyData], error) {
		return client.Transactions.List(ctx, &transactions.ListTransactionParams{
			PerPage: perPage, 
			Page:    page,
		})
	})

	for iter.Next() {
		tx := iter.Value()
		fmt.Printf("Transaction ID: %d, Status: %s\n", tx.ID, tx.Status)
	}

	if err := iter.Err(); err != nil {
		log.Fatal("Error during iteration:", err)
	}
}
```

### Webhook Verification

Easily and securely verify incoming webhooks from Paystack using HMAC validation and IP Allowlisting.

```go
package main

import (
	"log"
	"net/http"

	"github.com/samaasi/paystack-sdk-go/webhook"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Defense-in-depth: Verify request originates from Paystack's official IPs
	if !webhook.IsFromPaystackIP(r) {
		http.Error(w, "Unauthorized IP", http.StatusUnauthorized)
		return
	}

	// Parse and verify the payload signature using your secret key
	var event webhook.Event
	if err := webhook.Parse(r, "PAYSTACK_SECRET_KEY", &event); err != nil {
		log.Printf("Webhook validation failed: %v", err)
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	log.Printf("Received event: %s", event.Event)
	w.WriteHeader(http.StatusOK)
}
```

### Idempotency Keys

To prevent duplicate operations, you can pass an Idempotency Key using the context.

```go
import (
	"context"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

func main() {
	// ... init client ...

	ctx := context.Background()
	// Add Idempotency Key to context
	ctx = paystackapi.WithIdempotencyKey(ctx, "unique-transaction-id-123")

	// The key will be sent in the header as 'Idempotency-Key'
	resp, err := client.Transactions.Initialize(ctx, req)
}
```

### Custom Headers

You can also pass arbitrary custom headers via context.

```go
ctx := paystackapi.WithCustomHeader(context.Background(), "X-Custom-Header", "Value")
```

## Supported Services

- Apple Pay
- Bulk Charges
- Charges
- Customers
- Disputes
- Integration
- Miscellaneous (Banks, Countries, etc.)
- Payment Pages
- Payment Requests
- Plans
- Products
- Refunds
- Settlements
- Splits
- Status
- Subaccounts
- Subscriptions
- Terminal
- Transactions
- Transfer Control
- Transfer Recipients
- Transfers
- Verification
- Virtual Accounts

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

To get started:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/amazing-feature`).
3. Commit your changes (`git commit -m 'Add some amazing feature'`).
4. Push to the branch (`git push origin feature/amazing-feature`).
5. Open a Pull Request.

Please ensure you run tests before submitting:

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
