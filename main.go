package main

import (
	"fmt"
	"sync"
)

type Todo struct {
	Task      string
	Completed bool
}

var mu sync.Mutex
var todoList = []string{}

func main() {
	defer fmt.Println("Program finished.") // Last execution before program exit
	fmt.Println("Starting program...")

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

	var wg sync.WaitGroup

	tasks := []string{"Buy groceries", "Complete homework", "Go for a run"}

	for _, task := range tasks {
		wg.Add(1)
		go addTask(task, &wg)
	}

	wg.Wait()
	fmt.Println("Final Todo List:", todoList)

	err := deleteTask([]Todo{})
	if err != nil {
		fmt.Println("Error:", err)
	}

	doneTask("")
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

func addTask(task string, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	todoList = append(todoList, task)
	mu.Unlock()
}

func deleteTask(tasks []Todo) error {
	if len(tasks) == 0 {
		return fmt.Errorf("Task cannot be empty")
	}

	return nil
}

func doneTask(task string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if task == "" {
		panic("Task cannot be empty")
	}

	fmt.Printf("Task done: %s\n", task)
}
