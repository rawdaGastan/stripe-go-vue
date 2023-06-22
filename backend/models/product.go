// Package models for database models
package models

// Currency for currency types
type Currency string

const (
	USD Currency = "USD"
	EG  Currency = "EG"
)

// Product struct for products.
type Product struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Price    int64    `json:"price"`
	PriceID  string   `json:"price_id"`
	Currency Currency `json:"currency"`
	Amount   int64    `json:"amount"`
}
