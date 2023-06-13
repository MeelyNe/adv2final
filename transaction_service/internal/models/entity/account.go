package entity

import "time"

type Account struct {
	ID            int       `json:"id"`
	User          User      `json:"user"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}
