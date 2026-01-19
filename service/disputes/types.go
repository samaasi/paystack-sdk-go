package disputes

import (
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

// ListDisputesParams represents query parameters for listing disputes
type ListDisputesParams struct {
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	PerPage     int    `json:"perPage,omitempty"`
	Page        int    `json:"page,omitempty"`
	Transaction string `json:"transaction,omitempty"`
	Status      string `json:"status,omitempty"`
}

// UpdateDisputeRequest represents the request to update a dispute
type UpdateDisputeRequest struct {
	RefundAmount     string `json:"refund_amount"`
	UploadedFilename string `json:"uploaded_filename,omitempty"`
}

// AddEvidenceRequest represents the request to add evidence to a dispute
type AddEvidenceRequest struct {
	CustomerEmail   string `json:"customer_email"`
	CustomerName    string `json:"customer_name"`
	CustomerPhone   string `json:"customer_phone"`
	ServiceDetails  string `json:"service_details"`
	DeliveryAddress string `json:"delivery_address,omitempty"`
	DeliveryDate    string `json:"delivery_date,omitempty"`
}

// ResolveDisputeRequest represents the request to resolve a dispute
type ResolveDisputeRequest struct {
	Resolution       string `json:"resolution"` // "merchant-accepted", "declined"
	Message          string `json:"message"`
	RefundAmount     string `json:"refund_amount"`
	UploadedFilename string `json:"uploaded_filename"`
	Evidence         int    `json:"evidence,omitempty"`
}

// DisputeData represents the dispute object
type DisputeData struct {
	ID                   int             `json:"id"`
	RefundAmount         int             `json:"refund_amount"`
	Currency             string          `json:"currency"`
	Status               string          `json:"status"`
	Resolution           string          `json:"resolution"`
	Domain               string          `json:"domain"`
	Transaction          TransactionData `json:"transaction"`
	TransactionReference string          `json:"transaction_reference"`
	Category             string          `json:"category"`
	Customer             CustomerData    `json:"customer"`
	Bin                  string          `json:"bin"`
	Last4                string          `json:"last4"`
	DueAt                string          `json:"due_at"`
	ResolvedAt           string          `json:"resolved_at"`
	Evidence             EvidenceData    `json:"evidence"`
	History              []HistoryData   `json:"history"`
	Messages             []MessageData   `json:"messages"`
	CreatedAt            string          `json:"created_at"`
	UpdatedAt            string          `json:"updated_at"`
}

type TransactionData struct {
	ID        int    `json:"id"`
	Reference string `json:"reference"`
	Amount    int    `json:"amount"`
}

type CustomerData struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type EvidenceData struct {
	CustomerEmail   string `json:"customer_email"`
	CustomerName    string `json:"customer_name"`
	CustomerPhone   string `json:"customer_phone"`
	ServiceDetails  string `json:"service_details"`
	DeliveryAddress string `json:"delivery_address"`
	DeliveryDate    string `json:"delivery_date"`
}

type HistoryData struct {
	Status    string `json:"status"`
	By        string `json:"by"`
	CreatedAt string `json:"created_at"`
}

type MessageData struct {
	Sender    string `json:"sender"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

// DisputeResponse represents the generic response for dispute operations
type DisputeResponse struct {
	paystackapi.Response[DisputeData]
}

// DisputeListResponse represents the response for listing disputes
type DisputeListResponse struct {
	paystackapi.Response[[]DisputeData]
}

// UploadURLResponse represents the response for getting an upload URL
type UploadURLResponse struct {
	paystackapi.Response[UploadURLData]
}

type UploadURLData struct {
	SignedURL string `json:"signedUrl"`
	FileName  string `json:"fileName"`
}

// ExportDisputesResponse represents the response for exporting disputes
type ExportDisputesResponse struct {
	paystackapi.Response[ExportData]
}

type ExportData struct {
	Path string `json:"path"`
}
