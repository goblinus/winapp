package redis

import "errors"

var (
	ErrNoTaskFields = errors.New("empty fields set for task")
)
