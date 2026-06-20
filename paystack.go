package paystacksdkgo

import (
	"github.com/samaasi/paystack-sdk-go/v2/internal/backend"
	applepay "github.com/samaasi/paystack-sdk-go/v2/service/apple-pay"
	bulkcharges "github.com/samaasi/paystack-sdk-go/v2/service/bulk-charges"
	"github.com/samaasi/paystack-sdk-go/v2/service/charges"
	"github.com/samaasi/paystack-sdk-go/v2/service/customers"
	"github.com/samaasi/paystack-sdk-go/v2/service/disputes"
	"github.com/samaasi/paystack-sdk-go/v2/service/integration"
	"github.com/samaasi/paystack-sdk-go/v2/service/misc"
	paymentPages "github.com/samaasi/paystack-sdk-go/v2/service/payment-pages"
	paymentRequests "github.com/samaasi/paystack-sdk-go/v2/service/payment-requests"
	"github.com/samaasi/paystack-sdk-go/v2/service/plans"
	"github.com/samaasi/paystack-sdk-go/v2/service/products"
	"github.com/samaasi/paystack-sdk-go/v2/service/refunds"
	"github.com/samaasi/paystack-sdk-go/v2/service/settlements"
	"github.com/samaasi/paystack-sdk-go/v2/service/splits"
	"github.com/samaasi/paystack-sdk-go/v2/service/status"
	"github.com/samaasi/paystack-sdk-go/v2/service/subaccounts"
	"github.com/samaasi/paystack-sdk-go/v2/service/subscriptions"
	"github.com/samaasi/paystack-sdk-go/v2/service/terminal"
	"github.com/samaasi/paystack-sdk-go/v2/service/transactions"
	transferControl "github.com/samaasi/paystack-sdk-go/v2/service/transfer-control"
	transferRecipients "github.com/samaasi/paystack-sdk-go/v2/service/transfer-recipients"
	"github.com/samaasi/paystack-sdk-go/v2/service/transfers"
	"github.com/samaasi/paystack-sdk-go/v2/service/verification"
	virtualAccounts "github.com/samaasi/paystack-sdk-go/v2/service/virtual-accounts"
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
	c.TransferControl = transferControl.NewClient(backendClient)
	c.Misc = misc.NewClient(backendClient)
	c.TransferRecipients = transferRecipients.NewClient(backendClient)

	return c
}
