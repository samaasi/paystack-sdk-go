package settlements

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// Settlement represents a settlement
type Settlement struct {
	ID              int    `json:"id"`
	Domain          string `json:"domain"`
	Status          string `json:"status"`
	Currency        string `json:"currency"`
	Integration     int    `json:"integration"`
	TotalAmount     int    `json:"total_amount"`
	EffectiveAmount int    `json:"effective_amount"`
	TotalFees       int    `json:"total_fees"`
	TotalProcessed  int    `json:"total_processed"`
	Deductions      int    `json:"deductions"`
	SettlementDate  string `json:"settlement_date"`
	SettledBy       string `json:"settled_by"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// SettlementTransaction represents a transaction in a settlement
type SettlementTransaction struct {
	ID        int    `json:"id"`
	Domain    string `json:"domain"`
	Amount    int    `json:"amount"`
	Currency  string `json:"currency"`
	Reference string `json:"reference"`
	Status    string `json:"status"`
	SplitCode string `json:"split_code"`
	OrderCode string `json:"order_code"`
	PaidAt    string `json:"paid_at"`
	Channel   string `json:"channel"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ListSettlementsResponse represents the response for listing settlements
type ListSettlementsResponse struct {
	paystackapi.Response[[]Settlement]
}

// ListSettlementTransactionsResponse represents the response for listing settlement transactions
type ListSettlementTransactionsResponse struct {
	paystackapi.Response[[]SettlementTransaction]
}
