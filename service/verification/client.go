package verification

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

// ResolveAccount resolves an account number
func (c *Client) ResolveAccount(ctx context.Context, accountNumber, bankCode string) (*ResolveAccountResponse, error) {
	resp := &ResolveAccountResponse{}
	path := fmt.Sprintf("/bank/resolve?account_number=%s&bank_code=%s", accountNumber, bankCode)
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ValidateAccount validates an account
func (c *Client) ValidateAccount(ctx context.Context, req *ValidateAccountRequest) (*ValidateAccountResponse, error) {
	resp := &ValidateAccountResponse{}
	err := c.backend.Call(ctx, "POST", "/bank/validate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResolveCardBIN resolves a card BIN
func (c *Client) ResolveCardBIN(ctx context.Context, bin string) (*ResolveCardBINResponse, error) {
	resp := &ResolveCardBINResponse{}
	err := c.backend.Call(ctx, "GET", fmt.Sprintf("/decision/bin/%s", bin), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
