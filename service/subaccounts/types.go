package subaccounts

import (
	"encoding/json"

	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

type Subaccount struct {
	ID                  int             `json:"id"`
	SubaccountCode      string          `json:"subaccount_code"`
	BusinessName        string          `json:"business_name"`
	Description         string          `json:"description"`
	PrimaryContactName  string          `json:"primary_contact_name"`
	PrimaryContactEmail string          `json:"primary_contact_email"`
	PrimaryContactPhone string          `json:"primary_contact_phone"`
	Metadata            json.RawMessage `json:"metadata"`
	PercentageCharge    float64         `json:"percentage_charge"`
	SettlementBank      string          `json:"settlement_bank"`
	AccountNumber       string          `json:"account_number"`
	SettlementSchedule  string          `json:"settlement_schedule"`
	Active              bool            `json:"active"`
	Migrate             bool            `json:"migrate"`
	Currency            string          `json:"currency"`
	CreatedAt           string          `json:"created_at"`
	UpdatedAt           string          `json:"updated_at"`
}

type CreateSubaccountRequest struct {
	BusinessName        string                 `json:"business_name"`
	SettlementBank      string                 `json:"settlement_bank"`
	AccountNumber       string                 `json:"account_number"`
	PercentageCharge    float64                `json:"percentage_charge"`
	Description         string                 `json:"description,omitempty"`
	PrimaryContactName  string                 `json:"primary_contact_name,omitempty"`
	PrimaryContactEmail string                 `json:"primary_contact_email,omitempty"`
	PrimaryContactPhone string                 `json:"primary_contact_phone,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateSubaccountRequest struct {
	BusinessName        string                 `json:"business_name,omitempty"`
	SettlementBank      string                 `json:"settlement_bank,omitempty"`
	AccountNumber       string                 `json:"account_number,omitempty"`
	PercentageCharge    float64                `json:"percentage_charge,omitempty"`
	Description         string                 `json:"description,omitempty"`
	PrimaryContactName  string                 `json:"primary_contact_name,omitempty"`
	PrimaryContactEmail string                 `json:"primary_contact_email,omitempty"`
	PrimaryContactPhone string                 `json:"primary_contact_phone,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
	Active              *bool                  `json:"active,omitempty"`
}

type SubaccountResponse struct {
	paystackapi.Response[Subaccount]
}

type ListSubaccountsResponse struct {
	paystackapi.Response[[]Subaccount]
}
