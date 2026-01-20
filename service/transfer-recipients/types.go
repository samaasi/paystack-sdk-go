package transferRecipients

import "github.com/samaasi/paystack-sdk-go/paystackapi"

type Recipient struct {
	ID            int    `json:"id"`
	Integration   int    `json:"integration"`
	Domain        string `json:"domain"`
	Name          string `json:"name"`
	RecipientCode string `json:"recipient_code"`
	Type          string `json:"type"`
	Details       struct {
		AccountNumber string `json:"account_number"`
		AccountName   string `json:"account_name"`
		BankCode      string `json:"bank_code"`
		BankName      string `json:"bank_name"`
	} `json:"details"`
	Currency  string `json:"currency"`
	Active    bool   `json:"active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateRequest struct {
	Type          string                 `json:"type"`
	Name          string                 `json:"name"`
	AccountNumber string                 `json:"account_number"`
	BankCode      string                 `json:"bank_code"`
	Currency      string                 `json:"currency,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

type CreateResponse struct {
	paystackapi.Response[Recipient]
}

type ListResponse struct {
	paystackapi.Response[[]Recipient]
}

type FetchResponse struct {
	paystackapi.Response[Recipient]
}

type UpdateRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UpdateResponse struct {
	paystackapi.Response[Recipient]
}

type BatchCreateRequest struct {
	Batch []CreateRequest `json:"batch"`
}

type BatchCreateResponse struct {
	paystackapi.Response[[]Recipient]
}
