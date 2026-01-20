package transferRecipients

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

type Client struct {
	backend *backend.Client
}

func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	resp := &CreateResponse{}
	err := c.backend.Call(ctx, "POST", "/transferrecipient", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) List(ctx context.Context, perPage, page int) (*ListResponse, error) {
	path := fmt.Sprintf("/transferrecipient?perPage=%d&page=%d", perPage, page)
	resp := &ListResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchResponse, error) {
	path := fmt.Sprintf("/transferrecipient/%s", idOrCode)
	resp := &FetchResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Update(ctx context.Context, idOrCode string, req *UpdateRequest) (*UpdateResponse, error) {
	path := fmt.Sprintf("/transferrecipient/%s", idOrCode)
	resp := &UpdateResponse{}
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Delete(ctx context.Context, idOrCode string) (*paystackapi.Response[string], error) {
	path := fmt.Sprintf("/transferrecipient/%s", idOrCode)
	resp := &paystackapi.Response[string]{}
	err := c.backend.Call(ctx, "DELETE", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) BulkCreate(ctx context.Context, req *BatchCreateRequest) (*BatchCreateResponse, error) {
	resp := &BatchCreateResponse{}
	err := c.backend.Call(ctx, "POST", "/transferrecipient/bulk", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
