package todo

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// gRPC server
type Server struct {
	Todos []*TodoObject
}

func (s *Server) AddTodo(ctx context.Context, newTodo *AddTodoParams) (*TodoObject, error) {
	log.Printf("Received new task %s", newTodo.Task)
	u, err := uuid.NewV4()
	_ = err
	todoObject := &TodoObject{

		Id:   u.String(),
		Task: newTodo.Task,
	}
	s.Todos = append(s.Todos, todoObject)
	return todoObject, nil
}

func (s *Server) GetTodos(ctx context.Context, _ *GetTodoParams) (*TodoResponse, error) {
	log.Printf("get tasks")
	return &TodoResponse{Todos: s.Todos}, nil
}

func (s *Server) DeleteTodo(ctx context.Context, delTodo *DeleteTodoParams) (*DeleteResponse, error) {
	var updatedTodos []*TodoObject
	for index, todo := range s.Todos {
		if todo.Id == delTodo.Id {
			updatedTodos = append(s.Todos[:index], s.Todos[index+1:]...)
			break
		}
	}
	s.Todos = updatedTodos
	return &DeleteResponse{Message: "success"}, nil
}
