package misc

import "github.com/samaasi/paystack-sdk-go/paystackapi"

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
	Country string `query:"country"`
	PerPage int    `query:"perPage"`
	Page    int    `query:"page"`
}

type Country struct {
	ID                  int         `json:"id"`
	Name                string      `json:"name"`
	ISOCode             string      `json:"iso_code"`
	DefaultCurrencyCode string      `json:"default_currency_code"`
	IntegrationDefaults interface{} `json:"integration_defaults"`
	Relationships       interface{} `json:"relationships"`
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

type CardBIN struct {
	Bin          string `json:"bin"`
	Brand        string `json:"brand"`
	SubBrand     string `json:"sub_brand"`
	Type         string `json:"type"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	Bank         string `json:"bank"`
	LinkedBankID int    `json:"linked_bank_id"`
}

type ResolveCardBINResponse struct {
	paystackapi.Response[CardBIN]
}

type AccountResolve struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankID        int    `json:"bank_id"`
}

type ResolveAccountResponse struct {
	paystackapi.Response[AccountResolve]
}
