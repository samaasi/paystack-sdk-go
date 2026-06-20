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

// Client is the main entry point for the Paystack SDK.
type Client struct {
	// Transaction service for handling transaction-related operations
	Transactions transactions.Service

	// Transfer service for handling transfer-related operations
	Transfers transfers.Service

	// ApplePay service for handling Apple Pay-related operations
	ApplePay applepay.Service

	// BulkCharges service for handling bulk charge operations
	BulkCharges bulkcharges.Service

	// Charges service for handling charge operations
	Charges charges.Service

	// PaymentPages service for handling payment page operations
	PaymentPages paymentPages.Service

	// PaymentRequests service for handling payment request operations
	PaymentRequests paymentRequests.Service

	// Customers service for handling customer operations
	Customers customers.Service

	// Disputes service for handling dispute operations
	Disputes disputes.Service

	// Integration service for handling integration operations
	Integration integration.Service

	// Plans service for handling plans
	Plans plans.Service

	// Products service for handling products
	Products products.Service

	// Refunds service for handling refunds
	Refunds refunds.Service

	// Settlements service for handling settlements
	Settlements settlements.Service

	// Splits service for handling splits
	Splits splits.Service

	// Status service for handling status
	Status status.Service

	// Subaccounts service for handling subaccounts
	Subaccounts subaccounts.Service

	// Subscriptions service for handling subscriptions
	Subscriptions subscriptions.Service

	// Terminal service for handling terminal
	Terminal terminal.Service

	// Verification service for handling verification
	Verification verification.Service

	// VirtualAccounts service for handling virtual accounts
	VirtualAccounts virtualAccounts.Service

	// TransferControl service for handling transfer control
	TransferControl transferControl.Service

	// TransferRecipients service for handling transfer recipients
	TransferRecipients transferRecipients.Service

	// Misc service for handling miscellaneous operations
	Misc misc.Service

	// internal backend client
	backend *backend.Client
}
