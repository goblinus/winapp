package inmemory

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/goblinus/winapp/internal/domain"
)

type (
	TaskRepository struct {
		data map[string]*domain.Task
	}
)

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		data: make(map[string]*domain.Task),
	}
}

func (r *TaskRepository) Delete(ctx context.Context, UUID string) (*domain.Task, error) {
	for taskUUID := range r.data {
		if taskUUID == UUID {
			result := r.data[taskUUID]
			delete(r.data, taskUUID)

			return result, nil
		}
	}

	return nil, ErrTaskNotFound
}

func (r *TaskRepository) Create(ctx context.Context, taskInput *domain.TaskInput) (*domain.Task, error) {
	result := &domain.Task{
		UUID:        uuid.New().String(),
		Name:        taskInput.Name,
		Description: taskInput.Description,
		Timestamp:   time.Now(),
	}

	r.data[result.UUID] = result

	return r.data[result.UUID], nil
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]*domain.Task, error) {
	var result []*domain.Task
	for taskID := range r.data {
		result = append(result, r.data[taskID])
	}

	return result, nil
}

func (r *TaskRepository) GetByUUID(ctx context.Context, UUID string) (*domain.Task, error) {
	for taskUUID := range r.data {
		if taskUUID == UUID {
			result := r.data[UUID]

			return result, nil
		}
	}

	return nil, ErrTaskNotFound
}
