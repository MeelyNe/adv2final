package main

import (
	"github.com/DiasOrazbaev/adv2final/user_service/internal/repository"
	"github.com/DiasOrazbaev/adv2final/user_service/internal/service"
	grpc2 "github.com/DiasOrazbaev/adv2final/user_service/internal/transport/grpc"
	"github.com/DiasOrazbaev/adv2final/user_service/pkg/database/postgres"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	db, err := postgres.NewConnection("postgres", "postgres", "localhost", "5432", "adv2final")
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUser(userRepo)

	srv := grpc.NewServer()

	grpc2.NewUser(userService, &log).Register(srv)

	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return
	}

	log.Info().Msg("starting server")

	if err = srv.Serve(conn); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return
	}
}
