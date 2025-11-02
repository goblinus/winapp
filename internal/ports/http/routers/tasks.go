package routers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goblinus/winapp/internal/domain"
)

type (
	TaskService interface {
		DeleteTask(ctx context.Context, UUID string) (*domain.Task, error)
		CreateTask(ctx context.Context, taskInput *domain.TaskInput) (*domain.Task, error)
		GetTasksAll(ctx context.Context) ([]*domain.Task, error)
		GetTaskByUUID(ctx context.Context, UUID string) (*domain.Task, error)
	}

	TasksRoutes struct {
		tasks TaskService
	}
)

func NewTaskRoutes(taskService TaskService) *TasksRoutes {
	return &TasksRoutes{
		tasks: taskService,
	}
}

// /tasks/:id [DELETE]
func (s TasksRoutes) DeleteTask(c *gin.Context) {
	taskUUID := c.Params.ByName("uuid")
	task, err := s.tasks.DeleteTask(c.Request.Context(), taskUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"stutus": "OK",
		"msg": fmt.Sprintf(
			"%s [%s] deleted successfully",
			task.Name,
			task.UUID,
		),
	})
}

// /tasks [POST]
func (s TasksRoutes) CreateTask(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    fmt.Errorf("cannot read request: %w", err),
		})
	}

	defer c.Request.Body.Close()

	taskInput := new(domain.TaskInput)
	if err := taskInput.Unmarshal(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    fmt.Errorf("cannot unumarshal task: %w", err),
		})
	}

	result, err := s.tasks.CreateTask(c.Request.Context(), taskInput)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"status": "error",
			"msg":    fmt.Errorf("cannot unumarshal task: %w", err),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   result,
	})
}

// /tasks [GET]
func (s TasksRoutes) GetTasksAll(c *gin.Context) {
	tasks, err := s.tasks.GetTasksAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   tasks,
	})
}

// /tasks/:id [GET]
func (s TasksRoutes) GetTaskByID(c *gin.Context) {
	taskUUID := c.Params.ByName("uuid")
	task, err := s.tasks.GetTaskByUUID(c.Request.Context(), taskUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   task,
	})
}
