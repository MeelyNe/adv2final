package entity

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
