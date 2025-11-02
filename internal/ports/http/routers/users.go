package routers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goblinus/winapp/internal/domain"
)

type (
	UserService interface {
		GetUserByName(ctx context.Context, name string) (*domain.User, error)
	}

	UserRouter struct {
		users UserService
	}
)

func NewUserRouter(userService UserService) *UserRouter {
	return &UserRouter{
		users: userService,
	}
}

// users/:name
func (r UserRouter) GetByName(c *gin.Context) {
	userName := c.Params.ByName("name")
	user, err := r.users.GetUserByName(c.Request.Context(), userName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"user":   userName,
			"status": "not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":   user.Name,
		"status": "found",
	})

}
