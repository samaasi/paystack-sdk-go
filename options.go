package paystacksdkgo

import (
	"net/http"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// ClientOption is a function that configures the Client.
type ClientOption = backend.ClientOption

// WithHTTPClient allows you to provide a custom HTTP client.
// This is useful for testing or if you need to configure proxies, timeouts, etc.
func WithHTTPClient(client *http.Client) ClientOption {
	return backend.WithHTTPClient(client)
}

// WithBaseURL allows you to override the default Paystack API base URL.
func WithBaseURL(url string) ClientOption {
	return backend.WithBaseURL(url)
}
