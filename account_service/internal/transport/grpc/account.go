package grpc

import (
	"context"
	"github.com/DiasOrazbaev/adv2final/account_service/internal/models/entity"
	accountpb "github.com/DiasOrazbaev/adv2final/account_service/pkg/proto"
	"github.com/rs/zerolog"
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

func (a *Account) GetAccountByID(context.Context, *accountpb.GetAccountByIDRequest) (*accountpb.GetAccountByIDResponse, error) {
	a.log.Info().Msg("GetAccountByID")
	return &accountpb.GetAccountByIDResponse{}, nil
}

func (a *Account) GetAccountsByUserID(context.Context, *accountpb.GetAccountsByUserIDRequest) (*accountpb.GetAccountsByUserIDResponse, error) {
	a.log.Info().Msg("GetAccountsByUserID")
	return &accountpb.GetAccountsByUserIDResponse{}, nil
}
