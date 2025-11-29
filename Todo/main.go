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
func addTodo(t Todo) {
	fmt.Println(t.Task)
}
func saveTodos(todos Todos) {

}

func main() {
	// todosFileOpen, err := os.Open("todos.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer todosFileOpen.Close()

	// todosRead, err := os.ReadFile("todos.json")
	// if err != nil {
	// 	panic(err)
	// }
	// var todos Todos
	// fmt.Println(todosRead)
	// json.Unmarshal([]byte(todosRead), &todos)
	// fmt.Println(todos)

	var todos Todos
	file, _ := os.Open("todos.json")
	defer file.Close()

	data, _ := io.ReadAll(file)
	fmt.Println(string(data)) // actual content
	json.Unmarshal([]byte(data), &todos)

	fmt.Println("Welcome to Todo App")
	fmt.Print("1.add todo\n2.list todos\n3.mark completed\n4.delete todo\n")

	fmt.Println("Enter your choice: ")
	var choice int
	fmt.Scan(&choice)

	reader := bufio.NewReader(os.Stdin)
	switch choice {
	case 1:
		fmt.Println("Enter task: ")
		task, _ := reader.ReadString('\n')
		// task[:len(task)-1] - removes the newline character at the end
		todo := Todo{Task: task[:len(task)-1], IsCompleted: false, CreatedAt: time.Now(), CompletedAt: nil}
		todos = append(todos, todo)
		fmt.Println(todos)

		// finalJson, err := json.Marshal(todos) when using this the json is in one line

		finalJson, err := json.MarshalIndent(todos, "", " ")

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("JSON -> %s", finalJson)

		// save to file
		writeErr := os.WriteFile("todos.json", finalJson, 0644) // 0644 grants read access to all, write access to owner
		if writeErr != nil {
			panic(writeErr) // When panic(err) is called, the current function's execution immediately stops.
		}
		fmt.Println("added todo to todos json successfully")
		// listTodos(todos)

	}

}
