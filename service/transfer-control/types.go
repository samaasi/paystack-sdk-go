package transferControl

import "github.com/samaasi/paystack-sdk-go/paystackapi"

type Balance struct {
	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
}

type CheckBalanceResponse struct {
	paystackapi.Response[[]Balance]
}

type ResendOTPRequest struct {
	TransferCode string `json:"transfer_code"`
	Reason       string `json:"reason"`
}

type ResendOTPResponse struct {
	paystackapi.Response[string]
}

type DisableOTPResponse struct {
	paystackapi.Response[string]
}

type FinalizeDisableOTPRequest struct {
	OTP string `json:"otp"`
}

type FinalizeDisableOTPResponse struct {
	paystackapi.Response[string]
}

type EnableOTPResponse struct {
	paystackapi.Response[string]
}
