package webhook

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Verify verifies the Paystack webhook signature.
// It takes the secret key, the request body, and the signature header.
// Returns true if the signature is valid, false otherwise.
func Verify(secretKey string, body []byte, signature string) bool {
	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	return expectedSignature == signature
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
