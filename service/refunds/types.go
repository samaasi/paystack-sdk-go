package refunds

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// Refund represents a refund
type Refund struct {
	ID             int    `json:"id"`
	Integration    int    `json:"integration"`
	Domain         string `json:"domain"`
	Transaction    int    `json:"transaction"`
	Dispute        int    `json:"dispute"`
	Amount         int    `json:"amount"`
	Currency       string `json:"currency"`
	DeductedAmount int    `json:"deducted_amount"`
	RefundedBy     string `json:"refunded_by"`
	RefundedAt     string `json:"refunded_at"`
	FullyDeducted  bool   `json:"fully_deducted"`
	Status         string `json:"status"`
	CustomerNote   string `json:"customer_note"`
	MerchantNote   string `json:"merchant_note"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// CreateRefundRequest represents the request to create a refund
type CreateRefundRequest struct {
	Transaction  string `json:"transaction"`
	Amount       int    `json:"amount,omitempty"`
	Currency     string `json:"currency,omitempty"`
	CustomerNote string `json:"customer_note,omitempty"`
	MerchantNote string `json:"merchant_note,omitempty"`
}

// RefundResponse represents the response for a single refund
type RefundResponse struct {
	paystackapi.Response[Refund]
}

// ListRefundsParams represents the query parameters for listing refunds
type ListRefundsParams struct {
	PerPage     *int    `query:"perPage,omitempty"`
	Page        *int    `query:"page,omitempty"`
	From        *string `query:"from,omitempty"`
	To          *string `query:"to,omitempty"`
	Transaction *string `query:"transaction,omitempty"`
	Currency    *string `query:"currency,omitempty"`
}

// ListRefundsResponse represents the response for listing refunds
type ListRefundsResponse struct {
	paystackapi.Response[[]Refund]
}
