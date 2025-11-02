package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/goblinus/winapp/internal/domain"
)

type (
	TaskRepository struct {
		api *redis.Client
	}
)

func NewTaskRepository(apiRedis *redis.Client) *TaskRepository {
	return &TaskRepository{
		api: apiRedis,
	}
}

func (r *TaskRepository) Delete(ctx context.Context, UUID string) (*domain.Task, error) {
	task, err := r.GetByUUID(ctx, UUID)
	if err != nil {
		return nil, err
	}

	if err := r.api.Unlink(ctx, fmt.Sprintf("tasks:%s", UUID)).Err(); err != nil {
		return nil, fmt.Errorf("unlink task with uuid=%s: %w", UUID, err)
	}

	if err := r.api.ZRem(ctx, "tasks", UUID).Err(); err != nil {
		return nil, fmt.Errorf("clear memory for task with uuid=%s: %w", UUID, err)
	}

	return task, nil
}

func (r *TaskRepository) Create(ctx context.Context, taskInput *domain.TaskInput) (*domain.Task, error) {
	newTask := &domain.Task{
		UUID:        uuid.New().String(),
		Name:        taskInput.Name,
		Description: taskInput.Description,
		Timestamp:   time.Now(),
	}

	taskHash := fmt.Sprintf("task:%s", newTask.UUID)

	hset := r.api.HSet(ctx, taskHash,
		"uuid", newTask.UUID,
		"name", newTask.Name,
		"description", newTask.Description,
		"timestamp", newTask.Timestamp)
	if hset.Err() != nil {
		return nil, fmt.Errorf("create hash set for new task: %w", hset.Err())
	}

	z := redis.Z{
		Score:  float64(newTask.Timestamp.Unix()),
		Member: newTask.UUID,
	}

	zadd := r.api.ZAdd(ctx, "task", z)
	if zadd.Err() != nil {
		return nil, fmt.Errorf("zadd for task: %w", zadd.Err())
	}

	return newTask, nil
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]*domain.Task, error) {
	zrange := r.api.ZRange(ctx, "task", 0, -1)
	if zrange.Err() != nil {
		return nil, fmt.Errorf("get tasks: %w", zrange.Err())
	}

	taskUUIDs, err := zrange.Result()
	if err != nil {
		return nil, fmt.Errorf("resulting tasks identifiers: %w", err)
	}

	result := make([]*domain.Task, 0, len(taskUUIDs))

	for _, taskUUID := range taskUUIDs {
		task, err := r.GetByUUID(ctx, taskUUID)
		if err != nil {
			return nil, err
		}

		result = append(result, task)
	}

	return result, nil
}

func (r *TaskRepository) GetByUUID(ctx context.Context, UUID string) (*domain.Task, error) {
	hset := r.api.HGetAll(ctx, fmt.Sprintf("task:%s", UUID))
	if hset.Err() != nil {
		return nil, fmt.Errorf("get task: %w", hset.Err())
	}

	var result domain.Task
	if err := hset.Scan(&result); err != nil {
		return nil, fmt.Errorf("fetch task's fields: %w", err)
	}

	return &result, nil
}
