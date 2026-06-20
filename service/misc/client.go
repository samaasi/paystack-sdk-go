package misc

import (
	"context"
	"fmt"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

// Service represents the interface for misc operations.
type Service interface {
	ListBanks(ctx context.Context, params *ListBanksParams) (*ListBanksResponse, error)
	ListCountries(ctx context.Context) (*ListCountriesResponse, error)
	ListStates(ctx context.Context, country string) (*ListStatesResponse, error)
}

type Client struct {
	backend *backend.Client
}

func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

func (c *Client) ListBanks(ctx context.Context, params *ListBanksParams) (*ListBanksResponse, error) {
	path := "/bank"
	if params != nil {
		query, err := backend.EncodeQueryParams(params)
		if err != nil {
			return nil, err
		}
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
	}

	resp := &ListBanksResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ListCountries(ctx context.Context) (*ListCountriesResponse, error) {
	resp := &ListCountriesResponse{}
	err := c.backend.Call(ctx, "GET", "/country", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ListStates(ctx context.Context, country string) (*ListStatesResponse, error) {
	path := fmt.Sprintf("/address_verification/states?country=%s", country)
	resp := &ListStatesResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
