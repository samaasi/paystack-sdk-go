package transactions

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// InitializeRequest represents the payload for initializing a transaction
type InitializeRequest struct {
	Amount            string               `json:"amount"`
	Email             string               `json:"email"`
	Currency          string               `json:"currency,omitempty"`
	Reference         string               `json:"reference,omitempty"`
	CallbackURL       string               `json:"callback_url,omitempty"`
	Plan              string               `json:"plan,omitempty"`
	InvoiceLimit      int                  `json:"invoice_limit,omitempty"`
	Metadata          paystackapi.Metadata `json:"metadata,omitempty"`
	Channels          []string             `json:"channels,omitempty"`
	SplitCode         string               `json:"split_code,omitempty"`
	Subaccount        string               `json:"subaccount,omitempty"`
	TransactionCharge int                  `json:"transaction_charge,omitempty"`
	Bearer            string               `json:"bearer,omitempty"`
}

// InitializeResponse represents the response for transaction initialization
type InitializeResponse struct {
	paystackapi.Response[InitializeData]
}

type InitializeData struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

// VerifyResponse represents the response for transaction verification
type VerifyResponse struct {
	paystackapi.Response[VerifyData]
}

type VerifyData struct {
	ID              int                  `json:"id"`
	Domain          string               `json:"domain"`
	Status          string               `json:"status"`
	Reference       string               `json:"reference"`
	Amount          int                  `json:"amount"`
	Message         string               `json:"message"`
	GatewayResponse string               `json:"gateway_response"`
	PaidAt          string               `json:"paid_at"`
	CreatedAt       string               `json:"created_at"`
	Channel         string               `json:"channel"`
	Currency        string               `json:"currency"`
	IPAddress       string               `json:"ip_address"`
	Metadata        paystackapi.Metadata `json:"metadata"`
	Log             *Log                 `json:"log,omitempty"`
	Fees            int                  `json:"fees"`
	FeesBreakdown   interface{}          `json:"fees_breakdown,omitempty"`
	Authorization   Authorization        `json:"authorization"`
	Customer        Customer             `json:"customer"`
	Plan            interface{}          `json:"plan"`
	Split           interface{}          `json:"split"`
	OrderID         interface{}          `json:"order_id"`
	RequestedAmount int                  `json:"requested_amount"`
}

type Log struct {
	StartTime int           `json:"start_time"`
	TimeSpent int           `json:"time_spent"`
	Attempts  int           `json:"attempts"`
	Errors    int           `json:"errors"`
	Success   bool          `json:"success"`
	Mobile    bool          `json:"mobile"`
	Input     []interface{} `json:"input"`
	History   []History     `json:"history"`
}

type History struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

type Authorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	Channel           string `json:"channel"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reusable          bool   `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

type Customer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	CustomerCode string `json:"customer_code"`
	Phone        string `json:"phone"`
	RiskAction   string `json:"risk_action"`
}

// ListTransactionParams represents query parameters for listing transactions
type ListTransactionParams struct {
	PerPage int `query:"perPage"`
	Page    int `query:"page"`
}

// ListTransactionResponse represents the response for listing transactions
type ListTransactionResponse struct {
	paystackapi.Response[[]VerifyData]
}
