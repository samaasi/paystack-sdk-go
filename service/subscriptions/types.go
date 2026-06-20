package subscriptions

import (
	"encoding/json"

	"github.com/samaasi/paystack-sdk-go/v2/paystackapi"
)

type Subscription struct {
	ID               int             `json:"id"`
	SubscriptionCode string          `json:"subscription_code"`
	Amount           float64         `json:"amount"`
	CronExpression   string          `json:"cron_expression"`
	NextPaymentDate  string          `json:"next_payment_date"`
	OpenInvoice      string          `json:"open_invoice"`
	EmailToken       string          `json:"email_token"`
	Quantity         int             `json:"quantity"`
	InvoiceLimit     int             `json:"invoice_limit"`
	SplitCode        string          `json:"split_code"`
	CreatedAt        string          `json:"created_at"`
	Plan             json.RawMessage `json:"plan"`
	Customer         json.RawMessage `json:"customer"`
	Authorization    json.RawMessage `json:"authorization"`
	MostRecentInvoice *Invoice       `json:"most_recent_invoice,omitempty"`
	Status           string          `json:"status"`
}

// Invoice represents a subscription invoice.
type Invoice struct {
	Subscription int    `json:"subscription"`
	Integration  int    `json:"integration"`
	Domain       string `json:"domain"`
	InvoiceCode  string `json:"invoice_code"`
	Customer     int    `json:"customer"`
	Transaction  int    `json:"transaction"`
	Amount       int    `json:"amount"`
	PeriodStart  string `json:"period_start"`
	PeriodEnd    string `json:"period_end"`
	Status       string `json:"status"`
	Paid         bool   `json:"paid"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CreateSubscriptionRequest struct {
	Customer      string `json:"customer"`
	Plan          string `json:"plan"`
	Authorization string `json:"authorization,omitempty"`
	StartDate     string `json:"start_date,omitempty"`
}

type SubscriptionResponse struct {
	paystackapi.Response[Subscription]
}

type ListSubscriptionsResponse struct {
	paystackapi.Response[[]Subscription]
}

type EnableDisableSubscriptionRequest struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}

// GenerateLinkData holds the subscription management link.
type GenerateLinkData struct {
	Link string `json:"link"`
}

// GenerateLinkResponse represents the response for generating a subscription management link.
type GenerateLinkResponse struct {
	paystackapi.Response[GenerateLinkData]
}
