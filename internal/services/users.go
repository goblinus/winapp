package services

import (
	"context"

	"github.com/goblinus/winapp/internal/domain"
)

type (
	UserStorage interface {
		GetByName(ctx context.Context, name string) (*domain.User, error)
	}

	UserService struct {
		users UserStorage
	}
)

func NewUserService(usersDB UserStorage) *UserService {
	return &UserService{
		users: usersDB,
	}
}

func (s UserService) GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	return s.users.GetByName(ctx, name)
}
