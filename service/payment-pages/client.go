package paymentpages

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Payment Pages service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Payment Pages client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Create creates a new payment page
func (c *Client) Create(ctx context.Context, req *CreatePageRequest) (*PageResponse, error) {
	resp := &PageResponse{}
	err := c.backend.Call(ctx, "POST", "/page", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists payment pages
func (c *Client) List(ctx context.Context) (*ListPagesResponse, error) {
	resp := &ListPagesResponse{}
	err := c.backend.Call(ctx, "GET", "/page", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a payment page by ID or slug
func (c *Client) Fetch(ctx context.Context, idOrSlug string) (*PageResponse, error) {
	path := fmt.Sprintf("/page/%s", idOrSlug)
	resp := &PageResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Update updates a payment page
func (c *Client) Update(ctx context.Context, idOrSlug string, req *UpdatePageRequest) (*PageResponse, error) {
	path := fmt.Sprintf("/page/%s", idOrSlug)
	resp := &PageResponse{}
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CheckSlugAvailability checks if a slug is available
func (c *Client) CheckSlugAvailability(ctx context.Context, slug string) (*CheckSlugResponse, error) {
	path := fmt.Sprintf("/page/check_slug_availability/%s", slug)
	resp := &CheckSlugResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AddProducts adds products to a payment page
func (c *Client) AddProducts(ctx context.Context, id int, req *AddProductsRequest) (*AddProductsResponse, error) {
	path := fmt.Sprintf("/page/%d/product", id)
	resp := &AddProductsResponse{}
	err := c.backend.Call(ctx, "POST", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
