package transactions

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// InitializeRequest represents the payload for initializing a transaction
type InitializeRequest struct {
	Amount            string                `json:"amount"`
	Email             string                `json:"email"`
	Currency          paystackapi.Currency  `json:"currency,omitempty"`
	Reference         string                `json:"reference,omitempty"`
	CallbackURL       string                `json:"callback_url,omitempty"`
	Plan              string                `json:"plan,omitempty"`
	InvoiceLimit      int                   `json:"invoice_limit,omitempty"`
	Metadata          paystackapi.Metadata  `json:"metadata,omitempty"`
	Channels          []paystackapi.Channel `json:"channels,omitempty"`
	SplitCode         string                `json:"split_code,omitempty"`
	Subaccount        string                `json:"subaccount,omitempty"`
	TransactionCharge int                   `json:"transaction_charge,omitempty"`
	Bearer            paystackapi.Bearer    `json:"bearer,omitempty"`
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
	Status          paystackapi.Status   `json:"status"`
	Reference       string               `json:"reference"`
	Amount          int                  `json:"amount"`
	Message         string               `json:"message"`
	GatewayResponse string               `json:"gateway_response"`
	PaidAt          string               `json:"paid_at"`
	CreatedAt       string               `json:"created_at"`
	Channel         paystackapi.Channel  `json:"channel"`
	Currency        paystackapi.Currency `json:"currency"`
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
	AuthorizationCode string              `json:"authorization_code"`
	Bin               string              `json:"bin"`
	Last4             string              `json:"last4"`
	ExpMonth          string              `json:"exp_month"`
	ExpYear           string              `json:"exp_year"`
	Channel           paystackapi.Channel `json:"channel"`
	CardType          string              `json:"card_type"`
	Bank              string              `json:"bank"`
	CountryCode       string              `json:"country_code"`
	Brand             string              `json:"brand"`
	Reusable          bool                `json:"reusable"`
	Signature         string              `json:"signature"`
	AccountName       string              `json:"account_name"`
}

type Customer struct {
	ID           int                    `json:"id"`
	FirstName    string                 `json:"first_name"`
	LastName     string                 `json:"last_name"`
	Email        string                 `json:"email"`
	CustomerCode string                 `json:"customer_code"`
	Phone        string                 `json:"phone"`
	RiskAction   paystackapi.RiskAction `json:"risk_action"`
}

// ListTransactionParams represents query parameters for listing transactions
type ListTransactionParams struct {
	PerPage *int    `query:"perPage,omitempty"`
	Page    *int    `query:"page,omitempty"`
	Status  *string `query:"status,omitempty"`
	From    *string `query:"from,omitempty"`
	To      *string `query:"to,omitempty"`
}

// ListTransactionResponse represents the response for listing transactions
type ListTransactionResponse struct {
	paystackapi.Response[[]VerifyData]
}

// FetchResponse represents the response for fetching a single transaction by ID.
type FetchResponse struct {
	paystackapi.Response[VerifyData]
}

// ChargeAuthorizationRequest represents the payload for charging a saved authorization.
type ChargeAuthorizationRequest struct {
	Amount            string                `json:"amount"`
	Email             string                `json:"email"`
	AuthorizationCode string                `json:"authorization_code"`
	Reference         string                `json:"reference,omitempty"`
	Currency          paystackapi.Currency  `json:"currency,omitempty"`
	Metadata          paystackapi.Metadata  `json:"metadata,omitempty"`
	Channels          []paystackapi.Channel `json:"channels,omitempty"`
	Subaccount        string                `json:"subaccount,omitempty"`
	TransactionCharge int                   `json:"transaction_charge,omitempty"`
	Bearer            paystackapi.Bearer    `json:"bearer,omitempty"`
	Queue             bool                  `json:"queue,omitempty"`
}

// ChargeAuthorizationResponse represents the response for charging a saved authorization.
type ChargeAuthorizationResponse struct {
	paystackapi.Response[VerifyData]
}

// TimelineData represents the transaction timeline.
type TimelineData struct {
	TimeSpent      int           `json:"time_spent"`
	Attempts       int           `json:"attempts"`
	Authentication string        `json:"authentication"`
	Errors         int           `json:"errors"`
	Success        bool          `json:"success"`
	Mobile         bool          `json:"mobile"`
	Input          []interface{} `json:"input"`
	Channel        string        `json:"channel"`
	History        []History     `json:"history"`
}

// TimelineResponse represents the response for fetching a transaction timeline.
type TimelineResponse struct {
	paystackapi.Response[TimelineData]
}

// TotalsParams represents optional date-range filters for transaction totals.
type TotalsParams struct {
	From string `query:"from,omitempty"`
	To   string `query:"to,omitempty"`
}

// CurrencyTotal represents total amounts grouped by currency.
type CurrencyTotal struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

// TotalsData represents the aggregated transaction totals for the account.
type TotalsData struct {
	TotalTransactions          int             `json:"total_transactions"`
	TotalVolume                int             `json:"total_volume"`
	TotalVolumeByCurrency      []CurrencyTotal `json:"total_volume_by_currency"`
	PendingTransfers           int             `json:"pending_transfers"`
	PendingTransfersByCurrency []CurrencyTotal `json:"pending_transfers_by_currency"`
}

// TotalsResponse represents the response for transaction totals.
type TotalsResponse struct {
	paystackapi.Response[TotalsData]
}

// ExportParams represents optional query parameters for exporting transactions.
type ExportParams struct {
	From        string `query:"from,omitempty"`
	To          string `query:"to,omitempty"`
	Customer    *int   `query:"customer,omitempty"`
	Status      string `query:"status,omitempty"`
	Currency    string `query:"currency,omitempty"`
	Amount      *int   `query:"amount,omitempty"`
	Settled     *bool  `query:"settled,omitempty"`
	Settlement  *int   `query:"settlement,omitempty"`
	PaymentPage *int   `query:"payment_page,omitempty"`
}

// ExportData holds the path to the exported transactions file.
type ExportData struct {
	Path string `json:"path"`
}

// ExportResponse represents the response for exporting transactions.
type ExportResponse struct {
	paystackapi.Response[ExportData]
}

// PartialDebitRequest represents the payload for a partial debit.
type PartialDebitRequest struct {
	AuthorizationCode string               `json:"authorization_code"`
	Currency          paystackapi.Currency `json:"currency"`
	Amount            string               `json:"amount"`
	Email             string               `json:"email"`
	Reference         string               `json:"reference,omitempty"`
	AtLeast           string               `json:"at_least,omitempty"`
}

// PartialDebitResponse represents the response for a partial debit.
type PartialDebitResponse struct {
	paystackapi.Response[VerifyData]
}
