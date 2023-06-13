package repository

import (
	"context"
	"database/sql"
	"github.com/DiasOrazbaev/adv2final/account_service/internal/models/entity"
)

//	CREATE TABLE accounts
//	(
//		id             SERIAL PRIMARY KEY,
//		user_id        INTEGER REFERENCES users (id),
//		account_number VARCHAR(20) UNIQUE,
//		balance        DECIMAL(10, 2),
//		created_at     TIMESTAMP
//	);

type Account struct {
	db *sql.DB
}

func (a *Account) CreateAccount(ctx context.Context, account *entity.Account) error {
	query := `INSERT INTO accounts (user_id, account_number, balance, created_at) VALUES ($1, $2, $3, $4)`
	_, err := a.db.ExecContext(ctx, query, account.UserID, account.AccountNumber, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error) {
	query := `SELECT * FROM accounts WHERE id = $1`
	row := a.db.QueryRowContext(ctx, query, accountID)
	account := &entity.Account{}
	err := row.Scan(&account.ID, &account.UserID, &account.AccountNumber, &account.Balance, &account.CreatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) GetAccountsByUserID(ctx context.Context, userID string) ([]*entity.Account, error) {
	query := `SELECT * FROM accounts WHERE user_id = $1`
	rows, err := a.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]*entity.Account, 0)
	for rows.Next() {
		account := &entity.Account{}
		err := rows.Scan(&account.ID, &account.UserID, &account.AccountNumber, &account.Balance, &account.CreatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}
