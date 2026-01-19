package customers

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// CreateCustomerRequest represents the request to create a customer
type CreateCustomerRequest struct {
	Email     string                 `json:"email"`
	FirstName string                 `json:"first_name,omitempty"`
	LastName  string                 `json:"last_name,omitempty"`
	Phone     string                 `json:"phone,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// UpdateCustomerRequest represents the request to update a customer
type UpdateCustomerRequest struct {
	FirstName string                 `json:"first_name,omitempty"`
	LastName  string                 `json:"last_name,omitempty"`
	Phone     string                 `json:"phone,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// ValidateCustomerRequest represents the request to validate a customer
type ValidateCustomerRequest struct {
	Country       string `json:"country"`
	Type          string `json:"type"`
	Value         string `json:"value"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
	Bvn           string `json:"bvn,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
}

// SetRiskActionRequest represents the request to whitelist or blacklist a customer
type SetRiskActionRequest struct {
	Customer   string `json:"customer"`
	RiskAction string `json:"risk_action,omitempty"` // "default", "allow", "deny"
}

// DeactivateAuthorizationRequest represents the request to deactivate an authorization
type DeactivateAuthorizationRequest struct {
	AuthorizationCode string `json:"authorization_code"`
}

// CustomerResponse represents the generic response for customer operations
type CustomerResponse struct {
	paystackapi.Response[CustomerData]
}

// CustomerListResponse represents the response for listing customers
type CustomerListResponse struct {
	paystackapi.Response[[]CustomerData]
}

// ValidateCustomerResponse represents the response for validating a customer
type ValidateCustomerResponse struct {
	paystackapi.Response[bool] // Data is usually boolean or empty on success depending on endpoint
}

// CustomerData represents the customer object
type CustomerData struct {
	ID              int                    `json:"id"`
	Integration     int                    `json:"integration,omitempty"`
	FirstName       string                 `json:"first_name"`
	LastName        string                 `json:"last_name"`
	Email           string                 `json:"email"`
	Phone           string                 `json:"phone"`
	Metadata        map[string]interface{} `json:"metadata"`
	Domain          string                 `json:"domain"`
	CustomerCode    string                 `json:"customer_code"`
	RiskAction      string                 `json:"risk_action"`
	Identified      bool                   `json:"identified"`
	Identifications interface{}            `json:"identifications"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
}

// ListCustomersParams represents query parameters for listing customers
type ListCustomersParams struct {
	PerPage int    `json:"perPage,omitempty"`
	Page    int    `json:"page,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
}
