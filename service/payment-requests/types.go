package paymentrequests

import (
	"encoding/json"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// PaymentRequest represents a payment request
type PaymentRequest struct {
	ID               int             `json:"id"`
	Domain           string          `json:"domain"`
	Amount           int             `json:"amount"`
	Currency         string          `json:"currency"`
	DueDate          string          `json:"due_date"`
	HasInvoice       bool            `json:"has_invoice"`
	InvoiceNumber    int             `json:"invoice_number"`
	Description      string          `json:"description"`
	PDFURL           string          `json:"pdf_url"`
	LineItems        []interface{}   `json:"line_items"`
	Tax              []interface{}   `json:"tax"`
	RequestCode      string          `json:"request_code"`
	Status           string          `json:"status"`
	Paid             bool            `json:"paid"`
	PaidAt           string          `json:"paid_at"`
	Metadata         json.RawMessage `json:"metadata"`
	Notifications    []interface{}   `json:"notifications"`
	OfflineReference string          `json:"offline_reference"`
	Customer         int             `json:"customer"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at"`
}

// CreatePaymentRequestRequest represents the request to create a payment request
type CreatePaymentRequestRequest struct {
	Customer         string        `json:"customer"`
	Amount           int           `json:"amount"`
	DueDate          string        `json:"due_date,omitempty"`
	Description      string        `json:"description,omitempty"`
	LineItems        []interface{} `json:"line_items,omitempty"`
	Tax              []interface{} `json:"tax,omitempty"`
	Currency         string        `json:"currency,omitempty"`
	SendNotification bool          `json:"send_notification,omitempty"`
	Draft            bool          `json:"draft,omitempty"`
	HasInvoice       bool          `json:"has_invoice,omitempty"`
	InvoiceNumber    int           `json:"invoice_number,omitempty"`
	SplitCode        string        `json:"split_code,omitempty"`
}

// PaymentRequestResponse represents the response for a single payment request
type PaymentRequestResponse struct {
	paystackapi.Response[PaymentRequest]
}

// ListPaymentRequestsResponse represents the response for listing payment requests
type ListPaymentRequestsResponse struct {
	paystackapi.Response[[]PaymentRequest]
}

// UpdatePaymentRequestRequest represents the request to update a payment request
type UpdatePaymentRequestRequest struct {
	Customer         string        `json:"customer,omitempty"`
	Amount           int           `json:"amount,omitempty"`
	DueDate          string        `json:"due_date,omitempty"`
	Description      string        `json:"description,omitempty"`
	LineItems        []interface{} `json:"line_items,omitempty"`
	Tax              []interface{} `json:"tax,omitempty"`
	Currency         string        `json:"currency,omitempty"`
	SendNotification bool          `json:"send_notification,omitempty"`
	Draft            bool          `json:"draft,omitempty"`
}

// PaymentRequestTotalResponse represents the response for payment request total
type PaymentRequestTotalResponse struct {
	paystackapi.Response[struct {
		Total []struct {
			Currency string `json:"currency"`
			Amount   int    `json:"amount"`
		} `json:"total"`
	}]
}
