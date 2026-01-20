package status

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

const DefaultStatusURL = "https://status.paystack.com/api/v2/summary.json"

// Client is the client for the Status service
type Client struct {
	httpClient *http.Client
	url        string
}

// NewClient creates a new Status client
// backend param is kept for compatibility but not used for HTTP calls
func NewClient(backend *backend.Client) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		url:        DefaultStatusURL,
	}
}

// WithURL sets a custom URL for the status check (useful for testing)
func (c *Client) WithURL(url string) *Client {
	c.url = url
	return c
}

// Fetch fetches the system status summary
func (c *Client) Fetch(ctx context.Context) (*Summary, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status check failed with code: %d", resp.StatusCode)
	}

	var summary Summary
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &summary, nil
}
