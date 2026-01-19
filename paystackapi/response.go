package paystackapi

// Response is a generic wrapper for Paystack API responses
type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Meta    *Meta  `json:"meta,omitempty"`
}
