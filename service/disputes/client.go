package disputes

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/v2/internal/backend"
)

// Client is the client for the Disputes service

// Service represents the interface for disputes operations.
type Service interface {
	List(ctx context.Context, params *ListDisputesParams) (*DisputeListResponse, error)
	Fetch(ctx context.Context, id string) (*DisputeResponse, error)
	ListTransactionDisputes(ctx context.Context, transactionID string) (*DisputeListResponse, error)
	Update(ctx context.Context, id string, req *UpdateDisputeRequest) (*DisputeResponse, error)
	AddEvidence(ctx context.Context, id string, req *AddEvidenceRequest) (*DisputeResponse, error)
	GetUploadURL(ctx context.Context, id string, fileName string) (*UploadURLResponse, error)
	Resolve(ctx context.Context, id string, req *ResolveDisputeRequest) (*DisputeResponse, error)
	Export(ctx context.Context, params *ListDisputesParams) (*ExportDisputesResponse, error)
}

type Client struct {
	backend *backend.Client
}

// NewClient creates a new Disputes client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// List lists disputes
func (c *Client) List(ctx context.Context, params *ListDisputesParams) (*DisputeListResponse, error) {
	path := "/dispute"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}

	resp := &DisputeListResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a dispute
func (c *Client) Fetch(ctx context.Context, id string) (*DisputeResponse, error) {
	resp := &DisputeResponse{}
	path := fmt.Sprintf("/dispute/%s", id)
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListTransactionDisputes lists disputes for a transaction
func (c *Client) ListTransactionDisputes(ctx context.Context, transactionID string) (*DisputeListResponse, error) {
	resp := &DisputeListResponse{}
	path := fmt.Sprintf("/dispute/transaction/%s", transactionID)
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Update updates a dispute
func (c *Client) Update(ctx context.Context, id string, req *UpdateDisputeRequest) (*DisputeResponse, error) {
	resp := &DisputeResponse{}
	path := fmt.Sprintf("/dispute/%s", id)
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AddEvidence adds evidence to a dispute
func (c *Client) AddEvidence(ctx context.Context, id string, req *AddEvidenceRequest) (*DisputeResponse, error) {
	resp := &DisputeResponse{}
	path := fmt.Sprintf("/dispute/%s/evidence", id)
	err := c.backend.Call(ctx, "POST", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetUploadURL gets a URL to upload evidence
func (c *Client) GetUploadURL(ctx context.Context, id string, fileName string) (*UploadURLResponse, error) {
	resp := &UploadURLResponse{}
	path := fmt.Sprintf("/dispute/%s/upload_url?upload_filename=%s", id, fileName)
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Resolve resolves a dispute
func (c *Client) Resolve(ctx context.Context, id string, req *ResolveDisputeRequest) (*DisputeResponse, error) {
	resp := &DisputeResponse{}
	path := fmt.Sprintf("/dispute/%s/resolve", id)
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Export exports disputes
func (c *Client) Export(ctx context.Context, params *ListDisputesParams) (*ExportDisputesResponse, error) {
	path := "/dispute/export"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}

	resp := &ExportDisputesResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
