package applepay

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestRegisterDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/apple-pay/domain" {
			t.Errorf("Expected path /apple-pay/domain, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Domain registered", "data": null}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.RegisterDomain(context.Background(), &RegisterDomainRequest{DomainName: "example.com"})
	if err != nil {
		t.Fatalf("RegisterDomain failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestListDomains(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/apple-pay/domain" {
			t.Errorf("Expected path /apple-pay/domain, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Domains retrieved", "data": {"domainNames": ["example.com", "test.com"]}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.ListDomains(context.Background())
	if err != nil {
		t.Fatalf("ListDomains failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}

	if len(resp.Data.DomainNames) != 2 {
		t.Errorf("Expected 2 domains, got %d", len(resp.Data.DomainNames))
	}
	if resp.Data.DomainNames[0] != "example.com" {
		t.Errorf("Expected domain example.com, got %s", resp.Data.DomainNames[0])
	}
}

func TestUnregisterDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected method DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/apple-pay/domain" {
			t.Errorf("Expected path /apple-pay/domain, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Domain unregistered", "data": null}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.UnregisterDomain(context.Background(), &UnregisterDomainRequest{DomainName: "example.com"})
	if err != nil {
		t.Fatalf("UnregisterDomain failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
}
