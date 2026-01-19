package paystacksdkgo

import (
	"github.com/samaasi/paystack-sdk-go/internal/backend"
	applepay "github.com/samaasi/paystack-sdk-go/service/apple-pay"
	bulkcharges "github.com/samaasi/paystack-sdk-go/service/bulk-charges"
	"github.com/samaasi/paystack-sdk-go/service/charges"
	"github.com/samaasi/paystack-sdk-go/service/transaction"
	"github.com/samaasi/paystack-sdk-go/service/transfer"
)

// Client is the main entry point for the Paystack SDK.
type Client struct {
	// Transaction service for handling transaction-related operations
	Transaction *transaction.Client

	// Transfer service for handling transfer-related operations
	Transfer *transfer.Client

	// ApplePay service for handling Apple Pay-related operations
	ApplePay *applepay.Client

	// BulkCharges service for handling bulk charge operations
	BulkCharges *bulkcharges.Client

	// Charges service for handling charge operations
	Charges *charges.Client

	// internal backend client
	backend *backend.Client
}
