package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// Decode uses json.Decoder to decode the response body into v.
// It also checks for API errors.
func Decode(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var apiErr paystackapi.APIError
		// Try to decode into APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			// Fallback for non-JSON errors or decode failures
			return &paystackapi.APIError{
				Status:  false,
				Message: fmt.Sprintf("HTTP %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode)),
			}
		}
		return &apiErr
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}
	return nil
}
