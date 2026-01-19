package paymentrequests

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Client is the client for the Payment Requests service
type Client struct {
	backend *backend.Client
}

// NewClient creates a new Payment Requests client
func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

// Create creates a new payment request
func (c *Client) Create(ctx context.Context, req *CreatePaymentRequestRequest) (*PaymentRequestResponse, error) {
	resp := &PaymentRequestResponse{}
	err := c.backend.Call(ctx, "POST", "/paymentrequest", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List lists payment requests
func (c *Client) List(ctx context.Context) (*ListPaymentRequestsResponse, error) {
	resp := &ListPaymentRequestsResponse{}
	err := c.backend.Call(ctx, "GET", "/paymentrequest", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Fetch fetches a payment request by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*PaymentRequestResponse, error) {
	path := fmt.Sprintf("/paymentrequest/%s", idOrCode)
	resp := &PaymentRequestResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Verify verifies a payment request
func (c *Client) Verify(ctx context.Context, code string) (*PaymentRequestResponse, error) {
	path := fmt.Sprintf("/paymentrequest/verify/%s", code)
	resp := &PaymentRequestResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendNotification sends a notification for a payment request
func (c *Client) SendNotification(ctx context.Context, code string) (*PaymentRequestResponse, error) {
	path := fmt.Sprintf("/paymentrequest/notify/%s", code)
	resp := &PaymentRequestResponse{}
	err := c.backend.Call(ctx, "POST", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Total fetches the total payment requests
func (c *Client) Total(ctx context.Context) (*PaymentRequestTotalResponse, error) {
	resp := &PaymentRequestTotalResponse{}
	err := c.backend.Call(ctx, "GET", "/paymentrequest/total", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Finalize finalizes a payment request
func (c *Client) Finalize(ctx context.Context, code string) (*PaymentRequestResponse, error) {
	path := fmt.Sprintf("/paymentrequest/finalize/%s", code)
	resp := &PaymentRequestResponse{}
	err := c.backend.Call(ctx, "POST", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Update updates a payment request
func (c *Client) Update(ctx context.Context, idOrCode string, req *UpdatePaymentRequestRequest) (*PaymentRequestResponse, error) {
	path := fmt.Sprintf("/paymentrequest/%s", idOrCode)
	resp := &PaymentRequestResponse{}
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Archive archives a payment request
func (c *Client) Archive(ctx context.Context, code string) (*PaymentRequestResponse, error) {
	path := fmt.Sprintf("/paymentrequest/archive/%s", code)
	resp := &PaymentRequestResponse{}
	err := c.backend.Call(ctx, "POST", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
