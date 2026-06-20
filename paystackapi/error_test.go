package paystackapi

import (
	"errors"
	"testing"
)

func TestAPIError_Error_WithCode(t *testing.T) {
	e := &APIError{Message: "Unauthorized", Code: "auth_error", StatusCode: 401}
	got := e.Error()
	if got != "paystack: Unauthorized (code: auth_error, status: 401)" {
		t.Errorf("unexpected error string: %q", got)
	}
}

func TestAPIError_Error_WithoutCode(t *testing.T) {
	e := &APIError{Message: "Server error", StatusCode: 500}
	got := e.Error()
	if got != "paystack: Server error (status: 500)" {
		t.Errorf("unexpected error string: %q", got)
	}
}

func TestAPIError_Is_StatusCode(t *testing.T) {
	e := &APIError{StatusCode: 404, Message: "Not found"}
	target := &APIError{StatusCode: 404}
	if !errors.Is(e, target) {
		t.Error("expected Is() to match on StatusCode")
	}
}

func TestAPIError_Is_Code(t *testing.T) {
	e := &APIError{StatusCode: 400, Code: "invalid_amount", Message: "bad"}
	target := &APIError{Code: "invalid_amount"}
	if !errors.Is(e, target) {
		t.Error("expected Is() to match on Code")
	}
}

func TestAPIError_Is_Mismatch(t *testing.T) {
	e := &APIError{StatusCode: 404}
	target := &APIError{StatusCode: 500}
	if errors.Is(e, target) {
		t.Error("expected Is() to not match different StatusCode")
	}
}

func TestAPIError_Is_NonAPIError(t *testing.T) {
	e := &APIError{StatusCode: 404}
	if errors.Is(e, errors.New("some other error")) {
		t.Error("expected Is() to return false for non-APIError target")
	}
}

func TestAPIError_Errors_Field(t *testing.T) {
	e := &APIError{
		Message: "Validation failed",
		StatusCode: 400,
		Errors: map[string][]string{
			"amount": {"is required"},
			"email":  {"is invalid"},
		},
	}
	if len(e.Errors["amount"]) != 1 || e.Errors["amount"][0] != "is required" {
		t.Errorf("unexpected Errors field: %v", e.Errors)
	}
}

func TestRequestError_Error(t *testing.T) {
	inner := errors.New("connection refused")
	e := &RequestError{Err: inner}
	if e.Error() != "paystack request failed: connection refused" {
		t.Errorf("unexpected error string: %q", e.Error())
	}
}

func TestRequestError_Unwrap(t *testing.T) {
	inner := errors.New("timeout")
	e := &RequestError{Err: inner}
	if !errors.Is(e, inner) {
		t.Error("expected Unwrap to expose inner error")
	}
}
