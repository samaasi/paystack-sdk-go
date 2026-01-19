package paystacksdkgo

import (
	"github.com/samaasi/paystack-sdk-go/internal/backend"
	"github.com/samaasi/paystack-sdk-go/service/transaction"
)

// Client is the main entry point for the Paystack SDK.
type Client struct {
	// Transaction service for handling transaction-related operations
	Transaction *transaction.Client

	// internal backend client
	backend *backend.Client
}
