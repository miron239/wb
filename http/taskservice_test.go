package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/miron239/wb/config"
	"github.com/miron239/wb/domain"
	"github.com/miron239/wb/mock"
)

func TestGetTask(t *testing.T) {
	var ts mock.TaskService
	var ac mock.AlwaysAllow
	tsHTTP := &TaskService{Service: &ts, AuthzClient: &ac}

	// Mock Task() call.
	ts.TaskFn = func(id int) (*domain.Task, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &domain.Task{ID: 100, Name: "my-task-1"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/blacklist/100", nil)

	server := InitServer(&config.Config{
		Server: config.ServerConf{Port: 8080},
		Frontend: config.FrontendConf{
			Endpoints: []string{"*"},
		},
		Authn: config.AuthnConf{NotEnforced: true},
	})
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.TaskInvoked {
		t.Fatal("expected Task() to be invoked")
	}
}

func TestGetTasks(t *testing.T) {
	var ts mock.TaskService
	var ac mock.AlwaysAllow
	tsHTTP := &TaskService{Service: &ts, AuthzClient: &ac}

	// Mock Tasks() call.
	ts.TasksFn = func() ([]*domain.Task, error) {
		return []*domain.Task{{ID: 100, Name: "my-task-1"}}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/blacklist/", nil)

	server := InitServer(&config.Config{
		Server: config.ServerConf{Port: 8080},
		Frontend: config.FrontendConf{
			Endpoints: []string{"*"},
		},
		Authn: config.AuthnConf{NotEnforced: true},
	})
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.TasksInvoked {
		t.Fatal("expected Tasks() to be invoked")
	}
}

func TestCreateTask(t *testing.T) {
	var ts mock.TaskService
	var ac mock.AlwaysAllow
	tsHTTP := &TaskService{Service: &ts, AuthzClient: &ac}

	// Mock our CreateTask() call.
	ts.CreateTaskFn = func(task *domain.Task) error {
		if task.Name != "my-task-1" {
			t.Fatalf("unexpected name: %s", task.Name)
		}
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	request, err := json.Marshal(&CreateTaskRequest{&Task{Name: "my-task-1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}
	reader := strings.NewReader(string(request))
	r, _ := http.NewRequest("POST", "/blacklist/", reader)

	server := InitServer(&config.Config{
		Server: config.ServerConf{Port: 8080},
		Frontend: config.FrontendConf{
			Endpoints: []string{"*"},
		},
		Authn: config.AuthnConf{NotEnforced: true},
	})
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.CreateTaskInvoked {
		t.Fatal("expected CreateTask() to be invoked")
	}
}

func TestCreateTaskForbidden(t *testing.T) {
	var ts mock.TaskService
	var ac mock.AlwaysDeny
	tsHTTP := &TaskService{Service: &ts, AuthzClient: &ac}

	// Mock our CreateTask() call.
	ts.CreateTaskFn = func(task *domain.Task) error {
		if task.Name != "my-task-1" {
			t.Fatalf("unexpected name: %s", task.Name)
		}
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	request, err := json.Marshal(&CreateTaskRequest{&Task{Name: "my-task-1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}
	reader := strings.NewReader(string(request))
	r, _ := http.NewRequest("POST", "/blacklist/", reader)

	server := InitServer(&config.Config{
		Server: config.ServerConf{Port: 8080},
		Frontend: config.FrontendConf{
			Endpoints: []string{"*"},
		},
		Authn: config.AuthnConf{NotEnforced: true},
	})
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	if got, want := w.Code, 403; got != want {
		t.Fatalf("response code - got: %d, want: %d", got, want)
	}

	// Validate mock.
	if ts.CreateTaskInvoked {
		t.Fatal("not expected CreateTask() to be invoked")
	}
}

func TestDeleteTask(t *testing.T) {
	var ts mock.TaskService
	var ac mock.AlwaysAllow
	tsHTTP := &TaskService{Service: &ts, AuthzClient: &ac}

	// Mock DeleteTask() call.
	ts.DeleteTaskFn = func(id int) error {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return nil
	}
	// Mock Task() call.
	ts.TaskFn = func(id int) (*domain.Task, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &domain.Task{ID: 100, Name: "my-task-1"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/blacklist/100", nil)

	server := InitServer(&config.Config{
		Server: config.ServerConf{Port: 8080},
		Frontend: config.FrontendConf{
			Endpoints: []string{"*"},
		},
		Authn: config.AuthnConf{NotEnforced: true},
	})
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.DeleteTaskInvoked {
		t.Fatal("expected DeleteTask() to be invoked")
	}
}

func TestDeleteTasks(t *testing.T) {
	var ts mock.TaskService
	var ac mock.AlwaysAllow
	tsHTTP := &TaskService{Service: &ts, AuthzClient: &ac}

	// Mock Tasks() call.
	ts.DeleteTasksFn = func() error {
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/blacklist/", nil)

	server := InitServer(&config.Config{
		Server: config.ServerConf{Port: 8080},
		Frontend: config.FrontendConf{
			Endpoints: []string{"*"},
		},
		Authn: config.AuthnConf{NotEnforced: true},
	})
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.DeleteTasksInvoked {
		t.Fatal("expected DeleteTasks() to be invoked")
	}
}
