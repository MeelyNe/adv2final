package grpc

import (
	"context"
	"github.com/DiasOrazbaev/adv2final/account_service/internal/models/entity"
	accountpb "github.com/DiasOrazbaev/adv2final/account_service/pkg/proto"
	"github.com/rs/zerolog"
	status "google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type accountService interface {
	CreateAccount(ctx context.Context, userID int, accountNumber int) (*entity.Account, error)
	GetAccountByID(ctx context.Context, accountID int) (*entity.Account, error)
	GetAccountsByUserID(ctx context.Context, userID int) ([]*entity.Account, error)
}

type Account struct {
	accountService accountService
	log            *zerolog.Logger
	accountpb.UnimplementedAccountServiceServer
}

func (a *Account) Register(grpc *grpc.Server) {
	accountpb.RegisterAccountServiceServer(grpc, a)
}

func (a *Account) GetAccountByID(ctx context.Context, req *accountpb.GetAccountByIDRequest) (*accountpb.GetAccountByIDResponse, error) {
	a.log.Info().Msg("GetAccountByID")
	acc, err := a.accountService.GetAccountByID(ctx, int(req.AccountId))
	if err != nil {
		a.log.Error().Err(err).Msg("GetAccountByID")
		return nil, status.Error(12, err.Error())
	}
	return &accountpb.GetAccountByIDResponse{
		Account: &accountpb.Account{
			Id:            int64(acc.ID),
			AccountNumber: int64(acc.AccountNumber),
			UserId:        int64(acc.UserID),
		},
	}, nil
}

func (a *Account) GetAccountsByUserID(ctx context.Context, req *accountpb.GetAccountsByUserIDRequest) (*accountpb.GetAccountsByUserIDResponse, error) {
	a.log.Info().Msg("GetAccountsByUserID")
	accs, err := a.accountService.GetAccountsByUserID(ctx, int(req.UserId))
	if err != nil {
		a.log.Error().Err(err).Msg("GetAccountsByUserID")
		return nil, status.Error(12, err.Error())
	}
	accounts := make([]*accountpb.Account, 0, len(accs))
	for _, acc := range accs {
		accounts = append(accounts, &accountpb.Account{
			Id:            int64(acc.ID),
			AccountNumber: int64(acc.AccountNumber),
			UserId:        int64(acc.UserID),
		})
	}
	return &accountpb.GetAccountsByUserIDResponse{
		Accounts: accounts,
	}, nil
}
