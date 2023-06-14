package grpc

import (
	"context"
	"github.com/DiasOrazbaev/transaction_service/internal/models/entity"
	transactionpb "github.com/DiasOrazbaev/transaction_service/pkg/proto"
	"github.com/rs/zerolog"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type transactionService interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	GetTransactionByID(ctx context.Context, transactionID int) (*entity.Transaction, error)
	GetTransactionsByAccountID(ctx context.Context, accountID int) ([]*entity.Transaction, error)
}

type TransactionServiceServer struct {
	transactionService transactionService
	log                *zerolog.Logger
	transactionpb.UnimplementedTransactionServiceServer
}

func NewTransactionServiceServer(transactionService transactionService, log *zerolog.Logger) *TransactionServiceServer {
	return &TransactionServiceServer{transactionService: transactionService, log: log}
}

func (t *TransactionServiceServer) Register(grpcServer *grpc2.Server) {
	transactionpb.RegisterTransactionServiceServer(grpcServer, t)
}

func (t *TransactionServiceServer) CreateTransaction(ctx context.Context, req *transactionpb.CreateTransactionRequest) (*transactionpb.CreateTransactionResponse, error) {
	t.log.Info().Msg("CreateTransaction")
	transaction := &entity.Transaction{
		AccountID: int(req.AccountId),
		Amount:    req.Amount,
	}

	transaction, err := t.transactionService.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, status.Error(12, err.Error())
	}

	return &transactionpb.CreateTransactionResponse{
		TransactionId: int64(transaction.ID),
	}, nil
}
func (t *TransactionServiceServer) GetTransactionByID(ctx context.Context, req *transactionpb.GetTransactionByIDRequest) (*transactionpb.GetTransactionByIDResponse, error) {
	t.log.Info().Msg("GetTransactionByID")
	transaction, err := t.transactionService.GetTransactionByID(ctx, int(req.TransactionId))
	if err != nil {
		return nil, status.Error(12, err.Error())
	}

	return &transactionpb.GetTransactionByIDResponse{
		Transaction: &transactionpb.Transaction{
			AccountId: int64(transaction.AccountID),
			Amount:    transaction.Amount,
			Timestamp: timestamppb.New(transaction.Timestamp),
		},
	}, nil
}
func (t *TransactionServiceServer) GetTransactionsByAccountID(ctx context.Context, req *transactionpb.GetTransactionsByAccountIDRequest) (*transactionpb.GetTransactionsByAccountIDResponse, error) {
	t.log.Info().Msg("GetTransactionsByAccountID")

	transactions, err := t.transactionService.GetTransactionsByAccountID(ctx, int(req.AccountId))
	if err != nil {
		return nil, status.Error(12, err.Error())
	}

	var transactionsPb []*transactionpb.Transaction
	for _, transaction := range transactions {
		transactionsPb = append(transactionsPb, &transactionpb.Transaction{
			AccountId: int64(transaction.AccountID),
			Amount:    transaction.Amount,
			Timestamp: timestamppb.New(transaction.Timestamp),
		})
	}

	return &transactionpb.GetTransactionsByAccountIDResponse{
		Transactions: transactionsPb,
	}, nil
}
