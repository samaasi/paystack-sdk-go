package verification

import "github.com/samaasi/paystack-sdk-go/paystackapi"

// ResolveAccountResponse represents the response for account resolution
type ResolveAccountResponse struct {
	paystackapi.Response[AccountResolution]
}

// AccountResolution represents the resolved account details
type AccountResolution struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankID        int    `json:"bank_id"`
}

// ValidateAccountRequest represents the request to validate an account
type ValidateAccountRequest struct {
	BankCode       string `json:"bank_code"`
	CountryCode    string `json:"country_code"`
	AccountNumber  string `json:"account_number"`
	AccountName    string `json:"account_name"`
	AccountType    string `json:"account_type"`
	DocumentType   string `json:"document_type"`
	DocumentNumber string `json:"document_number"`
}

// ValidateAccountResponse represents the response for account validation
type ValidateAccountResponse struct {
	paystackapi.Response[AccountValidation]
}

// AccountValidation represents the validation result
type AccountValidation struct {
	Verified bool   `json:"verified"`
	Message  string `json:"message"`
}

// ResolveCardBINResponse represents the response for card BIN resolution
type ResolveCardBINResponse struct {
	paystackapi.Response[CardBIN]
}

// CardBIN represents card BIN details
type CardBIN struct {
	Bin          string `json:"bin"`
	Brand        string `json:"brand"`
	SubBrand     string `json:"sub_brand"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	CardType     string `json:"card_type"`
	Bank         string `json:"bank"`
	LinkedBankID int    `json:"linked_bank_id"`
}
