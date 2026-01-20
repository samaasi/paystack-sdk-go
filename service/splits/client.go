package splits

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Splits service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Splits client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Create creates a new split
func (c *Client) Create(ctx context.Context, req *CreateSplitRequest) (*SplitResponse, error) {
	resp := &SplitResponse{}
	err := c.backend.Call(ctx, "POST", "/split", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists splits
func (c *Client) List(ctx context.Context) (*ListSplitsResponse, error) {
	resp := &ListSplitsResponse{}
	err := c.backend.Call(ctx, "GET", "/split", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a split by ID
func (c *Client) Fetch(ctx context.Context, id string) (*SplitResponse, error) {
	path := fmt.Sprintf("/split/%s", id)
	resp := &SplitResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Update updates a split
func (c *Client) Update(ctx context.Context, id string, req *UpdateSplitRequest) (*SplitResponse, error) {
	path := fmt.Sprintf("/split/%s", id)
	resp := &SplitResponse{}
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AddSubaccount adds a subaccount to a split
func (c *Client) AddSubaccount(ctx context.Context, id string, req *SubaccountRequest) (*SplitResponse, error) {
	path := fmt.Sprintf("/split/%s/subaccount/add", id)
	resp := &SplitResponse{}
	err := c.backend.Call(ctx, "POST", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RemoveSubaccount removes a subaccount from a split
func (c *Client) RemoveSubaccount(ctx context.Context, id string, req *SubaccountRequest) (*SplitResponse, error) {
	path := fmt.Sprintf("/split/%s/subaccount/remove", id)
	resp := &SplitResponse{}
	err := c.backend.Call(ctx, "POST", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
