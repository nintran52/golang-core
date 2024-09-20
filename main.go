package main

import (
	"fmt"
)

type Todo struct {
	Task      string
	Completed bool
}

type TodoList struct {
	todos []Todo
}

func (t *TodoList) Add(task string) {
	todo := Todo{Task: task, Completed: false}
	t.todos = append(t.todos, todo)
	fmt.Printf("Added: %s\n", task)
}

func (t *TodoList) Complete(task string) {
	for i := range t.todos {
		if t.todos[i].Task == task {
			t.todos[i].Completed = true
			fmt.Printf("Completed: %s\n", task)
			return
		}
	}
	fmt.Printf("Task not found: %s\n", task)
}

func main() {
	todoList := TodoList{}
	todoList.Add("Buy groceries")
	todoList.Add("Complete homework")

	todoList.Complete("Buy groceries")
}
