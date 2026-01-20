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
	transferControl "github.com/samaasi/paystack-sdk-go/service/transfer-control"
	transferRecipients "github.com/samaasi/paystack-sdk-go/service/transfer-recipients"
	"github.com/samaasi/paystack-sdk-go/service/transfers"
	"github.com/samaasi/paystack-sdk-go/service/verification"
	virtualAccounts "github.com/samaasi/paystack-sdk-go/service/virtual-accounts"
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

	// Plans service for handling plans
	Plans *plans.Client

	// Products service for handling products
	Products *products.Client

	// Refunds service for handling refunds
	Refunds *refunds.Client

	// Settlements service for handling settlements
	Settlements *settlements.Client

	// Splits service for handling splits
	Splits *splits.Client

	// Status service for handling status
	Status *status.Client

	// Subaccounts service for handling subaccounts
	Subaccounts *subaccounts.Client

	// Subscriptions service for handling subscriptions
	Subscriptions *subscriptions.Client

	// Terminal service for handling terminal
	Terminal *terminal.Client

	// Verification service for handling verification
	Verification *verification.Client

	// VirtualAccounts service for handling virtual accounts
	VirtualAccounts *virtualAccounts.Client

	// TransferControl service for handling transfer control
	TransferControl *transferControl.Client

	// TransferRecipients service for handling transfer recipients
	TransferRecipients *transferRecipients.Client

	// internal backend client
	backend *backend.Client
}
