package service

import (
	"server/model"
	"server/repository"
)

type Service interface {
	GetTodos() interface{}
	CreateTodo(i interface{}) interface{}
}

type TodoService struct {
	repository repository.Repository
}

func NewTodoService(repository repository.Repository) Service {
	return TodoService{repository}
}

func (s TodoService) GetTodos() interface{} {
	return s.repository.GetTodos()
}

func (s TodoService) CreateTodo(i interface{}) interface{} {
	req, _ := i.(model.Todo)
	todo := model.Todo{
		Title:     req.Title,
		Completed: false,
	}
	return s.repository.CreateTodo(todo)
}
