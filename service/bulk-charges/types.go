package bulkcharges

import "github.com/samaasi/paystack-sdk-go/paystackapi"

// BulkChargeUnit represents a single charge in a bulk request
type BulkChargeUnit struct {
	Authorization string `json:"authorization"`
	Amount        int    `json:"amount"`
	Reference     string `json:"reference,omitempty"`
}

// InitiateBulkChargeRequest represents the request to initiate a bulk charge
type InitiateBulkChargeRequest []BulkChargeUnit

// InitiateBulkChargeResponse represents the response for initiating a bulk charge
type InitiateBulkChargeResponse struct {
	paystackapi.Response[BulkChargeData]
}

// BulkChargeData represents the data returned after initiating a bulk charge
type BulkChargeData struct {
	BatchCode string `json:"batch_code"`
	Reference string `json:"reference"`
	ID        int    `json:"id"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

// ListBulkChargesResponse represents the response for listing bulk charges
type ListBulkChargesResponse struct {
	paystackapi.Response[[]BulkChargeData]
}

// ListBulkChargesParams represents query parameters for listing bulk charges
type ListBulkChargesParams struct {
	PerPage int    `json:"perPage,omitempty"`
	Page    int    `json:"page,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
}

// FetchBulkChargeResponse represents the response for fetching a bulk charge batch
type FetchBulkChargeResponse struct {
	paystackapi.Response[BulkChargeBatchDetails]
}

// BulkChargeBatchDetails represents detailed information about a bulk charge batch
type BulkChargeBatchDetails struct {
	Domain         string `json:"domain"`
	BatchCode      string `json:"batch_code"`
	Status         string `json:"status"`
	ID             int    `json:"id"`
	TotalCharges   int    `json:"total_charges"`
	PendingCharges int    `json:"pending_charges"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// FetchChargesInBatchResponse represents the response for fetching charges in a batch
type FetchChargesInBatchResponse struct {
	paystackapi.Response[[]BulkChargeUnitDetails]
}

// FetchChargesInBatchParams represents query parameters for fetching charges in a batch
type FetchChargesInBatchParams struct {
	Status  string `json:"status,omitempty"`
	PerPage int    `json:"perPage,omitempty"`
	Page    int    `json:"page,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
}

// BulkChargeUnitDetails represents details of a specific charge in a bulk batch
type BulkChargeUnitDetails struct {
	Status        string `json:"status"`
	Reference     string `json:"reference"`
	Amount        int    `json:"amount"`
	Authorization string `json:"authorization_code"`
	Message       string `json:"message"`
}

// PauseBulkChargeResponse represents the response for pausing a bulk charge batch
type PauseBulkChargeResponse struct {
	paystackapi.Response[interface{}]
}

// ResumeBulkChargeResponse represents the response for resuming a bulk charge batch
type ResumeBulkChargeResponse struct {
	paystackapi.Response[interface{}]
}
