package main

import (
	"github.com/DiasOrazbaev/transaction_service/internal/repository"
	"github.com/DiasOrazbaev/transaction_service/internal/service"
	"github.com/DiasOrazbaev/transaction_service/internal/transport/grpc"
	"github.com/DiasOrazbaev/transaction_service/pkg/database/postgres"
	"github.com/rs/zerolog"
	grpc2 "google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Info().Msg("Starting transaction service")

	db, err := postgres.NewConnection("postgres", "postgres", "localhost", "5432", "adv2final")
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return
	}
	defer db.Close()

	transactionRepo := repository.NewTransaction(db)
	transactionService := service.NewTransaction(transactionRepo)

	grpcServer := grpc2.NewServer()

	grpc.NewTransactionServiceServer(transactionService, &log).Register(grpcServer)

	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return
	}

	log.Info().Msg("starting server")

	if err = grpcServer.Serve(conn); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return
	}
}
