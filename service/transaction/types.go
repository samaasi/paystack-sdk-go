package transaction

import "github.com/samaasi/paystack-sdk-go/paystackapi"

// InitializeRequest represents the payload for initializing a transaction
type InitializeRequest struct {
	Amount   string `json:"amount"`
	Email    string `json:"email"`
	Currency string `json:"currency,omitempty"`
}

// InitializeResponse represents the response for transaction initialization
type InitializeResponse struct {
	paystackapi.Response[InitializeData]
}

type InitializeData struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

// VerifyResponse represents the response for transaction verification
type VerifyResponse struct {
	paystackapi.Response[VerifyData]
}

type VerifyData struct {
	ID              int    `json:"id"`
	Domain          string `json:"domain"`
	Status          string `json:"status"`
	Reference       string `json:"reference"`
	Amount          int    `json:"amount"`
	Message         string `json:"message"`
	GatewayResponse string `json:"gateway_response"`
	PaidAt          string `json:"paid_at"`
	CreatedAt       string `json:"created_at"`
	Channel         string `json:"channel"`
	Currency        string `json:"currency"`
	IPAddress       string `json:"ip_address"`
}

// ListTransactionParams represents query parameters for listing transactions
type ListTransactionParams struct {
	PerPage int `query:"perPage"`
	Page    int `query:"page"`
}

// ListTransactionResponse represents the response for listing transactions
type ListTransactionResponse struct {
	paystackapi.Response[[]VerifyData]
}
