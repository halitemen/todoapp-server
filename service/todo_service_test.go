package service

import (
	"server/model"
	"server/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	var todoList []model.Todo
	todoList = append(todoList, model.Todo{
		Title:     "todo",
		Completed: false,
	})

	repository := repository.NewMockRepository(controller)
	repository.EXPECT().GetTodos().Return(todoList).Times(1)
	service := TodoService{repository}

	all := service.GetTodos().([]model.Todo)
	assert.NotNil(t, all)
	assert.NotEmpty(t, all)
	assert.Equal(t, 1, len(all))
	assert.Equal(t, "todo", all[0].Title)
	assert.False(t, all[0].Completed)
}

func TestCreateOne(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := repository.NewMockRepository(controller)
	repository.EXPECT().CreateTodo(model.Todo{Title: "todo", Completed: false}).Return(model.Todo{
		Title:     "todo",
		Completed: false,
	}).Times(1)

	service := NewTodoService(repository)
	created := service.CreateTodo(model.Todo{Title: "todo", Completed: false})

	assert.NotNil(t, created)
	todo := created.(model.Todo)
	assert.Equal(t, "todo", todo.Title)
	assert.False(t, todo.Completed)
}
