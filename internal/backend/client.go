package backend

import (
	"context"
	"net/http"
	"time"

	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

const DefaultBaseURL = "https://api.paystack.co"

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
	backoff    *Backoff
}

func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		baseURL:    DefaultBaseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		backoff:    DefaultBackoff(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

type ClientOption func(*Client)

func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// Call makes an HTTP request and decodes the response into v.
func (c *Client) Call(ctx context.Context, method, path string, body, v interface{}) error {
	url := c.baseURL + path

	op := func() error {
		req, err := NewRequest(method, url, c.apiKey, body, nil)
		if err != nil {
			return err
		}
		req = req.WithContext(ctx)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return &paystackapi.RequestError{Err: err}
		}

		return Decode(resp, v)
	}

	return c.backoff.Retry(ctx, op)
}
