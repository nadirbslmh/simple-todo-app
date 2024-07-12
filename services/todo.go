package services

import "simple-todo-app/models"

var database []models.Todo = []models.Todo{}

func AddTodo(newTodo models.Todo) {
	database = append(database, newTodo)
}

func GetTodoList() []models.Todo {
	return database
}
