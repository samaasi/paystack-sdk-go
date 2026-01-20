package products

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samaasi/paystack-sdk-go/internal/backend"
)

func TestCreateProduct(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.URL.Path != "/product" {
			t.Errorf("Expected path /product, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Product created", "data": {"name": "Test Product", "product_code": "PROD_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Create(context.Background(), &CreateProductRequest{Name: "Test Product", Price: 5000, Currency: "NGN"})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.ProductCode != "PROD_123" {
		t.Errorf("Expected product code PROD_123, got %s", resp.Data.ProductCode)
	}
}

func TestListProducts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/product" {
			t.Errorf("Expected path /product, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Products retrieved", "data": [{"product_code": "PROD_123"}]}`)
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
		t.Errorf("Expected 1 product, got %d", len(resp.Data))
	}
}

func TestFetchProduct(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Path != "/product/PROD_123" {
			t.Errorf("Expected path /product/PROD_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Product retrieved", "data": {"product_code": "PROD_123"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Fetch(context.Background(), "PROD_123")
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.ProductCode != "PROD_123" {
		t.Errorf("Expected product code PROD_123, got %s", resp.Data.ProductCode)
	}
}

func TestUpdateProduct(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/product/PROD_123" {
			t.Errorf("Expected path /product/PROD_123, got %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"status": true, "message": "Product updated", "data": {"product_code": "PROD_123", "name": "Updated Product"}}`)
	}))
	defer server.Close()

	client := NewClient(backend.NewClient("secret", backend.WithBaseURL(server.URL)))

	resp, err := client.Update(context.Background(), "PROD_123", &UpdateProductRequest{Name: "Updated Product"})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status true, got false")
	}
	if resp.Data.Name != "Updated Product" {
		t.Errorf("Expected name Updated Product, got %s", resp.Data.Name)
	}
}
