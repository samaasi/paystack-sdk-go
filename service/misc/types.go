package misc

import (
	"encoding/json"

	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

type Bank struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Code        string `json:"code"`
	Longcode    string `json:"longcode"`
	Gateway     string `json:"gateway"`
	PayWithBank bool   `json:"pay_with_bank"`
	Active      bool   `json:"active"`
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	Type        string `json:"type"`
	IsDeleted   bool   `json:"is_deleted"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListBanksResponse struct {
	paystackapi.Response[[]Bank]
}

type ListBanksParams struct {
	Country *string `query:"country,omitempty"`
	PerPage *int    `query:"perPage,omitempty"`
	Page    *int    `query:"page,omitempty"`
}

type Country struct {
	ID                  int             `json:"id"`
	Name                string          `json:"name"`
	ISOCode             string          `json:"iso_code"`
	DefaultCurrencyCode string          `json:"default_currency_code"`
	IntegrationDefaults json.RawMessage `json:"integration_defaults"`
	Relationships       json.RawMessage `json:"relationships"`
}

type ListCountriesResponse struct {
	paystackapi.Response[[]Country]
}

type State struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Abbreviation string `json:"abbreviation"`
}

type ListStatesResponse struct {
	paystackapi.Response[[]State]
}
