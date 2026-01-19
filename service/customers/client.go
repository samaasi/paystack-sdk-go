package customers

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Customers service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Customers client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Create creates a new customer
func (c *Client) Create(ctx context.Context, req *CreateCustomerRequest) (*CustomerResponse, error) {
	resp := &CustomerResponse{}
	err := c.backend.Call(ctx, "POST", "/customer", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists customers
func (c *Client) List(ctx context.Context, params *ListCustomersParams) (*CustomerListResponse, error) {
	resp := &CustomerListResponse{}
	err := c.backend.Call(ctx, "GET", "/customer", params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a customer by email or code
func (c *Client) Fetch(ctx context.Context, emailOrCode string) (*CustomerResponse, error) {
	resp := &CustomerResponse{}
	path := fmt.Sprintf("/customer/%s", emailOrCode)
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Update updates a customer's details
func (c *Client) Update(ctx context.Context, code string, req *UpdateCustomerRequest) (*CustomerResponse, error) {
	resp := &CustomerResponse{}
	path := fmt.Sprintf("/customer/%s", code)
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Validate validates a customer's identity
func (c *Client) Validate(ctx context.Context, code string, req *ValidateCustomerRequest) (*ValidateCustomerResponse, error) {
	resp := &ValidateCustomerResponse{}
	path := fmt.Sprintf("/customer/%s/identification", code)
	err := c.backend.Call(ctx, "POST", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Whitelist whitelists a customer
func (c *Client) Whitelist(ctx context.Context, customerCode string) (*CustomerResponse, error) {
	req := &SetRiskActionRequest{
		Customer:   customerCode,
		RiskAction: "allow",
	}
	resp := &CustomerResponse{}
	err := c.backend.Call(ctx, "POST", "/customer/set_risk_action", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Blacklist blacklists a customer
func (c *Client) Blacklist(ctx context.Context, customerCode string) (*CustomerResponse, error) {
	req := &SetRiskActionRequest{
		Customer:   customerCode,
		RiskAction: "deny",
	}
	resp := &CustomerResponse{}
	err := c.backend.Call(ctx, "POST", "/customer/set_risk_action", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeactivateAuthorization deactivates an authorization
func (c *Client) DeactivateAuthorization(ctx context.Context, authorizationCode string) (*CustomerResponse, error) {
	req := &DeactivateAuthorizationRequest{
		AuthorizationCode: authorizationCode,
	}
	resp := &CustomerResponse{}
	err := c.backend.Call(ctx, "POST", "/customer/deactivate_authorization", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
