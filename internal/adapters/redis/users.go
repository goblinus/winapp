package redis

import (
	"context"

	"github.com/goblinus/winapp/internal/domain"
	"github.com/redis/go-redis/v9"
)

type (
	UserRepository struct {
		api *redis.Client
	}
)

func NewUserRepository(addr string) *UserRepository {
	apiClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", //no password
		DB:       0,  // use default DB
	})

	return &UserRepository{
		api: apiClient,
	}
}

func (r *UserRepository) GetByName(ctx context.Context, name string) (*domain.User, error) {
	return nil, nil
}
