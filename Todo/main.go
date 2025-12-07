package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// A normal time.Time cannot be nil. Its zero value is always: 0001-01-01 00:00:00 +0000 UTC
// a pointer can be nil *time.Time ;f we want to represent absence of time value
type Todo struct {
	Task        string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	UpdatedAt   time.Time
}
type Todos []Todo

func listTodos(todos Todos) {
	for i, t := range todos {

		if t.IsCompleted == true {
			fmt.Printf("%d %s - ✅", i+1, t.Task)
		} else {
			fmt.Printf("%d %s - ❌", i+1, t.Task)
		}
		fmt.Println()
	}
}
func markAsCompleted(todos *Todos, todoId int) {
	if len(*todos) > 0 {
		todoIdx := todoId - 1
		if todoIdx < 0 || todoIdx >= len(*todos) {
			panic("Enter valid todoId")
		}
		todo := &(*todos)[todoIdx]
		if !todo.IsCompleted {
			todo.IsCompleted = true
			todo.UpdatedAt = time.Now()
			todo.CompletedAt = &todo.UpdatedAt
		}
	} else {
		fmt.Println("No todos to mark completed")
	}
}
func deleteTodo(todos *Todos, todoId int) {
	if len(*todos) > 0 {

		if todoId < 1 || todoId > len(*todos) {
			fmt.Println("Invalid todo, select valid todo")
		}

		*todos = append((*todos)[:todoId-1], (*todos)[todoId:]...)
	} else {
		fmt.Println("No todos to delete")
	}
}
func saveTodos(todos *Todos) {
	finalJson, err := json.MarshalIndent(*todos, "", " ")

	if err != nil {
		fmt.Println(err)
	}
	// save to file
	writeErr := os.WriteFile("todos.json", finalJson, 0644) // 0644 grants read access to all, write access to owner
	if writeErr != nil {
		panic(writeErr) // When panic(err) is called, the current function's execution immediately stops.
	}
}
func main() {
	var todos Todos
	file, openErr := os.Open("todos.json")
	if openErr != nil {
		if os.IsNotExist(openErr) {
			fmt.Println("No existing todos found, starting fresh")
		} else {
			panic(openErr)
		}
	} else {
		defer file.Close()

		data, _ := io.ReadAll(file)
		// fmt.Println(string(data)) // actual content
		if len(data) > 0 {

			json.Unmarshal([]byte(data), &todos)

		}
	}
	listTodos(todos)

	for {
		fmt.Println("Welcome to Todo App")
		fmt.Print("1.add todo\n2.mark completed\n3.delete todo\n4.exit app\n\n")

		fmt.Println("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		reader := bufio.NewReader(os.Stdin)
		switch choice {
		case 1:
			fmt.Println("Enter task: ")
			task, _ := reader.ReadString('\n')
			// task[:len(task)-1] - removes the newline character at the end
			todo := Todo{Task: task[:len(task)-1], IsCompleted: false, CreatedAt: time.Now(), CompletedAt: nil, UpdatedAt: time.Now()}
			todos = append(todos, todo)
			saveTodos(&todos)
			listTodos(todos)
		case 2:
			fmt.Println("Enter todo number to mark completed: ")
			var todoId int
			fmt.Scan(&todoId)
			markAsCompleted(&todos, todoId)
			saveTodos(&todos)
			listTodos(todos)

		case 3:
			fmt.Println("Enter task number to remove todo.")
			var todoId int
			fmt.Scan(&todoId)
			deleteTodo(&todos, todoId)
			saveTodos(&todos)
			listTodos(todos)
		case 4:
			return
		default:
			fmt.Println("Invalid operation")
		}

	}
}
