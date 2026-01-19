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
	"github.com/samaasi/paystack-sdk-go/service/transactions"
	"github.com/samaasi/paystack-sdk-go/service/transfers"
)

// Client is the main entry point for the Paystack SDK.
type Client struct {
	// Transaction service for handling transaction-related operations
	Transactions *transactions.Client

	// Transfer service for handling transfer-related operations
	Transfers *transfers.Client

	// ApplePay service for handling Apple Pay-related operations
	ApplePay *applepay.Client

	// BulkCharges service for handling bulk charge operations
	BulkCharges *bulkcharges.Client

	// Charges service for handling charge operations
	Charges *charges.Client

	// PaymentPages service for handling payment page operations
	PaymentPages *paymentPages.Client

	// PaymentRequests service for handling payment request operations
	PaymentRequests *paymentRequests.Client

	// Customers service for handling customer operations
	Customers *customers.Client

	// Disputes service for handling dispute operations
	Disputes *disputes.Client

	// Integration service for handling integration operations
	Integration *integration.Client

	// internal backend client
	backend *backend.Client
}
