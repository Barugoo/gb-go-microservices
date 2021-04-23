package main

import "context"

type Repository interface {
	CreateTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error)
	UpdateTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error)
	GetTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error)

	CreateMovieOwnershop(ctx context.Context, movieOwnershop *MovieOwnership) (*MovieOwnership, error)

	GetUserBalance(ctx context.Context, userID int64) (int64, error)
	// select user_id, sum(amount) from transaction group by user_id where status = 'done' and user_id = '<user_id>';
}
