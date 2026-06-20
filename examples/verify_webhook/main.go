// Example: verify an incoming Paystack webhook and dispatch on event type.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samaasi/paystack-sdk-go/webhook"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")

	// Optionally restrict to known Paystack IPs.
	if !webhook.IsFromPaystackIP(r) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	var event webhook.Event
	if err := webhook.Parse(r, secretKey, &event); err != nil {
		http.Error(w, "invalid webhook", http.StatusBadRequest)
		return
	}

	switch event.Event {
	case webhook.EventChargeSuccess:
		var data webhook.ChargeSuccessEvent
		if err := event.UnmarshalData(&data); err != nil {
			log.Printf("failed to parse charge.success: %v", err)
			return
		}
		fmt.Printf("Payment received: %s — NGN %.2f\n", data.Reference, float64(data.Amount)/100)

	case webhook.EventTransferSuccess:
		var data webhook.TransferSuccessEvent
		if err := event.UnmarshalData(&data); err != nil {
			log.Printf("failed to parse transfer.success: %v", err)
			return
		}
		fmt.Printf("Transfer successful: %s\n", data.TransferCode)

	case webhook.EventSubscriptionCreate:
		var data webhook.SubscriptionEvent
		if err := event.UnmarshalData(&data); err != nil {
			log.Printf("failed to parse subscription.create: %v", err)
			return
		}
		fmt.Printf("New subscription: %s\n", data.SubscriptionCode)

	default:
		fmt.Printf("Unhandled event: %s\n", event.Event)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
