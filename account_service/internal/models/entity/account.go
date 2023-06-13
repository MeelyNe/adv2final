package entity

import "time"

type Account struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user"`
	AccountNumber int       `json:"account_number"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}
