package models

import (
	"errors"
)

type TodoItem struct {
	ID          int
	Description string
	Done        bool
}

var (
	todoItems []*TodoItem
	nextID    int
)

func NewTodoItem(description string) *TodoItem {
	return &TodoItem{
		Description: description,
		Done:        false,
	}
}

func AddTodoItem(newTodoItem TodoItem) (TodoItem, error) {
	newTodoItem.ID = nextID
	nextID++
	todoItems = append(todoItems, &newTodoItem)
	return newTodoItem, nil
}

func GetTodoItems() []*TodoItem {
	return todoItems
}

func ToggleDoneById(id int) error {
	for _, todoItem := range todoItems {
		if todoItem.ID == id {
			todoItem.Done = !todoItem.Done
			return nil
		}
	}

	return errors.New("todo item not found")
}

func RemoveTodoById(id int) error {
	for i, todoItem := range todoItems {
		if todoItem.ID == id {
			todoItems = append(todoItems[:i], todoItems[i+1:]...)
			return nil
		}
	}

	return errors.New("todo item not found")
}
