package domain

import (
	"encoding/json"
	"time"
)

type (
	Task struct {
		UUID        string    `json:"uuid" redis:"uuid"`
		Name        string    `json:"name" redis:"name"`
		Description string    `json:"description" redis:"description"`
		Timestamp   time.Time `json:"timestamp" redis:"timestamp"`
	}

	TaskInput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func (t *TaskInput) Unmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}
