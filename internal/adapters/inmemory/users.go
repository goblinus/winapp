package inmemory

import (
	"context"

	"github.com/goblinus/winapp/internal/domain"
	"github.com/google/uuid"
)

type (
	UserRepository struct {
		data map[string]*domain.User
	}
)

func NewUserRepository() *UserRepository {
	data := make(map[string]*domain.User)

	{
		newUUID := uuid.New()
		data[newUUID.String()] = &domain.User{
			UUID: newUUID.String(),
			Name: "John",
		}
	}

	{
		newUUID := uuid.New()
		data[newUUID.String()] = &domain.User{
			UUID: newUUID.String(),
			Name: "David",
		}
	}

	return &UserRepository{
		data: data,
	}
}

func (r *UserRepository) GetByName(ctx context.Context, name string) (*domain.User, error) {
	for userUUID := range r.data {
		if r.data[userUUID].Name == name {
			return r.data[userUUID], nil
		}
	}

	return nil, ErrUserNotFound
}
