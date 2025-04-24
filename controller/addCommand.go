package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/SaeedMPro/todo-list/model"
)

func HandleAddCommand(cmd []string, tasks *[]model.Task) error {
	var id int
	var title, description string
	var autoGenerateID bool

	switch len(cmd) {
	case 2:
		autoGenerateID = true
		title = cmd[0]
		description = cmd[1]
	case 3:
		parsedID, err := strconv.Atoi(cmd[0])
		if err != nil {
			autoGenerateID = true
			title = cmd[1]
			description = cmd[2]
		} else {
			// Check for duplicate ID
			for _, t := range *tasks {
				if t.ID == parsedID {
					return fmt.Errorf("task with ID %d already exists", parsedID)
				}
			}

			id = parsedID
			title = cmd[1]
			description = cmd[2]
		}

	default:
		return errors.New("invalid format for `add` command.\ntry 'help' to see available commands")
	}

	if autoGenerateID {
		id = generateNextID(*tasks)
	}

	newTask := model.Task{
		ID:          id,
		Title:       title,
		Description: description,
		IsDone:      false,
	}

	*tasks = append(*tasks, newTask)
	fmt.Printf("Task with id <%d> created successfully", id)

	return nil
}

func generateNextID(tasks []model.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	id := maxID + 1
	return id
}
