package paystacksdkgo

import (
	"github.com/samaasi/paystack-sdk-go/internal/backend"
	applepay "github.com/samaasi/paystack-sdk-go/service/apple-pay"
	bulkcharges "github.com/samaasi/paystack-sdk-go/service/bulk-charges"
	"github.com/samaasi/paystack-sdk-go/service/charges"
	"github.com/samaasi/paystack-sdk-go/service/customers"
	"github.com/samaasi/paystack-sdk-go/service/disputes"
	"github.com/samaasi/paystack-sdk-go/service/integration"
	paymentPages "github.com/samaasi/paystack-sdk-go/service/payment-pages"
	paymentRequests "github.com/samaasi/paystack-sdk-go/service/payment-requests"
	"github.com/samaasi/paystack-sdk-go/service/plans"
	"github.com/samaasi/paystack-sdk-go/service/products"
	"github.com/samaasi/paystack-sdk-go/service/refunds"
	"github.com/samaasi/paystack-sdk-go/service/settlements"
	"github.com/samaasi/paystack-sdk-go/service/splits"
	"github.com/samaasi/paystack-sdk-go/service/status"
	"github.com/samaasi/paystack-sdk-go/service/subaccounts"
	"github.com/samaasi/paystack-sdk-go/service/subscriptions"
	"github.com/samaasi/paystack-sdk-go/service/terminal"
	"github.com/samaasi/paystack-sdk-go/service/transactions"
	"github.com/samaasi/paystack-sdk-go/service/transfers"
	"github.com/samaasi/paystack-sdk-go/service/verification"
	virtualAccounts "github.com/samaasi/paystack-sdk-go/service/virtual-accounts"
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
	c.PaymentPages = paymentPages.NewClient(backendClient)
	c.PaymentRequests = paymentRequests.NewClient(backendClient)
	c.Plans = plans.NewClient(backendClient)
	c.Products = products.NewClient(backendClient)
	c.Refunds = refunds.NewClient(backendClient)
	c.Settlements = settlements.NewClient(backendClient)
	c.Splits = splits.NewClient(backendClient)
	c.Status = status.NewClient(backendClient)
	c.Terminal = terminal.NewClient(backendClient)
	c.Subaccounts = subaccounts.NewClient(backendClient)
	c.Subscriptions = subscriptions.NewClient(backendClient)
	c.Verification = verification.NewClient(backendClient)
	c.VirtualAccounts = virtualAccounts.NewClient(backendClient)

	return c
}
