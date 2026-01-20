package products

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// Product represents a product
type Product struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	ProductCode     string        `json:"product_code"`
	Price           int           `json:"price"`
	Currency        string        `json:"currency"`
	Quantity        int           `json:"quantity"`
	QuantitySold    int           `json:"quantity_sold"`
	Active          bool          `json:"active"`
	Domain          string        `json:"domain"`
	Type            string        `json:"type"`
	InStock         bool          `json:"in_stock"`
	Unlimited       bool          `json:"unlimited"`
	Metadata        interface{}   `json:"metadata"`
	Files           []interface{} `json:"files"`
	SuccessMessage  string        `json:"success_message"`
	RedirectURL     string        `json:"redirect_url"`
	SplitCode       string        `json:"split_code"`
	NotificationEmail string      `json:"notification_email"`
	MinimumOrderable int          `json:"minimum_orderable"`
	MaximumOrderable int          `json:"maximum_orderable"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
	Integration     int           `json:"integration"`
	LowStockAlert   bool          `json:"low_stock_alert"`
	StockThreshold  int           `json:"stock_threshold"`
	ExpiresIn       string        `json:"expires_in"`
}

// CreateProductRequest represents the request to create a product
type CreateProductRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Price           int    `json:"price"`
	Currency        string `json:"currency"`
	Unlimited       bool   `json:"unlimited,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
}

// ProductResponse represents the response for a single product
type ProductResponse struct {
	paystackapi.Response[Product]
}

// ListProductsResponse represents the response for listing products
type ListProductsResponse struct {
	paystackapi.Response[[]Product]
}

// UpdateProductRequest represents the request to update a product
type UpdateProductRequest struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	Price           int    `json:"price,omitempty"`
	Currency        string `json:"currency,omitempty"`
	Unlimited       bool   `json:"unlimited,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
}
