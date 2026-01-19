package bulkcharges

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

// Initiate initiates a bulk charge
func (c *Client) Initiate(ctx context.Context, req InitiateBulkChargeRequest) (*InitiateBulkChargeResponse, error) {
	resp := &InitiateBulkChargeResponse{}
	err := c.backend.Call(ctx, "POST", "/bulkcharge", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists bulk charge batches
func (c *Client) List(ctx context.Context, params *ListBulkChargesParams) (*ListBulkChargesResponse, error) {
	path := "/bulkcharge"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}

	resp := &ListBulkChargesResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch retrieves a bulk charge batch by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchBulkChargeResponse, error) {
	resp := &FetchBulkChargeResponse{}
	err := c.backend.Call(ctx, "GET", "/bulkcharge/"+idOrCode, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FetchCharges retrieves charges in a specific batch
func (c *Client) FetchCharges(ctx context.Context, idOrCode string, params *FetchChargesInBatchParams) (*FetchChargesInBatchResponse, error) {
	path := fmt.Sprintf("/bulkcharge/%s/charges", idOrCode)
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}

	resp := &FetchChargesInBatchResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Pause pauses a bulk charge batch
func (c *Client) Pause(ctx context.Context, batchCode string) (*PauseBulkChargeResponse, error) {
	resp := &PauseBulkChargeResponse{}
	err := c.backend.Call(ctx, "GET", fmt.Sprintf("/bulkcharge/pause/%s", batchCode), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Resume resumes a paused bulk charge batch
func (c *Client) Resume(ctx context.Context, batchCode string) (*ResumeBulkChargeResponse, error) {
	resp := &ResumeBulkChargeResponse{}
	err := c.backend.Call(ctx, "GET", fmt.Sprintf("/bulkcharge/resume/%s", batchCode), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
