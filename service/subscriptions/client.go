package subscriptions

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

func (c *Client) Create(ctx context.Context, req *CreateSubscriptionRequest) (*SubscriptionResponse, error) {
	resp := &SubscriptionResponse{}
	err := c.backend.Call(ctx, "POST", "/subscription", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) List(ctx context.Context, perPage, page int) (*ListSubscriptionsResponse, error) {
	path := fmt.Sprintf("/subscription?perPage=%d&page=%d", perPage, page)
	resp := &ListSubscriptionsResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*SubscriptionResponse, error) {
	path := fmt.Sprintf("/subscription/%s", idOrCode)
	resp := &SubscriptionResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Enable(ctx context.Context, req *EnableDisableSubscriptionRequest) (*paystackapi.Response[interface{}], error) {
	resp := &paystackapi.Response[interface{}]{}
	err := c.backend.Call(ctx, "POST", "/subscription/enable", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Disable(ctx context.Context, req *EnableDisableSubscriptionRequest) (*paystackapi.Response[interface{}], error) {
	resp := &paystackapi.Response[interface{}]{}
	err := c.backend.Call(ctx, "POST", "/subscription/disable", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
