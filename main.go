package main

import "fmt"

type Todo struct {
	Task      string
	Completed bool
}

func main() {
	name := "Golang Learners"
	fmt.Printf("Welcome, %s, to the Todo List App!\n", name)

	todos := []Todo{
		{"Buy groceries", false},
		{"Complete homework", true},
	}
	fmt.Println(todos)

	printTodos(todos)

	if len(todos) == 0 {
		fmt.Println("Your todo list is empty!")
	} else {
		fmt.Println("You have tasks to do!")
	}
}

func printTodos(todos []Todo) {
	for _, todo := range todos {
		status := "not completed"
		if todo.Completed {
			status = "completed"
		}
		fmt.Printf("%s: %s\n", todo.Task, status)
	}
}
