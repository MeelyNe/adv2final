package grpc

import (
	"context"
	"github.com/DiasOrazbaev/transaction_service/internal/models/entity"
	accountpb "github.com/DiasOrazbaev/transaction_service/pkg/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type transactionService interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	GetTransactionByID(ctx context.Context, transactionID string) (*entity.Transaction, error)
	GetTransactionsByAccountID(ctx context.Context, accountID string) ([]*entity.Transaction, error)
}

type TransactionServiceServer struct {
	transactionService transactionService
	log                *zerolog.Logger
	accountpb.UnimplementedAccountServiceServer
}

func NewTransactionServiceServer(transactionService transactionService, log *zerolog.Logger) *TransactionServiceServer {
	return &TransactionServiceServer{transactionService: transactionService, log: log}
}

func (t *TransactionServiceServer) Register(s *grpc.Server) {
	accountpb.RegisterAccountServiceServer(s, t)
}

func (t *TransactionServiceServer) CreateUser(context.Context, *accountpb.CreateUserRequest) (*accountpb.CreateUserResponse, error) {
	t.log.Info().Msg("CreateUser")
	return &accountpb.CreateUserResponse{}, nil
}

func (t *TransactionServiceServer) GetUserByID(context.Context, *accountpb.GetUserByIDRequest) (*accountpb.GetUserByIDResponse, error) {
	t.log.Info().Msg("GetUserByID")
	return &accountpb.GetUserByIDResponse{}, nil
}

func (t *TransactionServiceServer) GetAccountByID(context.Context, *accountpb.GetAccountByIDRequest) (*accountpb.GetAccountByIDResponse, error) {
	t.log.Info().Msg("GetAccountByID")
	return &accountpb.GetAccountByIDResponse{}, nil
}

func (t *TransactionServiceServer) GetAccountsByUserID(context.Context, *accountpb.GetAccountsByUserIDRequest) (*accountpb.GetAccountsByUserIDResponse, error) {
	t.log.Info().Msg("GetAccountsByUserID")
	return &accountpb.GetAccountsByUserIDResponse{}, nil
}
