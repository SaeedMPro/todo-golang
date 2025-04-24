package controller

import "fmt"

var commands = []string{
	`help    
	- Show available commands`,
	`list    
	- List all tasks`,
	`add [id] <title> <description>
	- Add a new task`,
	`delete <id>
	- Delete task by id`,
	`complete <id>
	- Mark a task as complete`,
}

func ShowHelp() {
	fmt.Println("You can use this available commands:")
	for idx, cmd := range commands {
		fmt.Printf("%d. %s\n", idx+1, cmd)
	}
}
