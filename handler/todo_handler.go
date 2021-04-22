package handler

import (
	"net/http"
	"server/model"
	"server/service"

	"github.com/labstack/echo"
)

type Service struct {
	service service.Service
}

type Handler interface {
	SayHello(c echo.Context) error
	GetTodos(c echo.Context) error
	CreateTodo(c echo.Context) error
}

func NewTodoHandler(service service.Service) Handler {
	return Service{service}
}

func (h Service) SayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func (h Service) GetTodos(c echo.Context) error {
	todolist := h.service.GetTodos()
	if todolist == nil {
		return c.JSON(http.StatusNoContent, nil)
	}
	return c.JSON(http.StatusOK, todolist)
}

func (h Service) CreateTodo(c echo.Context) error {
	u := &model.Todo{}
	if err := c.Bind(u); err != nil {
		return err
	}
	err := h.service.CreateTodo(model.Todo{Title: u.Title, Completed: u.Completed})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, u)
}
