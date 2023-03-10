package http

import (
	"fmt"
	"strings"
	"time"

	"github.com/miron239/wb/domain"
)

// This section is implemented based on
// https://github.com/swaggo/swag#use-swaggertype-tag-to-supported-custom-type
type TimestampTime struct {
	time.Time
}

func (t *TimestampTime) MarshalJSON() ([]byte, error) {
	bin := make([]byte, 0, len("2019-10-12T07:20:50.52Z"))
	bin = append(bin, fmt.Sprintf("\"%s\"", t.Format(time.RFC3339))...)
	return bin, nil
}

func (t *TimestampTime) UnmarshalJSON(bin []byte) error {
	s := strings.Trim(string(bin), string([]byte{0, '"'}))
	parsedTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// Task is a struct with a subset of the fields of domain.Task. It is used when
// task needs to be provided as an input for task creation. So it excludes
// auto-generated fields.
type Task struct {
	UserId      string `json:"userId" example:"miron"`
	Name        string `json:"name" example:"my-task-1"`
	Description string `json:"description" example:"description why"`
	Phone       string `json:"phone" example:"+7931315455"`
}

type CreateTaskRequest struct {
	Task *Task `json:"task" binding:"required"`
}

func (t *Task) httpToModel() *domain.Task {
	return &domain.Task{
		UserId:      t.UserId,
		Name:        t.Name,
		Description: t.Description,
		Phone:       t.Phone,
	}
}
