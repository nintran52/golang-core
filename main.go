package main

import "fmt"

func main() {
	name := "Golang Learners"
	fmt.Printf("Welcome, %s, to the Todo List App!\n", name)

	todos := map[string]bool{
		"Buy groceries":     false,
		"Complete homework": true,
	}
	fmt.Println(todos)

	printTodos(todos)

	if len(todos) == 0 {
		fmt.Println("Your todo list is empty!")
	} else {
		fmt.Println("You have tasks to do!")
	}
}

func printTodos(todos map[string]bool) {
	for task, completed := range todos {
		status := "not completed"
		if completed {
			status = "completed"
		}
		fmt.Printf("%s: %s\n", task, status)
	}
}
