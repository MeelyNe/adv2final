package repository

import (
	"context"
	"database/sql"
	"github.com/DiasOrazbaev/transaction_service/internal/models/entity"
)

type Transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{db: db}
}

func (t *Transaction) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	query := `INSERT INTO transactions (account_id, type, amount, timestamp) VALUES ($1, $2, $3, $4)`
	_, err := t.db.ExecContext(ctx, query, transaction.ID, transaction.Type, transaction.Amount, transaction.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) GetTransactionByID(ctx context.Context, transactionID string) (*entity.Transaction, error) {
	query := `SELECT id, type, amount, timestamp FROM transactions WHERE id = $1`
	row := t.db.QueryRowContext(ctx, query, transactionID)
	transaction := &entity.Transaction{}
	err := row.Scan(&transaction.ID, &transaction.Type, &transaction.Amount, &transaction.Timestamp)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) GetTransactionsByAccountID(ctx context.Context, accountID string) ([]*entity.Transaction, error) {
	query := `SELECT id, type, amount, timestamp FROM transactions WHERE account_id = $1`
	rows, err := t.db.QueryContext(ctx, query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactions := make([]*entity.Transaction, 0)
	for rows.Next() {
		transaction := &entity.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.Type, &transaction.Amount, &transaction.Timestamp)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}
