package paymentpages

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreatePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/page" {
			t.Errorf("Expected path /page, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Page created", "data": {"name": "Test Page", "slug": "test-page"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Create(context.Background(), &CreatePageRequest{Name: "Test Page"})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.Name != "Test Page" {
		t.Errorf("Expected name Test Page, got %s", resp.Data.Name)
	}
}

func TestListPages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/page" {
			t.Errorf("Expected path /page, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Pages retrieved", "data": [{"name": "Test Page"}]}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.List(context.Background())
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 page, got %d", len(resp.Data))
	}
}

func TestFetchPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/page/test-page" {
			t.Errorf("Expected path /page/test-page, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Page retrieved", "data": {"name": "Test Page", "slug": "test-page"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "test-page")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.Slug != "test-page" {
		t.Errorf("Expected slug test-page, got %s", resp.Data.Slug)
	}
}

func TestUpdatePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/page/test-page" {
			t.Errorf("Expected path /page/test-page, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Page updated", "data": {"name": "Updated Page", "slug": "test-page"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Update(context.Background(), "test-page", &UpdatePageRequest{Name: "Updated Page"})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.Name != "Updated Page" {
		t.Errorf("Expected name Updated Page, got %s", resp.Data.Name)
	}
}

func TestCheckSlugAvailability(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/page/check_slug_availability/test-page" {
			t.Errorf("Expected path /page/check_slug_availability/test-page, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Slug available", "data": true}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.CheckSlugAvailability(context.Background(), "test-page")
	if err != nil {
		t.Fatalf("CheckSlugAvailability failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if !resp.Data {
		t.Errorf("Expected data true, got false")
	}
}

func TestAddProducts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/page/123/product" {
			t.Errorf("Expected path /page/123/product, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Products added", "data": {"id": 123}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.AddProducts(context.Background(), 123, &AddProductsRequest{Products: []int{1, 2}})
	if err != nil {
		t.Fatalf("AddProducts failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.ID != 123 {
		t.Errorf("Expected ID 123, got %d", resp.Data.ID)
	}
}
