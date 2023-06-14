package main

import (
	"github.com/DiasOrazbaev/adv2final/account_service/internal/config"
	"github.com/DiasOrazbaev/adv2final/account_service/internal/repository"
	"github.com/DiasOrazbaev/adv2final/account_service/internal/service"
	grpc2 "github.com/DiasOrazbaev/adv2final/account_service/internal/transport/grpc"
	"github.com/DiasOrazbaev/adv2final/account_service/pkg/database/postgres"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("failed to load env")
		return
	}
	cfg := config.GetConfig()
	db, err := postgres.NewConnection(
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return
	}
	defer db.Close()

	accoutRepo := repository.NewAccount(db)
	accountServ := service.NewAccount(accoutRepo)

	grpcs := grpc.NewServer()

	grpc2.NewAccount(accountServ, &log).Register(grpcs)

	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return
	}

	log.Info().Msg("starting server")

	if err = grpcs.Serve(conn); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return
	}
}
