package http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/goblinus/winapp/internal/domain"
	"github.com/goblinus/winapp/internal/ports/http/routers"
)

type (
	UserProvider interface {
		GetUserByName(ctx context.Context, name string) (*domain.User, error)
	}

	TaskProvider interface {
		DeleteTask(ctx context.Context, UUID string) (*domain.Task, error)
		CreateTask(ctx context.Context, taskInput *domain.TaskInput) (*domain.Task, error)
		GetTasksAll(ctx context.Context) ([]*domain.Task, error)
		GetTaskByUUID(ctx context.Context, UUID string) (*domain.Task, error)
	}
)

func RegisterRoutes(users UserProvider, tasks TaskProvider, apiRedis *redis.Client) *gin.Engine {
	r := gin.Default()

	{
		router := routers.NewPingRouter(apiRedis)
		r.GET("/ping", router.Ping)
	}

	v1Routers := r.Group("/v1")
	registerV1Routers(v1Routers, users, tasks)

	return r
}
