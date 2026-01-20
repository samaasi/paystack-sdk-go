package subaccounts

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

func (c *Client) Create(ctx context.Context, req *CreateSubaccountRequest) (*SubaccountResponse, error) {
	resp := &SubaccountResponse{}
	err := c.backend.Call(ctx, "POST", "/subaccount", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) List(ctx context.Context, perPage, page int) (*ListSubaccountsResponse, error) {
	path := fmt.Sprintf("/subaccount?perPage=%d&page=%d", perPage, page)
	resp := &ListSubaccountsResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*SubaccountResponse, error) {
	path := fmt.Sprintf("/subaccount/%s", idOrCode)
	resp := &SubaccountResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Update(ctx context.Context, idOrCode string, req *UpdateSubaccountRequest) (*SubaccountResponse, error) {
	path := fmt.Sprintf("/subaccount/%s", idOrCode)
	resp := &SubaccountResponse{}
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
