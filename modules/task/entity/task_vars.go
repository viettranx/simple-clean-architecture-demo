package entity

import "strings"

// TaskCreationData should be used for parsing data from HTTP Request body
// In many cases, main entity model contains data fields that clients
// cannot send, especially data from other modules (or services).
type TaskCreationData struct {
	Id          string `json:"-"` // just carry inserted id, so json tag is omitted
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Validate is entity logic, ensure data is validated
func (t *TaskCreationData) Validate() error {
	t.Title = strings.TrimSpace(t.Title)

	if err := validateTitle(t.Title); err != nil {
		return err
	}

	t.Status = strings.ToLower(strings.TrimSpace(t.Status))

	if err := validateStatus(t.Status); err != nil {
		return err
	}

	return nil
}

// TaskPatchData should be used for parsing data from HTTP Request body.
// The data fields are limited and pointers at all. The pointers help
// distinguish if they have carry value.
// Of course, we just do update for pointers are not nil.
type TaskPatchData struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

func (t *TaskPatchData) Validate() error {
	if s := t.Title; s != nil {
		title := strings.TrimSpace(*s)

		if err := validateTitle(title); err != nil {
			return err

		}
		t.Title = &title
	}

	if s := t.Description; s != nil {
		des := strings.TrimSpace(*s)

		if err := validateTitle(des); err != nil {
			return err

		}
		t.Description = &des
	}

	return nil
}

type Filter struct {
	Status *string `json:"status,omitempty" form:"status"`
}
