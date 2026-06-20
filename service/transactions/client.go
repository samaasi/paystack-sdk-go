package transactions

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/v2/internal/backend"
)

// Service represents the interface for transactions operations.
type Service interface {
	Initialize(ctx context.Context, req *InitializeRequest) (*InitializeResponse, error)
	Verify(ctx context.Context, reference string) (*VerifyResponse, error)
	List(ctx context.Context, params *ListTransactionParams) (*ListTransactionResponse, error)
	Fetch(ctx context.Context, id int) (*FetchResponse, error)
	ChargeAuthorization(ctx context.Context, req *ChargeAuthorizationRequest) (*ChargeAuthorizationResponse, error)
	GetTimeline(ctx context.Context, idOrReference string) (*TimelineResponse, error)
	Totals(ctx context.Context, params *TotalsParams) (*TotalsResponse, error)
	Export(ctx context.Context, params *ExportParams) (*ExportResponse, error)
	PartialDebit(ctx context.Context, req *PartialDebitRequest) (*PartialDebitResponse, error)
}

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

// Fetch retrieves a single transaction by its ID.
func (c *Client) Fetch(ctx context.Context, id int) (*FetchResponse, error) {
	resp := &FetchResponse{}
	err := c.backend.Call(ctx, "GET", fmt.Sprintf("/transaction/%d", id), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ChargeAuthorization charges a returning customer using a previously saved authorization code.
func (c *Client) ChargeAuthorization(ctx context.Context, req *ChargeAuthorizationRequest) (*ChargeAuthorizationResponse, error) {
	resp := &ChargeAuthorizationResponse{}
	err := c.backend.Call(ctx, "POST", "/transaction/charge_authorization", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTimeline fetches the timeline of a transaction using its ID or reference.
func (c *Client) GetTimeline(ctx context.Context, idOrReference string) (*TimelineResponse, error) {
	resp := &TimelineResponse{}
	err := c.backend.Call(ctx, "GET", "/transaction/timeline/"+idOrReference, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Totals returns the total amount received on your account, optionally filtered by date range.
func (c *Client) Totals(ctx context.Context, params *TotalsParams) (*TotalsResponse, error) {
	path := "/transaction/totals"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}
	resp := &TotalsResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Export exports a list of transactions carried out on your integration as a downloadable file.
func (c *Client) Export(ctx context.Context, params *ExportParams) (*ExportResponse, error) {
	path := "/transaction/export"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}
	resp := &ExportResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PartialDebit charges a customer's card less than the authorization amount for a transaction.
func (c *Client) PartialDebit(ctx context.Context, req *PartialDebitRequest) (*PartialDebitResponse, error) {
	resp := &PartialDebitResponse{}
	err := c.backend.Call(ctx, "POST", "/transaction/partial_debit", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
