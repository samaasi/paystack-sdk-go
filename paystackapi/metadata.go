package paystackapi

// Meta represents the pagination metadata returned by Paystack API
type Meta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

// Metadata represents the custom metadata object accepted and returned by Paystack API endpoints.
type Metadata map[string]interface{}
