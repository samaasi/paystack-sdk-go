package transferControl

import "github.com/samaasi/paystack-sdk-go/paystackapi"

type Balance struct {
	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
}

type CheckBalanceResponse struct {
	paystackapi.Response[[]Balance]
}

// LedgerEntry represents a single entry in the balance ledger
type LedgerEntry struct {
	Integration      int    `json:"integration"`
	Domain           string `json:"domain"`
	Balance          int64  `json:"balance"`
	Currency         string `json:"currency"`
	Difference       int64  `json:"difference"`
	Reason           string `json:"reason"`
	Model            string `json:"model"`
	ModelResponsible int    `json:"model_responsible"`
	TransferredAt    string `json:"transferred_at"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

// LedgerParams represents the query parameters for fetching the balance ledger
type LedgerParams struct {
	PerPage *int `query:"perPage,omitempty"`
	Page    *int `query:"page,omitempty"`
}

// LedgerResponse represents the response for the balance ledger
type LedgerResponse struct {
	paystackapi.Response[[]LedgerEntry]
}

type ResendOTPRequest struct {
	TransferCode string `json:"transfer_code"`
	Reason       string `json:"reason"`
}

type ResendOTPResponse struct {
	paystackapi.Response[string]
}

type DisableOTPResponse struct {
	paystackapi.Response[string]
}

type FinalizeDisableOTPRequest struct {
	OTP string `json:"otp"`
}

type FinalizeDisableOTPResponse struct {
	paystackapi.Response[string]
}

type EnableOTPResponse struct {
	paystackapi.Response[string]
}
