package grpc

import (
	"context"
	"github.com/DiasOrazbaev/adv2final/user_service/internal/models/entity"
	userpb "github.com/DiasOrazbaev/adv2final/user_service/pkg/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type userService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, userID int) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type User struct {
	userService userService
	log         *zerolog.Logger
	userpb.UnimplementedUserServiceServer
}

func NewUser(userService userService, log *zerolog.Logger) *User {
	return &User{userService: userService, log: log}
}

func (u *User) Register(s *grpc.Server) {
	userpb.RegisterUserServiceServer(s, u)
}

func (u *User) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u.log.Info().Msg("CreateUser")
	return &userpb.CreateUserResponse{}, nil
}

func (u *User) GetUserByID(ctx context.Context, in *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	u.log.Info().Msg("GetUserByID")
	return &userpb.GetUserByIDResponse{}, nil
}

func (u *User) GetUserByUsername(ctx context.Context, in *userpb.GetUserByUsernameRequest) (*userpb.GetUserByUsernameResponse, error) {
	u.log.Info().Msg("GetUserByUsername")
	return &userpb.GetUserByUsernameResponse{}, nil
}

func (u *User) GetUserByEmail(ctx context.Context, in *userpb.GetUserByEmailRequest) (*userpb.GetUserByEmailResponse, error) {
	u.log.Info().Msg("GetUserByEmail")
	return &userpb.GetUserByEmailResponse{}, nil
}
