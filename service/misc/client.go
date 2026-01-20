package misc

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

type Client struct {
	backend *backend.Client
}

func NewClient(backend *backend.Client) *Client {
	return &Client{backend: backend}
}

func (c *Client) ListBanks(ctx context.Context, country string, perPage, page int) (*ListBanksResponse, error) {
	v := url.Values{}
	if country != "" {
		v.Add("country", country)
	}
	if perPage > 0 {
		v.Add("perPage", strconv.Itoa(perPage))
	}
	if page > 0 {
		v.Add("page", strconv.Itoa(page))
	}
	
	path := "/bank"
	if len(v) > 0 {
		path = fmt.Sprintf("%s?%s", path, v.Encode())
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

func (c *Client) ResolveCardBIN(ctx context.Context, bin string) (*ResolveCardBINResponse, error) {
	path := fmt.Sprintf("/decision/bin/%s", bin)
	resp := &ResolveCardBINResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ResolveAccount(ctx context.Context, accountNumber, bankCode string) (*ResolveAccountResponse, error) {
	path := fmt.Sprintf("/bank/resolve?account_number=%s&bank_code=%s", accountNumber, bankCode)
	resp := &ResolveAccountResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
