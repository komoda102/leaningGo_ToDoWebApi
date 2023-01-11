package models

type TodoItem struct {
	ID          int
	Description string
	Done        bool
}

var (
	todoItems []TodoItem
)

func NewTodoItem(description string) TodoItem {
	return TodoItem{
		Description: description,
		Done:        false,
	}
}

func AddTodoItem(newTodoItem TodoItem) (TodoItem, error) {
	todoItems = append(todoItems, newTodoItem)
	return newTodoItem, nil
}

func GetTodoItems() []TodoItem {
	return todoItems
}
