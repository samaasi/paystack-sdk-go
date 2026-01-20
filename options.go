package paystacksdkgo

import (
	"net/http"
	"time"

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

// WithMaxRetries sets the maximum number of retries for failed requests.
// Default is 3. Set to 0 to disable retries.
func WithMaxRetries(retries int) ClientOption {
	return backend.WithMaxRetries(retries)
}

// WithTimeout sets the timeout for HTTP requests.
// Default is 30 seconds.
func WithTimeout(timeout time.Duration) ClientOption {
	return backend.WithTimeout(timeout)
}

