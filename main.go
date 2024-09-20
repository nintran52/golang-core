package main

import (
	"fmt"
	"time"
)

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

	tasks := make(chan string)

	go addTask("Buy groceries", tasks)
	go addTask("Complete homework", tasks)

	for i := 0; i < 2; i++ {
		task := <-tasks
		fmt.Printf("Task added: %s\n", task)
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

func addTask(task string, c chan string) {
	time.Sleep(2 * time.Second) // Simulate adding tasks that take time
	c <- task
}
