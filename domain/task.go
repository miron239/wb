package domain

import (
	"time"
)

type Task struct {
	ID          int
	UserId      string
	Name        string
	Description string
	Phone       string
	DueAt       time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskService interface {
	Task(id int) (*Task, error)
	TaskByName(name string) (*Task, error)
	TaskByPhone(phone string) (*Task, error)
	Tasks() ([]*Task, error)
	CreateTask(t *Task) error
	DeleteTask(id int) error
	DeleteTasks() error
}
