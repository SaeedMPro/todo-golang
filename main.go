package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/SaeedMPro/todo-list/model"
)

var (
	tasks []model.Task
	command []string
)

func init() {

	// Loading tasks from task.json into tasks slice
	fmt.Println("Loading tasks...")

	file, err := os.Open("./task.json")
	if err != nil {
		panic(fmt.Sprintf("Error in read json file : %v", err))
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &tasks)

	time.Sleep(2 * time.Second)

	// Set command from args
	command = os.Args[1:]
}

func main() {
	if len(command) == 0 {
		fmt.Println("No command provided. Try 'help' to see available commands.")
		return
	}
	
	switch command[0] {
	case "help":
		// TODO: print help
	case "show":
		// TODO: show tasks
	case "add":
		// TODO: add task
	case "delete":
		// TODO: delete task
	default:
		fmt.Printf("The command '%s' is not valid. Try 'help'.\n", command[0])
		os.Exit(1)
	}
}

