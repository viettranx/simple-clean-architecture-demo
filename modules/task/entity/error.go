package entity

import "errors"

var (
	ErrTaskNotFound       = errors.New("task not found")
	ErrCannotCreateTask   = errors.New("cannot create task")
	ErrCannotUpdateTask   = errors.New("cannot update task")
	ErrCannotDeleteTask   = errors.New("cannot delete task")
	ErrTitleCannotBeBlank = errors.New("title cannot be blank")
	ErrStatusNotValid     = errors.New("status must be 'doing' or 'done'")
)

func validateTitle(title string) error {
	if title == "" {
		return ErrTitleCannotBeBlank
	}

	return nil
}

func validateStatus(status string) error {
	if status != StatusDoing && status != StatusDone {
		return ErrStatusNotValid
	}

	return nil
}
