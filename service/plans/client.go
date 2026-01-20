package plans

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Plans service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Plans client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Create creates a new plan
func (c *Client) Create(ctx context.Context, req *CreatePlanRequest) (*PlanResponse, error) {
	resp := &PlanResponse{}
	err := c.backend.Call(ctx, "POST", "/plan", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists plans
func (c *Client) List(ctx context.Context) (*ListPlansResponse, error) {
	resp := &ListPlansResponse{}
	err := c.backend.Call(ctx, "GET", "/plan", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a plan by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*PlanResponse, error) {
	path := fmt.Sprintf("/plan/%s", idOrCode)
	resp := &PlanResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Update updates a plan
func (c *Client) Update(ctx context.Context, idOrCode string, req *UpdatePlanRequest) (*PlanResponse, error) {
	path := fmt.Sprintf("/plan/%s", idOrCode)
	resp := &PlanResponse{}
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
