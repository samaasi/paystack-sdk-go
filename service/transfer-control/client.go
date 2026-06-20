package transferControl

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/v2/internal/backend"
)

// Service represents the interface for transfer-control operations.
type Service interface {
	CheckBalance(ctx context.Context) (*CheckBalanceResponse, error)
	FetchLedger(ctx context.Context, params *LedgerParams) (*LedgerResponse, error)
	ResendOTP(ctx context.Context, req *ResendOTPRequest) (*ResendOTPResponse, error)
	DisableOTP(ctx context.Context) (*DisableOTPResponse, error)
	FinalizeDisableOTP(ctx context.Context, req *FinalizeDisableOTPRequest) (*FinalizeDisableOTPResponse, error)
	EnableOTP(ctx context.Context) (*EnableOTPResponse, error)
}

type Client struct {
	backend *backend.Client
}

func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

func (c *Client) CheckBalance(ctx context.Context) (*CheckBalanceResponse, error) {
	resp := &CheckBalanceResponse{}
	err := c.backend.Call(ctx, "GET", "/balance", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FetchLedger returns a paginated history of balance transactions (ledger)
func (c *Client) FetchLedger(ctx context.Context, params *LedgerParams) (*LedgerResponse, error) {
	path := "/balance/ledger"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}
	resp := &LedgerResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ResendOTP(ctx context.Context, req *ResendOTPRequest) (*ResendOTPResponse, error) {
	resp := &ResendOTPResponse{}
	err := c.backend.Call(ctx, "POST", "/transfer/resend_otp", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) DisableOTP(ctx context.Context) (*DisableOTPResponse, error) {
	resp := &DisableOTPResponse{}
	err := c.backend.Call(ctx, "POST", "/transfer/disable_otp", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) FinalizeDisableOTP(ctx context.Context, req *FinalizeDisableOTPRequest) (*FinalizeDisableOTPResponse, error) {
	resp := &FinalizeDisableOTPResponse{}
	err := c.backend.Call(ctx, "POST", "/transfer/disable_otp_finalize", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) EnableOTP(ctx context.Context) (*EnableOTPResponse, error) {
	resp := &EnableOTPResponse{}
	err := c.backend.Call(ctx, "POST", "/transfer/enable_otp", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
