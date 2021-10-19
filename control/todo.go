package control

import (
	"ProjectEcho/present/structdb"
	"ProjectEcho/present/structure"
	"gorm.io/gorm"
	"time"
)

func (ctrl APIControl) GetTodo(todo []structdb.Todo) (res interface{}, Error error) {

	if len(todo) < 1 {
		res = []structdb.Todo{}
		return
	}
	var userId uint
	var ArrayTodo []structure.Todo
	for _, m1 := range todo {
		userId = m1.UserRefer
		Todo := structure.Todo{
			Id:       m1.ID,
			Text:     m1.Text,
			IsActive: m1.IsActive,
		}
		ArrayTodo = append(ArrayTodo, Todo)
	}
	res = structure.UserRefer{
		UserId: userId,
		Todo:   ArrayTodo,
	}
	return
}

func (ctrl APIControl) AddTodo(UserId uint, todo *structure.AddTodo) (res structdb.Todo, Error error) {
	data := structdb.Todo{
		Text:      todo.Text,
		IsActive:  true,
		UserRefer: UserId,
	}
	res = data
	return
}

func (ctrl APIControl) EditTodo(TodoId, UserId uint, todo *structure.EditTodo) (res structdb.Todo, Error error) {
	data := structdb.Todo{
		Model: gorm.Model{
			ID:        TodoId,
			UpdatedAt: time.Now(),
		},
		Text:      todo.Text,
		IsActive:  todo.IsActive,
		UserRefer: UserId,
	}
	res = data
	return
}
