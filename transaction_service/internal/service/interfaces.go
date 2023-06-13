package service

import (
	"context"
	"github.com/DiasOrazbaev/transaction_service/internal/models/entity"
)

type transactionService interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	GetTransactionByID(ctx context.Context, transactionID string) (*entity.Transaction, error)
	GetTransactionsByAccountID(ctx context.Context, accountID string) ([]*entity.Transaction, error)
}
