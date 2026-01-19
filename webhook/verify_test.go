package webhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/http/httptest"
	"testing"
)

func TestVerify(t *testing.T) {
	secretKey := "secret"
	body := []byte(`{"event":"charge.success","data":{"id":1}}`)

	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	if !Verify(secretKey, body, signature) {
		t.Errorf("Verify failed for valid signature")
	}

	if Verify(secretKey, body, "invalid_signature") {
		t.Errorf("Verify passed for invalid signature")
	}
}

func TestParse(t *testing.T) {
	secretKey := "secret"
	body := []byte(`{"event":"charge.success","data":{"id":1}}`)

	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("x-paystack-signature", signature)

	var event Event
	if err := Parse(req, secretKey, &event); err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if event.Event != "charge.success" {
		t.Errorf("expected event charge.success, got %s", event.Event)
	}
}

func TestParseInvalidSignature(t *testing.T) {
	secretKey := "secret"
	body := []byte(`{"event":"charge.success","data":{"id":1}}`)

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("x-paystack-signature", "invalid_signature")

	var event Event
	if err := Parse(req, secretKey, &event); err == nil {
		t.Error("Parse should fail for invalid signature")
	}
}

func TestParseMissingSignature(t *testing.T) {
	secretKey := "secret"
	body := []byte(`{"event":"charge.success","data":{"id":1}}`)

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))

	var event Event
	if err := Parse(req, secretKey, &event); err == nil {
		t.Error("Parse should fail for missing signature")
	}
}
