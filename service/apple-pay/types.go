package applepay

import "github.com/samaasi/paystack-sdk-go/paystackapi"

// RegisterDomainRequest represents the payload for registering a domain
type RegisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

// RegisterDomainResponse represents the response for registering a domain
type RegisterDomainResponse struct {
	paystackapi.Response[interface{}]
}

// ListDomainsResponse represents the response for listing domains
type ListDomainsResponse struct {
	paystackapi.Response[DomainList]
}

// DomainList represents the data field in ListDomainsResponse
type DomainList struct {
	DomainNames []string `json:"domainNames"`
}

// UnregisterDomainRequest represents the payload for unregistering a domain
type UnregisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

// UnregisterDomainResponse represents the response for unregistering a domain
type UnregisterDomainResponse struct {
	paystackapi.Response[interface{}]
}
