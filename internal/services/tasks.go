package services

import (
	"context"

	"github.com/goblinus/winapp/internal/domain"
)

type (
	TaskStorage interface {
		Delete(ctx context.Context, UUID string) (*domain.Task, error)
		Create(ctx context.Context, task *domain.TaskInput) (*domain.Task, error)
		GetAll(ctx context.Context) ([]*domain.Task, error)
		GetByUUID(ctx context.Context, UUID string) (*domain.Task, error)
	}

	TaskService struct {
		tasks TaskStorage
	}
)

func NewTaskService(tasksRepo TaskStorage) *TaskService {
	return &TaskService{
		tasks: tasksRepo,
	}
}

func (s TaskService) DeleteTask(ctx context.Context, id string) (*domain.Task, error) {
	if len(id) == 0 {
		return nil, ErrWrongTaskID
	}

	return s.tasks.Delete(ctx, id)
}

func (s TaskService) CreateTask(ctx context.Context, taskInput *domain.TaskInput) (*domain.Task, error) {
	return s.tasks.Create(ctx, taskInput)
}

func (s TaskService) GetTasksAll(ctx context.Context) ([]*domain.Task, error) {
	return s.tasks.GetAll(ctx)
}

func (s TaskService) GetTaskByUUID(ctx context.Context, UUID string) (*domain.Task, error) {
	if len(UUID) == 0 {
		return nil, ErrWrongTaskID
	}

	return s.tasks.GetByUUID(ctx, UUID)
}
