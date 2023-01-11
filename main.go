package main

import (
	"fmt"

	"github.com/komoda102/leaningGo_ToDoWebApi/models"
	// "net/http"
)

func main() {
	//http.ListenAndServe(":3000", nil)
	//
	t := models.NewTodoItem("Wash your teeth")
	models.AddTodoItem(*t)
	t = models.NewTodoItem("Take out the trash")
	models.AddTodoItem(*t)
	printTodoList()

	models.ToggleDoneById(1)
	printTodoList()

	models.ToggleDoneById(1)
	printTodoList()

	models.RemoveTodoById(0)
	printTodoList()
}

func printTodoList() {
	allTodoItems := models.GetTodoItems()
	for _, todoItem := range allTodoItems {
		checkMark := " "
		if todoItem.Done {
			checkMark = "x"
		}

		fmt.Printf("[%v] %v - %v", checkMark, todoItem.Description, todoItem)
		fmt.Println()
	}
	fmt.Println()
}
