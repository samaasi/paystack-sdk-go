package products

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Products service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Products client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Create creates a new product
func (c *Client) Create(ctx context.Context, req *CreateProductRequest) (*ProductResponse, error) {
	resp := &ProductResponse{}
	err := c.backend.Call(ctx, "POST", "/product", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists products
func (c *Client) List(ctx context.Context) (*ListProductsResponse, error) {
	resp := &ListProductsResponse{}
	err := c.backend.Call(ctx, "GET", "/product", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a product by ID
func (c *Client) Fetch(ctx context.Context, id string) (*ProductResponse, error) {
	path := fmt.Sprintf("/product/%s", id)
	resp := &ProductResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Update updates a product
func (c *Client) Update(ctx context.Context, id string, req *UpdateProductRequest) (*ProductResponse, error) {
	path := fmt.Sprintf("/product/%s", id)
	resp := &ProductResponse{}
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
