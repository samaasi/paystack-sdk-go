package paystackapi

import "fmt"

// APIError represents an error returned by the Paystack API
type APIError struct {
	Status     bool                `json:"status"`
	Message    string              `json:"message"`
	Code       string              `json:"code,omitempty"`
	Errors     map[string][]string `json:"errors,omitempty"`
	StatusCode int                 `json:"-"`
	RetryAfter int                 `json:"-"`
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("paystack: %s (code: %s, status: %d)", e.Message, e.Code, e.StatusCode)
	}
	return fmt.Sprintf("paystack: %s (status: %d)", e.Message, e.StatusCode)
}

// Is allows checking if an error is a specific status code or code.
func (e *APIError) Is(target error) bool {
	t, ok := target.(*APIError)
	if !ok {
		return false
	}
	if t.StatusCode != 0 && e.StatusCode != t.StatusCode {
		return false
	}
	if t.Code != "" && e.Code != t.Code {
		return false
	}
	return true
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
