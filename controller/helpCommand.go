package controller

import "fmt"

var commands = []string{
	"help    - Show available commands",
	"list    - List all tasks",
	"add     - Add a new task",
	"delete  - Mark a task as complete",
}

func ShowHelp() {
	fmt.Println("You can use this available commands:")
	for idx, cmd := range commands {
		fmt.Printf("%d. %s\n", idx+1, cmd)
	}
}
