package transfers

import "github.com/samaasi/paystack-sdk-go/paystackapi"

// InitiateRequest represents the payload for initiating a transfer
type InitiateRequest struct {
	Source    string               `json:"source"`
	Amount    int                  `json:"amount"`
	Recipient string               `json:"recipient"`
	Reason    string               `json:"reason,omitempty"`
	Currency  string               `json:"currency,omitempty"`
	Reference string               `json:"reference,omitempty"`
	Metadata  paystackapi.Metadata `json:"metadata,omitempty"`
}

// InitiateResponse represents the response for transfer initialization
type InitiateResponse struct {
	paystackapi.Response[InitiateData]
}

type InitiateData struct {
	Reference    string `json:"reference"`
	Integration  int    `json:"integration"`
	Domain       string `json:"domain"`
	Amount       int    `json:"amount"`
	Currency     string `json:"currency"`
	Source       string `json:"source"`
	Reason       string `json:"reason"`
	Recipient    int    `json:"recipient"`
	Status       string `json:"status"`
	TransferCode string `json:"transfer_code"`
	ID           int    `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// FinalizeRequest represents the payload for finalizing a transfer
type FinalizeRequest struct {
	TransferCode string `json:"transfer_code"`
	OTP          string `json:"otp"`
}

// FinalizeResponse represents the response for transfer finalization
type FinalizeResponse struct {
	paystackapi.Response[FinalizeData]
}

type FinalizeData struct {
	Domain        string      `json:"domain"`
	Amount        int         `json:"amount"`
	Currency      string      `json:"currency"`
	Reference     string      `json:"reference"`
	Source        string      `json:"source"`
	SourceDetails interface{} `json:"source_details"`
	Reason        string      `json:"reason"`
	Recipient     Recipient   `json:"recipient"`
	Status        string      `json:"status"`
	TransferCode  string      `json:"transfer_code"`
	ID            int         `json:"id"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}

type Recipient struct {
	Domain        string               `json:"domain"`
	Type          string               `json:"type"`
	Currency      string               `json:"currency"`
	Name          string               `json:"name"`
	Details       RecipientDetails     `json:"details"`
	Description   string               `json:"description"`
	Metadata      paystackapi.Metadata `json:"metadata"`
	RecipientCode string               `json:"recipient_code"`
	Active        bool                 `json:"active"`
	ID            int                  `json:"id"`
	CreatedAt     string               `json:"created_at"`
	UpdatedAt     string               `json:"updated_at"`
}

type RecipientDetails struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
}

// ListTransferParams represents query parameters for listing transfers
type ListTransferParams struct {
	PerPage  *int    `query:"perPage,omitempty"`
	Page     *int    `query:"page,omitempty"`
	Customer *string `query:"customer,omitempty"`
	From     *string `query:"from,omitempty"`
	To       *string `query:"to,omitempty"`
}

// ListTransferResponse represents the response for listing transfers
type ListTransferResponse struct {
	paystackapi.Response[[]InitiateData]
}

// BulkTransferRequest represents the payload for initiating multiple transfers at once.
type BulkTransferRequest struct {
	Currency  string           `json:"currency"`
	Transfers []InitiateRequest `json:"transfers"`
}

// BulkTransferResponse represents the response for a bulk transfer request.
type BulkTransferResponse struct {
	paystackapi.Response[[]InitiateData]
}

// FetchResponse represents the response for fetching a transfer
type FetchResponse struct {
	paystackapi.Response[InitiateData]
}

// VerifyResponse represents the response for verifying a transfer
type VerifyResponse struct {
	paystackapi.Response[InitiateData]
}
