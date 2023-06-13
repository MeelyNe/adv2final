package repository

import (
	"context"
	"database/sql"
	"github.com/DiasOrazbaev/adv2final/user_service/internal/models/entity"
)

type (
	User struct {
		db *sql.DB
	}
)

func NewUserRepository(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) Create(ctx context.Context, user *entity.User) error {
	_, err := u.db.ExecContext(ctx, "INSERT INTO users (username, password, email, full_name, created_at) VALUES ($1, $2, $3, $4, now())",
		user.Username, user.Password, user.Email, user.FullName)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetByID(ctx context.Context, userID int) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.QueryRowContext(ctx, "SELECT id, username, password, email, full_name, created_at FROM users WHERE id=$1", userID).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FullName, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.QueryRowContext(ctx, "SELECT id, username, password, email, full_name, created_at FROM users WHERE username=$1", username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FullName, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.QueryRowContext(ctx, "SELECT id, username, password, email, full_name, created_at FROM users WHERE email=$1", email).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FullName, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
