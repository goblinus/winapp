package inmemory

import "errors"

var (
	ErrTaskNotFound = errors.New("no task found")
	ErrUserNotFound = errors.New("no user found")
)
