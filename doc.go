// Package paystacksdkgo provides a comprehensive and robust Go client for the Paystack API.
//
// This SDK covers all major Paystack services including Transactions, Customers, Plans,
// Subscriptions, Transfers, and more. It is designed to be idiomatic, context-aware,
// and easy to use.
//
// # Initialization
//
//	import (
//		"os"
//		paystack "github.com/samaasi/paystack-sdk-go"
//	)
//
//	func main() {
//		apiKey := os.Getenv("PAYSTACK_SECRET_KEY")
//		client := paystack.NewClient(apiKey)
//		// Use client to interact with Paystack API...
//	}
//
// # Services
//
// The client provides access to various services via fields:
//
//	client.Transactions.Initialize(...)
//	client.Customers.Create(...)
//
// See individual service documentation for details.
package paystacksdkgo
