package main

import "todolist/todolist"

func main() {
	var manager todolist.TaskManager = &todolist.TodoList{}

	manager.Add("Buy groceries")
	manager.Add("Complete homework")

	manager.Complete("Buy groceries")
}
