package main

import "fmt"

func main() {
	name := "Golang Learners"
	fmt.Printf("Welcome, %s, to the Todo List App!\n", name)

	todos := []string{"Buy groceries", "Complete homework", "Read Golang docs"}
	fmt.Println(todos)

	printTodos(todos)

	if len(todos) == 0 {
		fmt.Println("Your todo list is empty!")
	} else {
		fmt.Println("You have tasks to do!")
	}
}

func printTodos(todos []string) {
	for i, todo := range todos {
		fmt.Printf("%d: %s\n", i+1, todo)
	}
}
