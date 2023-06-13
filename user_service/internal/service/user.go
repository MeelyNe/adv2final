package service

import (
	"context"
	"fmt"
	"github.com/DiasOrazbaev/adv2final/user_service/internal/models/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const timeout = 5 * time.Second

type userRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, userID int) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type User struct {
	repo userRepository
}

func NewUser(repo userRepository) *User {
	return &User{repo: repo}
}

func (u *User) CreateUser(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err = u.repo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (u *User) GetUserByID(ctx context.Context, userID int) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	user, err := u.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}

func (u *User) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	user, err := u.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return user, nil
}
