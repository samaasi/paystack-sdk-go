package paystacksdkgo

import (
	"github.com/samaasi/paystack-sdk-go/internal/backend"
	"github.com/samaasi/paystack-sdk-go/service/transaction"
	"github.com/samaasi/paystack-sdk-go/service/transfer"
)

// Client is the main entry point for the Paystack SDK.
type Client struct {
	// Transaction service for handling transaction-related operations
	Transaction *transaction.Client

	// Transfer service for handling transfer-related operations
	Transfer *transfer.Client

	// internal backend client
	backend *backend.Client
}
