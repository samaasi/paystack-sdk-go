package webhook

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Verify verifies the Paystack webhook signature.
// It takes the secret key, the request body, and the signature header.
// Returns true if the signature is valid, false otherwise.
func Verify(secretKey string, body []byte, signature string) bool {
	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

// Parse parses the webhook request and verifies the signature.
// It returns the parsed event and any error encountered.
func Parse(r *http.Request, secretKey string, event *Event) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}
	defer r.Body.Close()

	signature := r.Header.Get("x-paystack-signature")
	if signature == "" {
		return fmt.Errorf("missing x-paystack-signature header")
	}

	if !Verify(secretKey, body, signature) {
		return fmt.Errorf("invalid signature")
	}

	if err := json.Unmarshal(body, event); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	return nil
}

// paystackIPs contains the official Paystack webhook IP addresses for production
var paystackIPs = []string{
	"52.31.239.247",
	"52.89.246.173",
	"52.214.14.220",
}

// IsFromPaystackIP checks if the incoming HTTP request is from a known Paystack IP address.
// This provides defense-in-depth security.
func IsFromPaystackIP(req *http.Request) bool {
	// Check X-Forwarded-For if behind a proxy
	forwardedFor := req.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		ips := strings.Split(forwardedFor, ",")
		if len(ips) > 0 {
			clientIP := strings.TrimSpace(ips[0])
			for _, ip := range paystackIPs {
				if clientIP == ip {
					return true
				}
			}
		}
	}

	// Fallback to RemoteAddr
	remoteAddr := req.RemoteAddr
	if idx := strings.LastIndex(remoteAddr, ":"); idx != -1 {
		remoteAddr = remoteAddr[:idx]
	}

	for _, ip := range paystackIPs {
		if remoteAddr == ip {
			return true
		}
	}

	return false
}
