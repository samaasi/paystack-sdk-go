package virtualAccounts

import "github.com/samaasi/paystack-sdk-go/paystackapi"

// CreateVirtualAccountRequest represents the request to create a dedicated virtual account
type CreateVirtualAccountRequest struct {
	Customer      string `json:"customer"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

// VirtualAccount represents a dedicated virtual account
type VirtualAccount struct {
	ID            int                    `json:"id"`
	AccountName   string                 `json:"account_name"`
	AccountNumber string                 `json:"account_number"`
	Assigned      bool                   `json:"assigned"`
	Currency      string                 `json:"currency"`
	Metadata      map[string]interface{} `json:"metadata"`
	Active        bool                   `json:"active"`
	SplitConfig   map[string]interface{} `json:"split_config"`
	Bank          Bank                   `json:"bank"`
	Customer      Customer               `json:"customer"`
	Assignment    Assignment             `json:"assignment"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
}

// Bank represents bank details
type Bank struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}

// Customer represents customer details in virtual account
type Customer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	CustomerCode string `json:"customer_code"`
	Phone        string `json:"phone"`
	RiskAction   string `json:"risk_action"`
}

// Assignment represents assignment details
type Assignment struct {
	AssigneeID   int    `json:"assignee_id"`
	AssigneeType string `json:"assignee_type"`
	AssignedAt   string `json:"assigned_at"`
	Expired      bool   `json:"expired"`
	AccountType  string `json:"account_type"`
}

// VirtualAccountResponse represents the response for a single virtual account
type VirtualAccountResponse struct {
	paystackapi.Response[VirtualAccount]
}

// ListVirtualAccountsRequest represents the query parameters for listing virtual accounts
type ListVirtualAccountsRequest struct {
	Active       *bool  `query:"active"`
	Currency     string `query:"currency"`
	ProviderSlug string `query:"provider_slug"`
	BankID       string `query:"bank_id"`
	Customer     string `query:"customer"`
}

// ListVirtualAccountsResponse represents the response for listing virtual accounts
type ListVirtualAccountsResponse struct {
	paystackapi.Response[[]VirtualAccount]
}

// SplitTransactionRequest represents the request to split a transaction
type SplitTransactionRequest struct {
	Customer      string `json:"customer"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
}

// RemoveSplitRequest represents the request to remove a split
type RemoveSplitRequest struct {
	AccountNumber string `json:"account_number"`
}
