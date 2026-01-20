package transfers

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

// Initiate initiates a new transfer.
func (c *Client) Initiate(ctx context.Context, req *InitiateRequest) (*InitiateResponse, error) {
	resp := &InitiateResponse{}
	err := c.backend.Call(ctx, "POST", "/transfer", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Finalize completes a transfer.
func (c *Client) Finalize(ctx context.Context, req *FinalizeRequest) (*FinalizeResponse, error) {
	resp := &FinalizeResponse{}
	err := c.backend.Call(ctx, "POST", "/transfer/finalize_transfer", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List retrieves a list of transfers.
func (c *Client) List(ctx context.Context, params *ListTransferParams) (*ListTransferResponse, error) {
	path := "/transfer"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}

	resp := &ListTransferResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch retrieves a transfer by ID or code.
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchResponse, error) {
	resp := &FetchResponse{}
	err := c.backend.Call(ctx, "GET", "/transfer/"+idOrCode, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Verify verifies a transfer by reference.
func (c *Client) Verify(ctx context.Context, reference string) (*VerifyResponse, error) {
	resp := &VerifyResponse{}
	err := c.backend.Call(ctx, "GET", "/transfer/verify/"+reference, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
