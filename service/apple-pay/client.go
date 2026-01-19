package applepay

import (
	"context"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

type Client struct {
	backend *backend.Client
}

func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// RegisterDomain registers a domain for Apple Pay
func (c *Client) RegisterDomain(ctx context.Context, req *RegisterDomainRequest) (*RegisterDomainResponse, error) {
	resp := &RegisterDomainResponse{}
	err := c.backend.Call(ctx, "POST", "/apple-pay/domain", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListDomains lists all registered domains for Apple Pay
func (c *Client) ListDomains(ctx context.Context) (*ListDomainsResponse, error) {
	resp := &ListDomainsResponse{}
	err := c.backend.Call(ctx, "GET", "/apple-pay/domain", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UnregisterDomain unregisters a domain for Apple Pay
func (c *Client) UnregisterDomain(ctx context.Context, req *UnregisterDomainRequest) (*UnregisterDomainResponse, error) {
	resp := &UnregisterDomainResponse{}
	err := c.backend.Call(ctx, "DELETE", "/apple-pay/domain", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
