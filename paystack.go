package paystacksdkgo

import (
	"github.com/samaasi/paystack-sdk-go/internal/backend"
	applepay "github.com/samaasi/paystack-sdk-go/service/apple-pay"
	bulkcharges "github.com/samaasi/paystack-sdk-go/service/bulk-charges"
	"github.com/samaasi/paystack-sdk-go/service/charges"
	"github.com/samaasi/paystack-sdk-go/service/customers"
	"github.com/samaasi/paystack-sdk-go/service/disputes"
	"github.com/samaasi/paystack-sdk-go/service/integration"
	"github.com/samaasi/paystack-sdk-go/service/transactions"
	"github.com/samaasi/paystack-sdk-go/service/transfers"
)

// NewClient creates a new Paystack client with the given secret key.
// You can pass the secret key from your environment variables or any other source.
//
// Example:
//
//	client := paystack.NewClient(os.Getenv("PAYSTACK_SECRET_KEY"))
func NewClient(secretKey string, opts ...ClientOption) *Client {
	// Convert public options to internal backend options
	var backendOpts []backend.ClientOption
	for _, opt := range opts {
		backendOpts = append(backendOpts, opt)
	}

	// Initialize backend client
	backendClient := backend.NewClient(secretKey, backendOpts...)

	// Initialize SDK client
	c := &Client{
		backend: backendClient,
	}

	// Initialize services
	c.Transactions = transactions.NewClient(backendClient)
	c.Transfers = transfers.NewClient(backendClient)
	c.ApplePay = applepay.NewClient(backendClient)
	c.BulkCharges = bulkcharges.NewClient(backendClient)
	c.Charges = charges.NewClient(backendClient)
	c.Customers = customers.NewClient(backendClient)
	c.Disputes = disputes.NewClient(backendClient)
	c.Integration = integration.NewClient(backendClient)

	return c
}
