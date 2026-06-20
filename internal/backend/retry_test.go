package backend

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/samaasi/paystack-sdk-go/v2/paystackapi"
)

func TestRetry_SuccessOnFirstAttempt(t *testing.T) {
	b := &Backoff{MaxRetries: 3, MinDelay: time.Millisecond, MaxDelay: 10 * time.Millisecond}
	calls := 0
	err := b.Retry(context.Background(), func() error {
		calls++
		return nil
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if calls != 1 {
		t.Errorf("expected 1 call, got %d", calls)
	}
}

func TestRetry_RetriesOn5xx(t *testing.T) {
	b := &Backoff{MaxRetries: 2, MinDelay: time.Millisecond, MaxDelay: 5 * time.Millisecond}
	calls := 0
	serverErr := &paystackapi.APIError{StatusCode: http.StatusInternalServerError, Message: "server error"}
	err := b.Retry(context.Background(), func() error {
		calls++
		if calls < 3 {
			return serverErr
		}
		return nil
	})
	if err != nil {
		t.Fatalf("expected success after retries, got %v", err)
	}
	if calls != 3 {
		t.Errorf("expected 3 calls, got %d", calls)
	}
}

func TestRetry_NoRetryOn4xx(t *testing.T) {
	b := &Backoff{MaxRetries: 3, MinDelay: time.Millisecond, MaxDelay: 5 * time.Millisecond}
	calls := 0
	clientErr := &paystackapi.APIError{StatusCode: http.StatusBadRequest, Message: "bad request"}
	err := b.Retry(context.Background(), func() error {
		calls++
		return clientErr
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if calls != 1 {
		t.Errorf("expected exactly 1 call (no retry on 4xx), got %d", calls)
	}
}

func TestRetry_Retries429(t *testing.T) {
	b := &Backoff{MaxRetries: 2, MinDelay: time.Millisecond, MaxDelay: 5 * time.Millisecond}
	calls := 0
	rateLimitErr := &paystackapi.APIError{StatusCode: http.StatusTooManyRequests, Message: "rate limited", RetryAfter: 0}
	err := b.Retry(context.Background(), func() error {
		calls++
		if calls < 2 {
			return rateLimitErr
		}
		return nil
	})
	if err != nil {
		t.Fatalf("expected success after 429 retry, got %v", err)
	}
	if calls != 2 {
		t.Errorf("expected 2 calls, got %d", calls)
	}
}

func TestRetry_ExhaustsMaxRetries(t *testing.T) {
	b := &Backoff{MaxRetries: 2, MinDelay: time.Millisecond, MaxDelay: 5 * time.Millisecond}
	calls := 0
	serverErr := &paystackapi.APIError{StatusCode: 500, Message: "always fails"}
	err := b.Retry(context.Background(), func() error {
		calls++
		return serverErr
	})
	if err == nil {
		t.Fatal("expected error after exhausting retries")
	}
	if calls != 3 { // initial attempt + 2 retries
		t.Errorf("expected 3 calls (1 + 2 retries), got %d", calls)
	}
}

func TestRetry_ContextCancellation(t *testing.T) {
	b := &Backoff{MaxRetries: 5, MinDelay: 100 * time.Millisecond, MaxDelay: 500 * time.Millisecond}
	ctx, cancel := context.WithCancel(context.Background())

	calls := 0
	serverErr := &paystackapi.APIError{StatusCode: 500, Message: "server error"}
	done := make(chan error, 1)
	go func() {
		done <- b.Retry(ctx, func() error {
			calls++
			if calls == 1 {
				cancel()
			}
			return serverErr
		})
	}()

	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expected context.Canceled, got %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("test timed out — retry did not respect context cancellation")
	}
}

func TestIsRetryable(t *testing.T) {
	cases := []struct {
		name     string
		err      error
		expected bool
	}{
		{"5xx retryable", &paystackapi.APIError{StatusCode: 500}, true},
		{"429 retryable", &paystackapi.APIError{StatusCode: 429}, true},
		{"400 not retryable", &paystackapi.APIError{StatusCode: 400}, false},
		{"404 not retryable", &paystackapi.APIError{StatusCode: 404}, false},
		{"non-API error retryable", errors.New("network error"), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isRetryable(tc.err); got != tc.expected {
				t.Errorf("isRetryable(%v) = %v, want %v", tc.err, got, tc.expected)
			}
		})
	}
}

func TestCalculateDelay(t *testing.T) {
	b := &Backoff{MinDelay: 100 * time.Millisecond, MaxDelay: 5 * time.Second}
	d0 := b.calculateDelay(0)
	d1 := b.calculateDelay(1)
	d2 := b.calculateDelay(2)

	if d0 < 100*time.Millisecond {
		t.Errorf("attempt 0 delay %v too small", d0)
	}
	if d1 <= d0 {
		t.Errorf("expected delay to grow: attempt 1 (%v) <= attempt 0 (%v)", d1, d0)
	}
	if d2 > 5*time.Second {
		t.Errorf("attempt 2 delay %v exceeds MaxDelay", d2)
	}
}
