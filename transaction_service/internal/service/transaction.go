package service

import (
	"context"
	"github.com/DiasOrazbaev/transaction_service/internal/models/entity"
)

type transactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
	GetTransactionByID(ctx context.Context, transactionID int) (*entity.Transaction, error)
	GetTransactionsByAccountID(ctx context.Context, accountID int) ([]*entity.Transaction, error)
}

type Transaction struct {
	transactionRepository transactionRepository
}

func NewTransaction(transactionRepository transactionRepository) *Transaction {
	return &Transaction{transactionRepository: transactionRepository}
}

func (t *Transaction) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	err := t.transactionRepository.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) GetTransactionByID(ctx context.Context, transactionID int) (*entity.Transaction, error) {
	transaction, err := t.transactionRepository.GetTransactionByID(ctx, transactionID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) GetTransactionsByAccountID(ctx context.Context, accountID int) ([]*entity.Transaction, error) {
	transactions, err := t.transactionRepository.GetTransactionsByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
