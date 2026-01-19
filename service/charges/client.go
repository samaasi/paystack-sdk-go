package charges

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

// Create initiates a payment by charging a customer's card or bank account
func (c *Client) Create(ctx context.Context, req *CreateChargeRequest) (*CreateChargeResponse, error) {
	resp := &CreateChargeResponse{}
	err := c.backend.Call(ctx, "POST", "/charge", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SubmitPIN submits a PIN for a charge
func (c *Client) SubmitPIN(ctx context.Context, req *SubmitPINRequest) (*SubmitResponse, error) {
	resp := &SubmitResponse{}
	err := c.backend.Call(ctx, "POST", "/charge/submit_pin", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SubmitOTP submits an OTP for a charge
func (c *Client) SubmitOTP(ctx context.Context, req *SubmitOTPRequest) (*SubmitResponse, error) {
	resp := &SubmitResponse{}
	err := c.backend.Call(ctx, "POST", "/charge/submit_otp", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SubmitPhone submits a phone number for a charge
func (c *Client) SubmitPhone(ctx context.Context, req *SubmitPhoneRequest) (*SubmitResponse, error) {
	resp := &SubmitResponse{}
	err := c.backend.Call(ctx, "POST", "/charge/submit_phone", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SubmitBirthday submits a birthday for a charge
func (c *Client) SubmitBirthday(ctx context.Context, req *SubmitBirthdayRequest) (*SubmitResponse, error) {
	resp := &SubmitResponse{}
	err := c.backend.Call(ctx, "POST", "/charge/submit_birthday", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SubmitAddress submits an address for a charge
func (c *Client) SubmitAddress(ctx context.Context, req *SubmitAddressRequest) (*SubmitResponse, error) {
	resp := &SubmitResponse{}
	err := c.backend.Call(ctx, "POST", "/charge/submit_address", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CheckPending checks the status of a pending charge
func (c *Client) CheckPending(ctx context.Context, reference string) (*CheckPendingChargeResponse, error) {
	resp := &CheckPendingChargeResponse{}
	err := c.backend.Call(ctx, "GET", fmt.Sprintf("/charge/%s", reference), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
