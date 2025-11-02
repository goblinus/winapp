package http

import (
	"github.com/gin-gonic/gin"

	"github.com/goblinus/winapp/internal/ports/http/routers"
)

func registerV1Routers(
	v1Routers *gin.RouterGroup,
	users UserProvider,
	tasks TaskProvider,
) {
	{
		g := v1Routers.Group("/users")
		router := routers.NewUserRouter(users)

		g.GET("/:name", router.GetByName)
	}

	{
		g := v1Routers.Group("/tasks")
		router := routers.NewTaskRoutes(tasks)

		g.GET("/", router.GetTasksAll)
		g.GET("/:uuid", router.GetTaskByID)
		g.POST("/", router.CreateTask)
		g.DELETE("/:uuid", router.DeleteTask)
	}
}
