package mock

import "github.com/miron239/wb/domain"

type TaskService struct {
	TaskFn      func(id int) (*domain.Task, error)
	TaskInvoked bool

	TaskFnName      func(name string) (*domain.Task, error)
	NameTaskInvoked bool

	TaskFnPhone      func(phone string) (*domain.Task, error)
	PhoneTaskInvoked bool

	TasksFn      func() ([]*domain.Task, error)
	TasksInvoked bool

	CreateTaskFn      func(t *domain.Task) error
	CreateTaskInvoked bool

	DeleteTaskFn      func(id int) error
	DeleteTaskInvoked bool

	DeleteTasksFn      func() error
	DeleteTasksInvoked bool
}

func (s *TaskService) Task(id int) (*domain.Task, error) {
	s.TaskInvoked = true
	return s.TaskFn(id)
}

func (s *TaskService) TaskByName(name string) (*domain.Task, error) {
	s.NameTaskInvoked = true
	return s.TaskFnName(name)
}

func (s *TaskService) TaskByPhone(phone string) (*domain.Task, error) {
	s.PhoneTaskInvoked = true
	return s.TaskFnPhone(phone)
}

func (s *TaskService) Tasks() ([]*domain.Task, error) {
	s.TasksInvoked = true
	return s.TasksFn()
}

func (s *TaskService) CreateTask(t *domain.Task) error {
	s.CreateTaskInvoked = true
	return s.CreateTaskFn(t)
}

func (s *TaskService) DeleteTask(id int) error {
	s.DeleteTaskInvoked = true
	return s.DeleteTaskFn(id)
}

func (s *TaskService) DeleteTasks() error {
	s.DeleteTasksInvoked = true
	return s.DeleteTasksFn()
}
