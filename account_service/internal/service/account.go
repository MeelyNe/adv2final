package service

import (
	"context"
	"github.com/DiasOrazbaev/adv2final/account_service/internal/models/entity"
	"time"
)

type accountRepository interface {
	CreateAccount(ctx context.Context, account *entity.Account) error
	GetAccountByID(ctx context.Context, accountID int) (*entity.Account, error)
	GetAccountsByUserID(ctx context.Context, userID int) ([]*entity.Account, error)
}

type Account struct {
	repo accountRepository
}

func (a *Account) CreateAccount(ctx context.Context, userID int, accountNumber int) (*entity.Account, error) {
	acc := &entity.Account{
		UserID:        userID,
		AccountNumber: accountNumber,
		Balance:       0,
		CreatedAt:     time.Now(),
	}
	err := a.repo.CreateAccount(ctx, acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (a *Account) GetAccountByID(ctx context.Context, accountID int) (*entity.Account, error) {
	acc, err := a.repo.GetAccountByID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (a *Account) GetAccountsByUserID(ctx context.Context, userID int) ([]*entity.Account, error) {
	accs, err := a.repo.GetAccountsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return accs, nil
}

func NewAccount(repo accountRepository) *Account {
	return &Account{repo: repo}
}
