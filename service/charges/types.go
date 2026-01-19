package charges

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// CreateChargeRequest represents the request to create a charge
type CreateChargeRequest struct {
	Email             string                 `json:"email"`
	Amount            string                 `json:"amount"`
	AuthorizationCode string                 `json:"authorization_code,omitempty"`
	Pin               string                 `json:"pin,omitempty"`
	Reference         string                 `json:"reference,omitempty"`
	Birthday          string                 `json:"birthday,omitempty"`
	DeviceID          string                 `json:"device_id,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	Bank              *BankSource            `json:"bank,omitempty"`
	MobileMoney       *MobileMoneySource     `json:"mobile_money,omitempty"`
	USSD              *USSDSource            `json:"ussd,omitempty"`
	EFT               *EFTSource             `json:"eft,omitempty"`
}

type BankSource struct {
	Code          string `json:"code"`
	AccountNumber string `json:"account_number"`
}

type MobileMoneySource struct {
	Phone    string `json:"phone"`
	Provider string `json:"provider"`
}

type USSDSource struct {
	Code string `json:"code"`
}

type EFTSource struct {
	Provider string `json:"provider"`
}

// CreateChargeResponse represents the response for creating a charge
type CreateChargeResponse struct {
	paystackapi.Response[ChargeData]
}

// ChargeData represents the data returned after creating a charge
type ChargeData struct {
	Reference   string `json:"reference"`
	Status      string `json:"status"`
	DisplayText string `json:"display_text,omitempty"`
	URL         string `json:"url,omitempty"`
	ID          int    `json:"id,omitempty"`
	Message     string `json:"message,omitempty"`
}

// SubmitPINRequest represents the request to submit a PIN
type SubmitPINRequest struct {
	Pin       string `json:"pin"`
	Reference string `json:"reference"`
}

// SubmitOTPRequest represents the request to submit an OTP
type SubmitOTPRequest struct {
	OTP       string `json:"otp"`
	Reference string `json:"reference"`
}

// SubmitPhoneRequest represents the request to submit a phone number
type SubmitPhoneRequest struct {
	Phone     string `json:"phone"`
	Reference string `json:"reference"`
}

// SubmitBirthdayRequest represents the request to submit a birthday
type SubmitBirthdayRequest struct {
	Birthday  string `json:"birthday"` // Format: YYYY-MM-DD
	Reference string `json:"reference"`
}

// SubmitAddressRequest represents the request to submit an address
type SubmitAddressRequest struct {
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
	Reference string `json:"reference"`
}

// SubmitResponse represents the generic response for submit actions
type SubmitResponse struct {
	paystackapi.Response[ChargeData]
}

// CheckPendingChargeResponse represents the response for checking a pending charge
type CheckPendingChargeResponse struct {
	paystackapi.Response[ChargeData]
}
