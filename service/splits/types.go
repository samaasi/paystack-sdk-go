package splits

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// Split represents a transaction split
type Split struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	Type             string       `json:"type"`
	Currency         string       `json:"currency"`
	Integration      int          `json:"integration"`
	Domain           string       `json:"domain"`
	SplitCode        string       `json:"split_code"`
	Active           bool         `json:"active"`
	BearerType       string       `json:"bearer_type"`
	BearerSubaccount int          `json:"bearer_subaccount"`
	CreatedAt        string       `json:"created_at"`
	UpdatedAt        string       `json:"updated_at"`
	Subaccounts      []Subaccount `json:"subaccounts"`
	TotalSubaccounts int          `json:"total_subaccounts"`
}

// Subaccount represents a subaccount in a split
type Subaccount struct {
	Subaccount struct {
		ID                  int         `json:"id"`
		SubaccountCode      string      `json:"subaccount_code"`
		BusinessName        string      `json:"business_name"`
		Description         string      `json:"description"`
		PrimaryContactName  string      `json:"primary_contact_name"`
		PrimaryContactEmail string      `json:"primary_contact_email"`
		PrimaryContactPhone string      `json:"primary_contact_phone"`
		Metadata            interface{} `json:"metadata"`
		PercentageCharge    float64     `json:"percentage_charge"`
		SettlementBank      string      `json:"settlement_bank"`
		AccountNumber       string      `json:"account_number"`
	} `json:"subaccount"`
	Share int `json:"share"`
}

// CreateSplitRequest represents the request to create a split
type CreateSplitRequest struct {
	Name             string        `json:"name"`
	Type             string        `json:"type"`
	Currency         string        `json:"currency"`
	Subaccounts      []interface{} `json:"subaccounts"`
	BearerType       string        `json:"bearer_type"`
	BearerSubaccount string        `json:"bearer_subaccount"`
}

// SplitResponse represents the response for a single split
type SplitResponse struct {
	paystackapi.Response[Split]
}

// ListSplitsResponse represents the response for listing splits
type ListSplitsResponse struct {
	paystackapi.Response[[]Split]
}

// UpdateSplitRequest represents the request to update a split
type UpdateSplitRequest struct {
	Name             string `json:"name"`
	Active           bool   `json:"active"`
	BearerType       string `json:"bearer_type,omitempty"`
	BearerSubaccount string `json:"bearer_subaccount,omitempty"`
}

// SubaccountRequest represents the request to add/remove subaccount
type SubaccountRequest struct {
	Subaccount string `json:"subaccount"`
	Share      int    `json:"share,omitempty"`
}
