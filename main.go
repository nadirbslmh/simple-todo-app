package main

import (
	"simple-todo-app/models"
	"simple-todo-app/services"
)

func main() {
	// specify struct / function with uppercase letter first -> accessible across other packages
	// ex: Todo, AddtoDB()
	// specify struct / function with lowercase first -> private (only accessible in the package itself)

	services.AddTodo(models.Todo{})
	services.GetTodoList()
}
