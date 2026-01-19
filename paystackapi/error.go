package paystackapi

import "fmt"

// APIError represents an error returned by the Paystack API
type APIError struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("paystack: %s (code: %s)", e.Message, e.Code)
	}
	return fmt.Sprintf("paystack: %s", e.Message)
}

// RequestError represents an error that occurred while making the request
type RequestError struct {
	Err error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("paystack request failed: %v", e.Err)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}
