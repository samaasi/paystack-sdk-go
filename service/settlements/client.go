package settlements

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Settlements service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Settlements client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// List lists settlements
func (c *Client) List(ctx context.Context) (*ListSettlementsResponse, error) {
	resp := &ListSettlementsResponse{}
	err := c.backend.Call(ctx, "GET", "/settlement", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FetchTransactions fetches transactions for a settlement
func (c *Client) FetchTransactions(ctx context.Context, id string) (*ListSettlementTransactionsResponse, error) {
	path := fmt.Sprintf("/settlement/%s/transactions", id)
	resp := &ListSettlementTransactionsResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
