package backend

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// Backoff handles retry logic with exponential backoff
type Backoff struct {
	MaxRetries int
	MinDelay   time.Duration
	MaxDelay   time.Duration
}

// DefaultBackoff returns a default configuration for backoff
func DefaultBackoff() *Backoff {
	return &Backoff{
		MaxRetries: 3,
		MinDelay:   1 * time.Second,
		MaxDelay:   5 * time.Second,
	}
}

// Retry executes the operation with retries
func (b *Backoff) Retry(ctx context.Context, op func() error) error {
	var err error
	for i := 0; i <= b.MaxRetries; i++ {
		err = op()
		if err == nil {
			return nil
		}

		if !isRetryable(err) {
			return err
		}

		if i == b.MaxRetries {
			break
		}

		delay := b.calculateDelay(i)
		var apiErr *paystackapi.APIError
		if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusTooManyRequests && apiErr.RetryAfter > 0 {
			delay = time.Duration(apiErr.RetryAfter) * time.Second
		}

		select {
		case <-time.After(delay):
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return err
}

func isRetryable(err error) bool {
	var apiErr *paystackapi.APIError
	if errors.As(err, &apiErr) {
		if apiErr.StatusCode >= 400 && apiErr.StatusCode < 500 && apiErr.StatusCode != http.StatusTooManyRequests {
			return false
		}
		return true
	}

	return true
}

func (b *Backoff) calculateDelay(attempt int) time.Duration {
	delay := float64(b.MinDelay) * math.Pow(2, float64(attempt))
	// Add jitter
	jitter := rand.Float64() * float64(delay) * 0.1
	delay += jitter

	if delay > float64(b.MaxDelay) {
		delay = float64(b.MaxDelay)
	}

	return time.Duration(delay)
}
