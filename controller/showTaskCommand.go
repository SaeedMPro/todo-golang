package controller

import (
	"fmt"

	"github.com/SaeedMPro/todo-list/model"
)

func ShowTaskByID(tasks []model.Task, id int) {
	for _, task := range tasks {
		if task.ID == id {
			status := " "
			if task.IsDone {
				status = "✓"
			}

			fmt.Printf("\n[%s] %d. %s\n", status, task.ID, task.Title)
			fmt.Printf("   Description: %s\n", task.Description)
			fmt.Printf("   Status:      ")
			if task.IsDone {
				fmt.Print("Completed\n\n")
			} else {
				fmt.Print("Pending\n\n")
			}
			return
		}
	}
	fmt.Printf("\nTask with ID %d not found\n\n", id)
}
