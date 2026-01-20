# Paystack SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/samaasi/paystack-sdk-go.svg)](https://pkg.go.dev/github.com/samaasi/paystack-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/samaasi/paystack-sdk-go)](https://goreportcard.com/report/github.com/samaasi/paystack-sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive and robust Go SDK for the [Paystack API](https://paystack.com/docs/api/). This library provides a clean, idiomatic Go interface for integrating Paystack payments into your applications.

## Features

- **Full Coverage**: Support for Transactions, Customers, Plans, Subscriptions, Transfers, and more.
- **Context Aware**: All methods support `context.Context` for timeouts and cancellation.
- **Idempotency**: Built-in support for Idempotency Keys via context.
- **Robust Error Handling**: Typed errors for better control flow.
- **Modular Design**: Services are isolated for better maintainability.

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
	"github.com/samaasi/paystack-sdk-go/service/transactions"
)

func main() {
	client := paystack.NewClient(os.Getenv("PAYSTACK_SECRET_KEY"))

	req := &transactions.InitializeRequest{
		Email:  "customer@email.com",
		Amount: "500000", // in kobo
	}

	resp, err := client.Transactions.Initialize(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Authorization URL: %s\n", resp.Data.AuthorizationURL)
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
