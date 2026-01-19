package transaction

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

type Client struct {
	backend *backend.Client
}

func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Initialize initiates a transaction.
func (c *Client) Initialize(ctx context.Context, req *InitializeRequest) (*InitializeResponse, error) {
	resp := &InitializeResponse{}
	err := c.backend.Call(ctx, "POST", "/transaction/initialize", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Verify confirms the status of a transaction.
func (c *Client) Verify(ctx context.Context, reference string) (*VerifyResponse, error) {
	resp := &VerifyResponse{}
	err := c.backend.Call(ctx, "GET", "/transaction/verify/"+reference, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List retrieves a list of transactions.
func (c *Client) List(ctx context.Context, params *ListTransactionParams) (*ListTransactionResponse, error) {
	path := "/transaction"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}

	resp := &ListTransactionResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
