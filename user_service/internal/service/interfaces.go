package service

import (
	"context"
	"github.com/DiasOrazbaev/adv2final/user_service/internal/models/entity"
)

type userService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, userID int) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}
