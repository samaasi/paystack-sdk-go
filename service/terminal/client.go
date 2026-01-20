package terminal

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

func (c *Client) SendEvent(ctx context.Context, terminalID string, req *SendEventRequest) (*TerminalEventResponse, error) {
	path := fmt.Sprintf("/terminal/%s/event", terminalID)
	resp := &TerminalEventResponse{}
	err := c.backend.Call(ctx, "POST", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) FetchEventStatus(ctx context.Context, terminalID, eventID string) (*TerminalEventResponse, error) {
	path := fmt.Sprintf("/terminal/%s/event/%s", terminalID, eventID)
	resp := &TerminalEventResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) FetchPresence(ctx context.Context, terminalID string) (*TerminalPresenceResponse, error) {
	path := fmt.Sprintf("/terminal/%s/presence", terminalID)
	resp := &TerminalPresenceResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) List(ctx context.Context, perPage, page int) (*ListTerminalsResponse, error) {
	path := fmt.Sprintf("/terminal?perPage=%d&page=%d", perPage, page)
	resp := &ListTerminalsResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Fetch(ctx context.Context, terminalID string) (*TerminalResponse, error) {
	path := fmt.Sprintf("/terminal/%s", terminalID)
	resp := &TerminalResponse{}
	err := c.backend.Call(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Update(ctx context.Context, terminalID string, req *UpdateTerminalRequest) (*TerminalResponse, error) {
	path := fmt.Sprintf("/terminal/%s", terminalID)
	resp := &TerminalResponse{}
	err := c.backend.Call(ctx, "PUT", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Commission(ctx context.Context, serial string) (*TerminalResponse, error) {
	req := map[string]string{"serial": serial}
	resp := &TerminalResponse{}
	err := c.backend.Call(ctx, "POST", "/terminal/commission_device", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Decommission(ctx context.Context, serial string) (*TerminalResponse, error) {
	req := map[string]string{"serial": serial}
	resp := &TerminalResponse{}
	err := c.backend.Call(ctx, "POST", "/terminal/decommission_device", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
