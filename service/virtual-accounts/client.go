package virtualAccounts

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/v2/internal/backend"
)

// Service represents the interface for virtual-accounts operations.
type Service interface {
	Create(ctx context.Context, req *CreateVirtualAccountRequest) (*VirtualAccountResponse, error)
	List(ctx context.Context, req *ListVirtualAccountsRequest) (*ListVirtualAccountsResponse, error)
	Fetch(ctx context.Context, id int) (*VirtualAccountResponse, error)
	Deactivate(ctx context.Context, id int) (*VirtualAccountResponse, error)
	FetchBanks(ctx context.Context) (*FetchBanksResponse, error)
	SplitTransaction(ctx context.Context, req *SplitTransactionRequest) (*VirtualAccountResponse, error)
	RemoveSplit(ctx context.Context, req *RemoveSplitRequest) (*VirtualAccountResponse, error)
}

type Client struct {
	backend *backend.Client
}

func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Create creates a dedicated virtual account
func (c *Client) Create(ctx context.Context, req *CreateVirtualAccountRequest) (*VirtualAccountResponse, error) {
	resp := &VirtualAccountResponse{}
	err := c.backend.Call(ctx, "POST", "/dedicated_account", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists dedicated virtual accounts
func (c *Client) List(ctx context.Context, req *ListVirtualAccountsRequest) (*ListVirtualAccountsResponse, error) {
	resp := &ListVirtualAccountsResponse{}
	queryParams, err := backend.EncodeQueryParams(req)
	if err != nil {
		return nil, err
	}

	path := "/dedicated_account"
	if queryParams != "" {
		path += "?" + queryParams
	}

	err = c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a dedicated virtual account by ID
func (c *Client) Fetch(ctx context.Context, id int) (*VirtualAccountResponse, error) {
	resp := &VirtualAccountResponse{}
	err := c.backend.Call(ctx, "GET", fmt.Sprintf("/dedicated_account/%d", id), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Deactivate deactivates a dedicated virtual account
func (c *Client) Deactivate(ctx context.Context, id int) (*VirtualAccountResponse, error) {
	resp := &VirtualAccountResponse{}
	err := c.backend.Call(ctx, "DELETE", fmt.Sprintf("/dedicated_account/%d", id), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FetchBanks returns the list of banks that support dedicated virtual accounts
func (c *Client) FetchBanks(ctx context.Context) (*FetchBanksResponse, error) {
	resp := &FetchBanksResponse{}
	err := c.backend.Call(ctx, "GET", "/dedicated_account/available_banks", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SplitTransaction splits a dedicated virtual account transaction
func (c *Client) SplitTransaction(ctx context.Context, req *SplitTransactionRequest) (*VirtualAccountResponse, error) {
	resp := &VirtualAccountResponse{}
	err := c.backend.Call(ctx, "POST", "/dedicated_account/split", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RemoveSplit removes a split from a dedicated virtual account
func (c *Client) RemoveSplit(ctx context.Context, req *RemoveSplitRequest) (*VirtualAccountResponse, error) {
	resp := &VirtualAccountResponse{}
	err := c.backend.Call(ctx, "DELETE", "/dedicated_account/split", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
