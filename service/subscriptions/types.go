package subscriptions

import (
	"encoding/json"

	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

type Subscription struct {
	ID               int             `json:"id"`
	SubscriptionCode string          `json:"subscription_code"`
	Amount           float64         `json:"amount"`
	CronExpression   string          `json:"cron_expression"`
	NextPaymentDate  string          `json:"next_payment_date"`
	OpenInvoice      string          `json:"open_invoice"`
	CreatedAt        string          `json:"created_at"`
	Plan             json.RawMessage `json:"plan"`
	Customer         json.RawMessage `json:"customer"`
	Authorization    json.RawMessage `json:"authorization"`
	Status           string          `json:"status"`
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
