package paymentpages

import (
	"encoding/json"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// Page represents a payment page
type Page struct {
	ID           int             `json:"id"`
	Integration  int             `json:"integration"`
	Domain       string          `json:"domain"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Amount       int             `json:"amount"`
	Currency     string          `json:"currency"`
	Slug         string          `json:"slug"`
	Active       bool            `json:"active"`
	RedirectURL  string          `json:"redirect_url"`
	CustomFields []interface{}   `json:"custom_fields"`
	Type         string          `json:"type"`
	Metadata     json.RawMessage `json:"metadata"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
}

// CreatePageRequest represents the request to create a payment page
type CreatePageRequest struct {
	Name         string        `json:"name"`
	Description  string        `json:"description,omitempty"`
	Amount       int           `json:"amount,omitempty"`
	Slug         string        `json:"slug,omitempty"`
	RedirectURL  string        `json:"redirect_url,omitempty"`
	CustomFields []interface{} `json:"custom_fields,omitempty"`
	Metadata     interface{}   `json:"metadata,omitempty"`
}

// PageResponse represents the response for a single page
type PageResponse struct {
	paystackapi.Response[Page]
}

// ListPagesResponse represents the response for listing pages
type ListPagesResponse struct {
	paystackapi.Response[[]Page]
}

// UpdatePageRequest represents the request to update a payment page
type UpdatePageRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Active      *bool  `json:"active,omitempty"`
}

// CheckSlugResponse represents the response for checking slug availability
type CheckSlugResponse struct {
	paystackapi.Response[bool] // Data is usually true/false or object with status
}

// AddProductsRequest represents the request to add products to a page
type AddProductsRequest struct {
	Products []int `json:"product"`
}

// AddProductsResponse represents the response for adding products
type AddProductsResponse struct {
	paystackapi.Response[Page]
}
