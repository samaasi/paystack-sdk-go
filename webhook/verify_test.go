package webhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/http"
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

func TestIsFromPaystackIP_XForwardedFor(t *testing.T) {
	cases := []struct {
		name     string
		header   string
		expected bool
	}{
		{"first known IP", "52.31.239.247", true},
		{"second known IP", "52.89.246.173", true},
		{"third known IP", "52.214.14.220", true},
		{"unknown IP", "1.2.3.4", false},
		{"known IP behind proxy", "52.31.239.247, 10.0.0.1", true},
		{"unknown first, known second", "1.2.3.4, 52.31.239.247", false},
		{"empty header falls through to RemoteAddr", "", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := &http.Request{
				Header:     make(http.Header),
				RemoteAddr: "1.2.3.4:9999",
			}
			if tc.header != "" {
				req.Header.Set("X-Forwarded-For", tc.header)
			}
			got := IsFromPaystackIP(req)
			if got != tc.expected {
				t.Errorf("IsFromPaystackIP with X-Forwarded-For=%q: got %v, want %v", tc.header, got, tc.expected)
			}
		})
	}
}

func TestIsFromPaystackIP_RemoteAddr(t *testing.T) {
	cases := []struct {
		name       string
		remoteAddr string
		expected   bool
	}{
		{"known IP with port", "52.31.239.247:12345", true},
		{"second known IP with port", "52.89.246.173:9000", true},
		{"unknown IP with port", "9.9.9.9:80", false},
		{"IPv6 loopback", "[::1]:80", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := &http.Request{
				Header:     make(http.Header),
				RemoteAddr: tc.remoteAddr,
			}
			got := IsFromPaystackIP(req)
			if got != tc.expected {
				t.Errorf("IsFromPaystackIP with RemoteAddr=%q: got %v, want %v", tc.remoteAddr, got, tc.expected)
			}
		})
	}
}
