package webhook

import "encoding/json"

// EventType constants representing the various Paystack webhook events
const (
	EventChargeDisputeCreate           = "charge.dispute.create"
	EventChargeDisputeRemind           = "charge.dispute.remind"
	EventChargeDisputeResolve          = "charge.dispute.resolve"
	EventChargeSuccess                 = "charge.success"
	EventCustomerIdentificationFailed  = "customeridentification.failed"
	EventCustomerIdentificationSuccess = "customeridentification.success"
	EventInvoiceCreate                 = "invoice.create"
	EventInvoicePaymentFailed          = "invoice.payment_failed"
	EventInvoiceUpdate                 = "invoice.update"
	EventPaymentRequestPending         = "paymentrequest.pending"
	EventPaymentRequestSuccess         = "paymentrequest.success"
	EventRefundPending                 = "refund.pending"
	EventRefundProcessing              = "refund.processing"
	EventRefundProcessed               = "refund.processed"
	EventRefundFailed                  = "refund.failed"
	EventSubscriptionCreate            = "subscription.create"
	EventSubscriptionDisable           = "subscription.disable"
	EventSubscriptionEnable            = "subscription.enable"
	EventSubscriptionNotRenew          = "subscription.not_renew"
	EventSubscriptionExpiringCards     = "subscription.expiring_cards"
	EventTransferSuccess               = "transfer.success"
	EventTransferFailed                = "transfer.failed"
	EventTransferReversed              = "transfer.reversed"
	EventDedicatedAccountAssignSuccess = "dedicatedaccount.assign.success"
	EventDedicatedAccountAssignFailed  = "dedicatedaccount.assign.failed"
)

// Event represents a Paystack webhook event
type Event struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

// ParseEvent parses the JSON body into a generic Event struct.
// After parsing, you can inspect the Event field to determine the event type
// and then use the UnmarshalData helper to parse the Data field.
func ParseEvent(body []byte) (*Event, error) {
	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// UnmarshalData unmarshals the event data into the provided target struct.
//
// Example:
//
//	event, _ := webhook.ParseEvent(body)
//	if event.Event == webhook.EventChargeSuccess {
//	    var data webhook.ChargeSuccessEvent
//	    if err := event.UnmarshalData(&data); err != nil {
//	        // handle error
//	    }
//	}
func (e *Event) UnmarshalData(v interface{}) error {
	return json.Unmarshal(e.Data, v)
}

// ChargeSuccessEvent represents the data for a charge.success event
type ChargeSuccessEvent struct {
	ID              int             `json:"id"`
	Domain          string          `json:"domain"`
	Status          string          `json:"status"`
	Reference       string          `json:"reference"`
	Amount          int             `json:"amount"`
	Message         string          `json:"message"`
	GatewayResponse string          `json:"gateway_response"`
	PaidAt          string          `json:"paid_at"`
	CreatedAt       string          `json:"created_at"`
	Channel         string          `json:"channel"`
	Currency        string          `json:"currency"`
	IPAddress       string          `json:"ip_address"`
	Metadata        json.RawMessage `json:"metadata,omitempty"`
	Log             interface{}     `json:"log,omitempty"`
	Fees            int             `json:"fees,omitempty"`
	Customer        Customer        `json:"customer"`
	Authorization   Authorization   `json:"authorization"`
	Plan            interface{}     `json:"plan,omitempty"`
}

type Customer struct {
	ID           int         `json:"id"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Email        string      `json:"email"`
	CustomerCode string      `json:"customer_code"`
	Phone        string      `json:"phone"`
	Metadata     interface{} `json:"metadata"`
	RiskAction   string      `json:"risk_action"`
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
}

// TransferSuccessEvent represents the data for a transfer.success event
type TransferSuccessEvent struct {
	Amount        int         `json:"amount"`
	Currency      string      `json:"currency"`
	Domain        string      `json:"domain"`
	Failures      interface{} `json:"failures"`
	ID            int         `json:"id"`
	Integration   Integration `json:"integration"`
	Reason        string      `json:"reason"`
	Reference     string      `json:"reference"`
	Source        string      `json:"source"`
	SourceDetails interface{} `json:"source_details"`
	Status        string      `json:"status"`
	TitanCode     interface{} `json:"titan_code"`
	TransferCode  string      `json:"transfer_code"`
	TransferredAt string      `json:"transferred_at"`
	Recipient     Recipient   `json:"recipient"`
	Session       interface{} `json:"session"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}

type Integration struct {
	ID           int    `json:"id"`
	IsLive       bool   `json:"is_live"`
	BusinessName string `json:"business_name"`
}

type Recipient struct {
	Active        bool        `json:"active"`
	Currency      string      `json:"currency"`
	Description   string      `json:"description"`
	Domain        string      `json:"domain"`
	Email         interface{} `json:"email"`
	ID            int         `json:"id"`
	Integration   int         `json:"integration"`
	Metadata      interface{} `json:"metadata"`
	Name          string      `json:"name"`
	RecipientCode string      `json:"recipient_code"`
	Type          string      `json:"type"`
	IsDeleted     bool        `json:"is_deleted"`
	Details       BankDetails `json:"details"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}

type BankDetails struct {
	AccountNumber string      `json:"account_number"`
	AccountName   interface{} `json:"account_name"`
	BankCode      string      `json:"bank_code"`
	BankName      string      `json:"bank_name"`
}

// TransferFailedEvent represents the data for a transfer.failed event
type TransferFailedEvent = TransferSuccessEvent

// TransferReversedEvent represents the data for a transfer.reversed event
type TransferReversedEvent = TransferSuccessEvent

// DisputeEvent represents the data for charge.dispute.* events
type DisputeEvent struct {
	ID           int                `json:"id"`
	RefundAmount int                `json:"refund_amount"`
	Currency     string             `json:"currency"`
	Status       string             `json:"status"`
	Resolution   string             `json:"resolution"`
	Reference    string             `json:"reference"`
	Transaction  ChargeSuccessEvent `json:"transaction"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	DisputeCode  string             `json:"dispute_code"`
	Reason       string             `json:"reason"`
}

// InvoiceEvent represents the data for invoice.* events
type InvoiceEvent struct {
	ID            int           `json:"id"`
	Domain        string        `json:"domain"`
	InvoiceCode   string        `json:"invoice_code"`
	Amount        int           `json:"amount"`
	PeriodStart   string        `json:"period_start"`
	PeriodEnd     string        `json:"period_end"`
	Status        string        `json:"status"`
	Paid          bool          `json:"paid"`
	PaidAt        string        `json:"paid_at"`
	Description   string        `json:"description"`
	Authorization Authorization `json:"authorization"`
	Customer      Customer      `json:"customer"`
	CreatedAt     string        `json:"created_at"`
}

// SubscriptionEvent represents the data for subscription.* events
type SubscriptionEvent struct {
	ID               int           `json:"id"`
	Domain           string        `json:"domain"`
	Amount           int           `json:"amount"`
	PeriodStart      string        `json:"period_start"`
	PeriodEnd        string        `json:"period_end"`
	Status           string        `json:"status"`
	SubscriptionCode string        `json:"subscription_code"`
	EmailToken       string        `json:"email_token"`
	EasyTransact     interface{}   `json:"easy_transact"`
	CronExpression   string        `json:"cron_expression"`
	NextPaymentDate  string        `json:"next_payment_date"`
	OpenInvoice      interface{}   `json:"open_invoice"`
	CreatedAt        string        `json:"created_at"`
	Plan             interface{}   `json:"plan"`
	Authorization    Authorization `json:"authorization"`
	Customer         Customer      `json:"customer"`
}

// PaymentRequestEvent represents the data for paymentrequest.* events
type PaymentRequestEvent struct {
	ID               int         `json:"id"`
	Domain           string      `json:"domain"`
	Amount           int         `json:"amount"`
	Currency         string      `json:"currency"`
	DueDate          string      `json:"due_date"`
	HasInvoice       bool        `json:"has_invoice"`
	InvoiceNumber    string      `json:"invoice_number"`
	Description      string      `json:"description"`
	PDFUrl           string      `json:"pdf_url"`
	LineItems        interface{} `json:"line_items"`
	Tax              interface{} `json:"tax"`
	RequestCode      string      `json:"request_code"`
	Status           string      `json:"status"`
	Paid             bool        `json:"paid"`
	PaidAt           string      `json:"paid_at"`
	Metadata         interface{} `json:"metadata"`
	Notifications    interface{} `json:"notifications"`
	OfflineReference string      `json:"offline_reference"`
	Customer         Customer    `json:"customer"`
	CreatedAt        string      `json:"created_at"`
}

// RefundEvent represents the data for refund.* events
type RefundEvent struct {
	ID                   int      `json:"id"`
	RefundedBy           string   `json:"refunded_by"`
	RefundedAt           string   `json:"refunded_at"`
	Amount               int      `json:"amount"`
	Currency             string   `json:"currency"`
	TransactionReference string   `json:"transaction_reference"`
	Status               string   `json:"status"`
	DeductedAmount       int      `json:"deducted_amount"`
	FullyDeducted        bool     `json:"fully_deducted"`
	RefundReference      string   `json:"refund_reference"`
	Customer             Customer `json:"customer"`
}

// DedicatedAccountEvent represents the data for dedicatedaccount.assign.* events
type DedicatedAccountEvent struct {
	Bank          BankDetails `json:"bank"`
	AccountName   string      `json:"account_name"`
	AccountNumber string      `json:"account_number"`
	Assigned      bool        `json:"assigned"`
	Currency      string      `json:"currency"`
	Active        bool        `json:"active"`
	ID            int         `json:"id"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	Customer      Customer    `json:"customer"`
}

// CustomerIdentificationEvent represents the data for customeridentification.* events
type CustomerIdentificationEvent struct {
	CustomerCode   string      `json:"customer_code"`
	Email          string      `json:"email"`
	Identification interface{} `json:"identification"`
	Reason         string      `json:"reason"`
}
