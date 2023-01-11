package main

import (
	"fmt"

	"github.com/komoda102/leaningGo_ToDoWebApi/models"
	// "net/http"
)

func main() {
	//http.ListenAndServe(":3000", nil)
	//
	t := models.NewTodoItem("Test description")
	models.AddTodoItem(t)
	t2 := models.NewTodoItem("Test description 1")
	models.AddTodoItem(t2)
	allTodoItems := models.GetTodoItems()

	for _, todoItem := range allTodoItems {
		checkMark := " "
		if todoItem.Done {
			checkMark = "x"
		}

		fmt.Printf("[%v] %v", checkMark, todoItem.Description)
		fmt.Println()
	}
}
