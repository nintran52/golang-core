package main

import "fmt"

func main() {
	name := "Golang Learners"
	fmt.Printf("Welcome, %s, to the Todo List App!\n", name)

	todos := []string{"Buy groceries", "Complete homework", "Read Golang docs"}
	fmt.Println(todos)

	for i, todo := range todos {
		fmt.Printf("%d: %s\n", i+1, todo)
	}
}
