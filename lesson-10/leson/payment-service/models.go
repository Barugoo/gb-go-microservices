package main

import "time"

type TransactionStatus string

const (
	TransactionStatusNew     TransactionStatus = "new"
	TransactionStatusDone    TransactionStatus = "done"
	TransactionStatusFailure TransactionStatus = "failure"
)

type Transaction struct {
	ID         int64             `json:"id"`
	UserID     int64             `json:"user_id"`
	Amount     int64             `json:"amount"`
	ExternalID string            `json:"external_id"`
	Status     TransactionStatus `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type MovieOwnership struct {
	ID        int64
	MovieID   int64
	UserID    int64
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
