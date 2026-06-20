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

func TestParseEvent_TransferSuccess(t *testing.T) {
	rawJSON := `{
		"event": "transfer.success",
		"data": {
			"amount": 50000,
			"currency": "NGN",
			"domain": "live",
			"id": 42,
			"reference": "TRF_ref_001",
			"transfer_code": "TRF_abc123",
			"status": "success",
			"source": "balance",
			"reason": "Salary payment",
			"transferred_at": "2023-01-15T10:00:00.000Z",
			"integration": {"id": 10, "is_live": true, "business_name": "Acme Ltd"},
			"recipient": {
				"id": 99,
				"type": "nuban",
				"name": "John Doe",
				"recipient_code": "RCP_001",
				"active": true,
				"currency": "NGN",
				"domain": "live",
				"is_deleted": false,
				"details": {"account_number": "0123456789", "bank_code": "011", "bank_name": "First Bank"},
				"created_at": "2022-01-01T00:00:00.000Z",
				"updated_at": "2022-01-01T00:00:00.000Z"
			},
			"failures": null,
			"session": null,
			"source_details": null,
			"titan_code": null,
			"created_at": "2023-01-15T09:59:00.000Z",
			"updated_at": "2023-01-15T10:00:00.000Z"
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}
	if event.Event != EventTransferSuccess {
		t.Errorf("expected %s, got %s", EventTransferSuccess, event.Event)
	}

	var data TransferSuccessEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}
	if data.Amount != 50000 {
		t.Errorf("expected amount 50000, got %d", data.Amount)
	}
	if data.TransferCode != "TRF_abc123" {
		t.Errorf("expected transfer code TRF_abc123, got %s", data.TransferCode)
	}
	if data.Recipient.RecipientCode != "RCP_001" {
		t.Errorf("expected recipient code RCP_001, got %s", data.Recipient.RecipientCode)
	}
}

func TestParseEvent_DisputeCreate(t *testing.T) {
	rawJSON := `{
		"event": "charge.dispute.create",
		"data": {
			"id": 7,
			"refund_amount": 5000,
			"currency": "NGN",
			"status": "awaiting-merchant-feedback",
			"resolution": "",
			"reference": "ref_dispute",
			"dispute_code": "DIS_001",
			"reason": "Customer complaint",
			"transaction": {
				"id": 123, "domain": "live", "status": "success",
				"reference": "ref_123", "amount": 5000,
				"customer": {"id": 1, "email": "j@example.com"},
				"authorization": {}
			},
			"created_at": "2023-01-01T00:00:00.000Z",
			"updated_at": "2023-01-01T00:00:00.000Z"
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}
	if event.Event != EventChargeDisputeCreate {
		t.Errorf("expected %s, got %s", EventChargeDisputeCreate, event.Event)
	}

	var data DisputeEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}
	if data.DisputeCode != "DIS_001" {
		t.Errorf("expected DisputeCode DIS_001, got %s", data.DisputeCode)
	}
	if data.RefundAmount != 5000 {
		t.Errorf("expected RefundAmount 5000, got %d", data.RefundAmount)
	}
}

func TestParseEvent_InvoiceCreate(t *testing.T) {
	rawJSON := `{
		"event": "invoice.create",
		"data": {
			"id": 55,
			"domain": "test",
			"invoice_code": "INV_001",
			"amount": 10000,
			"period_start": "2023-01-01T00:00:00.000Z",
			"period_end": "2023-02-01T00:00:00.000Z",
			"status": "pending",
			"paid": false,
			"paid_at": "",
			"description": "Monthly subscription",
			"authorization": {},
			"customer": {"id": 1, "email": "c@example.com"},
			"created_at": "2023-01-01T00:00:00.000Z"
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}

	var data InvoiceEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}
	if data.InvoiceCode != "INV_001" {
		t.Errorf("expected InvoiceCode INV_001, got %s", data.InvoiceCode)
	}
	if data.Amount != 10000 {
		t.Errorf("expected Amount 10000, got %d", data.Amount)
	}
}

func TestParseEvent_SubscriptionCreate(t *testing.T) {
	rawJSON := `{
		"event": "subscription.create",
		"data": {
			"id": 77,
			"domain": "test",
			"amount": 2000,
			"period_start": "2023-01-01T00:00:00.000Z",
			"period_end": "2023-02-01T00:00:00.000Z",
			"status": "active",
			"subscription_code": "SUB_abc",
			"email_token": "tok_abc123",
			"cron_expression": "0 0 1 * *",
			"next_payment_date": "2023-02-01T00:00:00.000Z",
			"created_at": "2023-01-01T00:00:00.000Z",
			"plan": {"id": 1, "name": "Monthly"},
			"authorization": {},
			"customer": {"id": 1, "email": "c@example.com"}
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}
	if event.Event != EventSubscriptionCreate {
		t.Errorf("expected %s, got %s", EventSubscriptionCreate, event.Event)
	}

	var data SubscriptionEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}
	if data.SubscriptionCode != "SUB_abc" {
		t.Errorf("expected SubscriptionCode SUB_abc, got %s", data.SubscriptionCode)
	}
	if data.EmailToken != "tok_abc123" {
		t.Errorf("expected EmailToken tok_abc123, got %s", data.EmailToken)
	}
	if data.Plan == nil {
		t.Error("expected Plan to be non-nil json.RawMessage")
	}
}

func TestParseEvent_RefundProcessed(t *testing.T) {
	rawJSON := `{
		"event": "refund.processed",
		"data": {
			"id": 10,
			"refunded_by": "admin@test.com",
			"refunded_at": "2023-01-05T00:00:00.000Z",
			"amount": 3000,
			"currency": "NGN",
			"transaction_reference": "ref_orig",
			"status": "processed",
			"deducted_amount": 3000,
			"fully_deducted": true,
			"refund_reference": "RFD_001",
			"customer": {"id": 1, "email": "c@example.com"}
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}
	if event.Event != EventRefundProcessed {
		t.Errorf("expected %s, got %s", EventRefundProcessed, event.Event)
	}

	var data RefundEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}
	if data.RefundReference != "RFD_001" {
		t.Errorf("expected RefundReference RFD_001, got %s", data.RefundReference)
	}
	if !data.FullyDeducted {
		t.Error("expected FullyDeducted to be true")
	}
}

func TestParseEvent_DedicatedAccountAssign(t *testing.T) {
	rawJSON := `{
		"event": "dedicatedaccount.assign.success",
		"data": {
			"id": 15,
			"account_name": "John Doe",
			"account_number": "9876543210",
			"assigned": true,
			"currency": "NGN",
			"active": true,
			"bank": {"account_number": "9876543210", "bank_code": "035", "bank_name": "Wema Bank"},
			"created_at": "2023-01-01T00:00:00.000Z",
			"updated_at": "2023-01-01T00:00:00.000Z",
			"customer": {"id": 1, "email": "c@example.com"}
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}

	var data DedicatedAccountEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}
	if data.AccountNumber != "9876543210" {
		t.Errorf("expected AccountNumber 9876543210, got %s", data.AccountNumber)
	}
	if !data.Assigned {
		t.Error("expected Assigned to be true")
	}
}

func TestParseEvent_CustomerIdentification(t *testing.T) {
	rawJSON := `{
		"event": "customeridentification.success",
		"data": {
			"customer_code": "CUS_abc",
			"email": "c@example.com",
			"reason": "Verified",
			"identification": {"country": "NG", "type": "bvn", "value": "12345678901"}
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}

	var data CustomerIdentificationEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}
	if data.CustomerCode != "CUS_abc" {
		t.Errorf("expected CustomerCode CUS_abc, got %s", data.CustomerCode)
	}
	if data.Identification == nil {
		t.Error("expected Identification to be non-nil json.RawMessage")
	}
}

func TestParseEvent_PaymentRequest(t *testing.T) {
	rawJSON := `{
		"event": "paymentrequest.success",
		"data": {
			"id": 88,
			"domain": "test",
			"amount": 15000,
			"currency": "NGN",
			"request_code": "PRQ_001",
			"status": "success",
			"paid": true,
			"paid_at": "2023-01-10T00:00:00.000Z",
			"has_invoice": true,
			"invoice_number": "INV-001",
			"description": "Service fee",
			"line_items": [{"name": "Consulting", "amount": 15000}],
			"tax": [],
			"notifications": [],
			"offline_reference": "offline_001",
			"customer": {"id": 1, "email": "c@example.com"},
			"created_at": "2023-01-01T00:00:00.000Z"
		}
	}`

	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent failed: %v", err)
	}

	var data PaymentRequestEvent
	if err := event.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData failed: %v", err)
	}
	if data.RequestCode != "PRQ_001" {
		t.Errorf("expected RequestCode PRQ_001, got %s", data.RequestCode)
	}
	if !data.Paid {
		t.Error("expected Paid to be true")
	}
}

func TestParseEvent_UnknownType(t *testing.T) {
	rawJSON := `{"event":"unknown.event","data":{"foo":"bar"}}`
	event, err := ParseEvent([]byte(rawJSON))
	if err != nil {
		t.Fatalf("ParseEvent should not error on unknown event types: %v", err)
	}
	if event.Event != "unknown.event" {
		t.Errorf("expected event unknown.event, got %s", event.Event)
	}
}

func TestParseEvent_InvalidJSON(t *testing.T) {
	_, err := ParseEvent([]byte(`{invalid`))
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}
