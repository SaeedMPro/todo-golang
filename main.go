package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/SaeedMPro/todo-list/controller"
	"github.com/SaeedMPro/todo-list/model"
	"github.com/SaeedMPro/todo-list/repository"
)

var (
	tasks   []model.Task
	command []string
)

func init() {
	// Loading tasks from task.json into tasks slice
	repository.LoadTasks(&tasks)

	// Set command from args
	command = os.Args[1:]
}

func main() {
	if len(command) == 0 {
		fmt.Println("No command provided.\nTry 'help' to see available commands.")
		return
	}

	switch strings.ToLower(command[0]) {
	case "help":
		controller.ShowHelp()

	case "list":
		if len(command) != 1 {
			fmt.Println("Invalid format for `list` command.\nTry 'help' to see available commands.")
			return
		}

		controller.ListAllTasks(tasks)

	case "show":
		if len(command) != 2 {
			fmt.Println("Invalid format for `show` command.\nTry 'help' to see available commands.")
			return
		}

		id, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		controller.ShowTaskByID(tasks, id)

	case "add":
		if len(command) < 3 || len(command) > 4 {
			fmt.Println("Invalid format for `add` command. add requires at least title and description.\nTry 'help' to see available commands.")
		}

		err := controller.HandleAddCommand(command[1:], &tasks)
		if err != nil {
			log.Fatal(err)
		}

		repository.SaveTasks(tasks)

	case "delete":
		if len(command) != 2 {
			fmt.Println("Invalid format for `delete` command.\nTry 'help' to see available commands.")
			return
		}

		id, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		err = controller.DeleteTaskByID(&tasks, id)
		if err != nil {
			log.Fatal(err)
		}

		repository.SaveTasks(tasks)

	case "complete":
		if len(command) != 2 {
			fmt.Println("Invalid format for `complete` command.\nTry 'help' to see available commands.")
			return
		}

		id, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		err = controller.CompleteTaskByID(&tasks, id)
		if err != nil {
			log.Fatal(err)
		}

		repository.SaveTasks(tasks)

	default:
		fmt.Printf("The command '%s' is not valid. Try 'help'.\n", command[0])
		os.Exit(1)

	}
}
