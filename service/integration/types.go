package integration

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// PaymentSessionTimeoutResponse represents the response for fetching payment session timeout
type PaymentSessionTimeoutResponse struct {
	paystackapi.Response[PaymentSessionTimeoutData]
}

// PaymentSessionTimeoutData represents the data for payment session timeout
type PaymentSessionTimeoutData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}

// UpdatePaymentSessionTimeoutRequest represents the request to update payment session timeout
type UpdatePaymentSessionTimeoutRequest struct {
	Timeout int `json:"timeout"`
}

// UpdatePaymentSessionTimeoutResponse represents the response for updating payment session timeout
type UpdatePaymentSessionTimeoutResponse struct {
	paystackapi.Response[PaymentSessionTimeoutData]
}
