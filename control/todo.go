package control

import (
	"github.com/chiraponkub/Todo-Go/present/structdb"
	"github.com/chiraponkub/Todo-Go/present/structure"
	"gorm.io/gorm"
	"time"
)

func (ctrl APIControl) GetTodo(todo []structdb.User) (res interface{}, Error error) {
	if len(todo) < 1 {
		res = []structdb.Todo{}
		return
	}
	var userId uint
	var name string
	var ArrayTodo []structure.Todo
	for _, m1 := range todo {
		userId = m1.ID
		name = m1.FirstName + " " + m1.LastName
		for _, m2 := range m1.Todo {
			Todo := structure.Todo{
				Id:       m2.ID,
				Text:     m2.Text,
				IsActive: m2.IsActive,
			}
			ArrayTodo = append(ArrayTodo, Todo)
		}
	}
	res = structure.UserRefer{
		UserId: userId,
		Name:   name,
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
