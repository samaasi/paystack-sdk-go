package integration

import (
	"context"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Integration service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Integration client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// FetchPaymentSessionTimeout fetches the payment session timeout
func (c *Client) FetchPaymentSessionTimeout(ctx context.Context) (*PaymentSessionTimeoutResponse, error) {
	resp := &PaymentSessionTimeoutResponse{}
	err := c.backend.Call(ctx, "GET", "/integration/payment_session_timeout", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdatePaymentSessionTimeout updates the payment session timeout
func (c *Client) UpdatePaymentSessionTimeout(ctx context.Context, timeout int) (*UpdatePaymentSessionTimeoutResponse, error) {
	req := &UpdatePaymentSessionTimeoutRequest{Timeout: timeout}
	resp := &UpdatePaymentSessionTimeoutResponse{}
	err := c.backend.Call(ctx, "PUT", "/integration/payment_session_timeout", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
