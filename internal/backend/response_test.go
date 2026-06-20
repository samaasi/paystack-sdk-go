package backend

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/samaasi/paystack-sdk-go/v2/paystackapi"
)

func makeResp(statusCode int, body string, headers map[string]string) *http.Response {
	resp := &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
	for k, v := range headers {
		resp.Header.Set(k, v)
	}
	return resp
}

func TestDecode_Success(t *testing.T) {
	type result struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
	resp := makeResp(200, `{"status":true,"message":"ok"}`, nil)
	var out result
	if err := Decode(resp, &out); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !out.Status {
		t.Error("expected status true")
	}
	if out.Message != "ok" {
		t.Errorf("expected message ok, got %q", out.Message)
	}
}

func TestDecode_APIError_JSON(t *testing.T) {
	body := `{"status":false,"message":"Unauthorized","code":"auth_error"}`
	resp := makeResp(401, body, nil)
	err := Decode(resp, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var apiErr *paystackapi.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *paystackapi.APIError, got %T", err)
	}
	if apiErr.StatusCode != 401 {
		t.Errorf("expected StatusCode 401, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "Unauthorized" {
		t.Errorf("expected message Unauthorized, got %q", apiErr.Message)
	}
	if apiErr.Code != "auth_error" {
		t.Errorf("expected code auth_error, got %q", apiErr.Code)
	}
}

func TestDecode_APIError_ValidationErrors(t *testing.T) {
	body := `{"status":false,"message":"Validation failed","errors":{"amount":["is required"]}}`
	resp := makeResp(400, body, nil)
	err := Decode(resp, nil)
	var apiErr *paystackapi.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *paystackapi.APIError, got %T", err)
	}
	if msgs, ok := apiErr.Errors["amount"]; !ok || len(msgs) == 0 {
		t.Errorf("expected Errors[amount] to be populated, got %v", apiErr.Errors)
	}
}

func TestDecode_APIError_NonJSON(t *testing.T) {
	resp := makeResp(503, "Service Unavailable", nil)
	err := Decode(resp, nil)
	var apiErr *paystackapi.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *paystackapi.APIError, got %T", err)
	}
	if apiErr.StatusCode != 503 {
		t.Errorf("expected StatusCode 503, got %d", apiErr.StatusCode)
	}
}

func TestDecode_RetryAfterHeader(t *testing.T) {
	body := `{"status":false,"message":"Too Many Requests"}`
	resp := makeResp(429, body, map[string]string{"Retry-After": "30"})
	err := Decode(resp, nil)
	var apiErr *paystackapi.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *paystackapi.APIError, got %T", err)
	}
	if apiErr.RetryAfter != 30 {
		t.Errorf("expected RetryAfter 30, got %d", apiErr.RetryAfter)
	}
}

func TestDecode_NilDestination(t *testing.T) {
	resp := makeResp(200, `{"status":true}`, nil)
	if err := Decode(resp, nil); err != nil {
		t.Fatalf("expected no error with nil destination, got %v", err)
	}
}
