package grpc

import (
	"context"
	"github.com/DiasOrazbaev/adv2final/user_service/internal/models/entity"
	userpb "github.com/DiasOrazbaev/adv2final/user_service/pkg/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	cr := &entity.User{
		Username:  in.User.Username,
		Email:     in.User.Email,
		Password:  in.User.Password,
		FullName:  in.User.FullName,
		CreatedAt: in.User.CreatedAt.AsTime(),
	}

	err := u.userService.CreateUser(ctx, cr)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		UserId: int32(cr.ID),
	}, nil
}

func (u *User) GetUserByID(ctx context.Context, in *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	u.log.Info().Msg("GetUserByID")

	user, err := u.userService.GetUserByID(ctx, int(in.UserId))
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserByIDResponse{
		User: &userpb.User{
			Id:        int32(user.ID),
			Username:  user.Username,
			Password:  user.Password,
			Email:     user.Email,
			FullName:  user.FullName,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}

func (u *User) GetUserByUsername(ctx context.Context, in *userpb.GetUserByUsernameRequest) (*userpb.GetUserByUsernameResponse, error) {
	u.log.Info().Msg("GetUserByUsername")

	user, err := u.userService.GetUserByUsername(ctx, in.Username)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserByUsernameResponse{
		User: &userpb.User{
			Id:        int32(user.ID),
			Username:  user.Username,
			Password:  user.Password,
			Email:     user.Email,
			FullName:  user.FullName,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}

func (u *User) GetUserByEmail(ctx context.Context, in *userpb.GetUserByEmailRequest) (*userpb.GetUserByEmailResponse, error) {
	u.log.Info().Msg("GetUserByEmail")
	return nil, nil
}
