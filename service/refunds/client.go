package refunds

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Refunds service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Refunds client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Create creates a new refund
func (c *Client) Create(ctx context.Context, req *CreateRefundRequest) (*RefundResponse, error) {
	resp := &RefundResponse{}
	err := c.backend.Call(ctx, "POST", "/refund", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists refunds
func (c *Client) List(ctx context.Context) (*ListRefundsResponse, error) {
	resp := &ListRefundsResponse{}
	err := c.backend.Call(ctx, "GET", "/refund", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a refund by ID
func (c *Client) Fetch(ctx context.Context, id string) (*RefundResponse, error) {
	path := fmt.Sprintf("/refund/%s", id)
	resp := &RefundResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
