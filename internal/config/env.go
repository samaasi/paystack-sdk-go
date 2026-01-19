package config

import "os"

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment the value (which may be empty) is returned.
// Otherwise, the fallback value is returned.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetPaystackSecretKey retrieves the Paystack secret key from environment variables.
func GetPaystackSecretKey() string {
	return os.Getenv("PAYSTACK_SECRET_KEY")
}
