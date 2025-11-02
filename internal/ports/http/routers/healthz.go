package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type PingRouter struct {
	apiRedis *redis.Client
}

func NewPingRouter(apiRedis *redis.Client) *PingRouter {
	return &PingRouter{
		apiRedis: apiRedis,
	}
}

// /ping
func (r PingRouter) Ping(c *gin.Context) {
	redisResult, err := r.apiRedis.Ping(c.Request.Context()).Result()
	if err != nil {
		redisResult = "DISPONG"
	}

	c.JSON(http.StatusOK, gin.H{
		"http":  "PONG",
		"redis": redisResult,
	})
}
