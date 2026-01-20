package webhook

import (
	"testing"
)

func TestParseEvent(t *testing.T) {
	rawJSON := `{
		"event": "charge.success",
		"data": {
			"id": 123,
			"domain": "test",
			"status": "success",
			"reference": "ref_123",
			"amount": 5000,
			"gateway_response": "Successful",
			"paid_at": "2020-01-01T12:00:00.000Z",
			"created_at": "2020-01-01T12:00:00.000Z",
			"channel": "card",
			"currency": "NGN",
			"ip_address": "127.0.0.1",
			"customer": {
				"id": 1,
				"first_name": "John",
				"last_name": "Doe",
				"email": "john@example.com",
				"customer_code": "CUS_123"
			},
			"authorization": {
				"authorization_code": "AUTH_123",
				"bin": "408408",
				"last4": "4081",
				"exp_month": "12",
				"exp_year": "2030",
				"channel": "card",
				"card_type": "visa",
				"bank": "test bank",
				"country_code": "NG",
				"brand": "visa",
				"reusable": true,
				"signature": "SIG_123"
			}
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}

	if event.Event != EventChargeSuccess {
		t.Errorf("Expected event type %s, got %s", EventChargeSuccess, event.Event)
	}

	var data ChargeSuccessEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}

	if data.ID != 123 {
		t.Errorf("Expected ID 123, got %d", data.ID)
	}
	if data.Reference != "ref_123" {
		t.Errorf("Expected reference 'ref_123', got '%s'", data.Reference)
	}
	if data.Customer.Email != "john@example.com" {
		t.Errorf("Expected customer email 'john@example.com', got '%s'", data.Customer.Email)
	}
}
