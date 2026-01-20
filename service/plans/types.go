package plans

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// Plan represents a payment plan
type Plan struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Amount        int           `json:"amount"`
	Interval      string        `json:"interval"`
	Integration   int           `json:"integration"`
	Domain        string        `json:"domain"`
	PlanCode      string        `json:"plan_code"`
	SendInvoices  bool          `json:"send_invoices"`
	SendSMS       bool          `json:"send_sms"`
	HostedPage    bool          `json:"hosted_page"`
	Currency      string        `json:"currency"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	Subscriptions []interface{} `json:"subscriptions"`
	Pages         []interface{} `json:"pages"`
}

// CreatePlanRequest represents the request to create a plan
type CreatePlanRequest struct {
	Name         string `json:"name"`
	Amount       int    `json:"amount"`
	Interval     string `json:"interval"`
	Description  string `json:"description,omitempty"`
	SendInvoices bool   `json:"send_invoices,omitempty"`
	SendSMS      bool   `json:"send_sms,omitempty"`
	Currency     string `json:"currency,omitempty"`
	InvoiceLimit int    `json:"invoice_limit,omitempty"`
}

// PlanResponse represents the response for a single plan
type PlanResponse struct {
	paystackapi.Response[Plan]
}

// ListPlansResponse represents the response for listing plans
type ListPlansResponse struct {
	paystackapi.Response[[]Plan]
}

// UpdatePlanRequest represents the request to update a plan
type UpdatePlanRequest struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Amount       int    `json:"amount,omitempty"`
	Interval     string `json:"interval,omitempty"`
	SendInvoices bool   `json:"send_invoices,omitempty"`
	SendSMS      bool   `json:"send_sms,omitempty"`
	Currency     string `json:"currency,omitempty"`
	InvoiceLimit int    `json:"invoice_limit,omitempty"`
}
