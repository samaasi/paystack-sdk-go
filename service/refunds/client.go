package refunds

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/v2/internal/backend"
)

// Service represents the interface for refunds operations.
type Service interface {
	Create(ctx context.Context, req *CreateRefundRequest) (*RefundResponse, error)
	List(ctx context.Context, params *ListRefundsParams) (*ListRefundsResponse, error)
	Fetch(ctx context.Context, id string) (*RefundResponse, error)
}

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

// List lists refunds with optional filters
func (c *Client) List(ctx context.Context, params *ListRefundsParams) (*ListRefundsResponse, error) {
	path := "/refund"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}
	resp := &ListRefundsResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
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
