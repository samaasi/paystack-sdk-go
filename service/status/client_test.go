package status

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		response := `{
			"page": {
				"id": "y3b19k59530l",
				"name": "Paystack",
				"url": "https://status.paystack.com",
				"time_zone": "Africa/Lagos",
				"updated_at": "2023-10-27T10:00:00.000+01:00"
			},
			"status": {
				"indicator": "none",
				"description": "All Systems Operational"
			},
			"components": [
				{
					"id": "0l2p9nhq71w8",
					"name": "API",
					"status": "operational",
					"created_at": "2016-04-13T14:42:04.664+01:00",
					"updated_at": "2023-10-27T10:00:00.000+01:00",
					"position": 1,
					"description": "API requests",
					"showcase": false,
					"start_date": null,
					"group_id": null,
					"only_show_if_degraded": false
				}
			]
		}`

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, response)
	}))
	defer ts.Close()

	// Initialize client and override URL with mock server
	client := NewClient(nil).WithURL(ts.URL)

	summary, err := client.Fetch(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if summary.Page.Name != "Paystack" {
		t.Errorf("Expected page name Paystack, got %s", summary.Page.Name)
	}

	if summary.Status.Description != "All Systems Operational" {
		t.Errorf("Expected status description 'All Systems Operational', got %s", summary.Status.Description)
	}

	if len(summary.Components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(summary.Components))
	}

	if summary.Components[0].Name != "API" {
		t.Errorf("Expected component name API, got %s", summary.Components[0].Name)
	}
}
