package controller

import (
	"fmt"

	"github.com/SaeedMPro/todo-list/model"
)

func ListAllTasks(tasks []model.Task) {
	fmt.Println("\nAll Tasks:")

	fmt.Println("\nUndone Tasks:")
	getTasksByStatus(tasks, false, showTasks)

	fmt.Println("\nDone Tasks:")
	getTasksByStatus(tasks, true, showTasks)
}

func getTasksByStatus(ts []model.Task, done bool, show func([]model.Task)) {
	var filtered []model.Task
	for _, task := range ts {
		if task.IsDone == done {
			filtered = append(filtered, task)
		}
	}
	show(filtered)
}

func showTasks(ts []model.Task) {
	if len(ts) == 0 {
		fmt.Println("  No tasks in this category")
		return
	}

	for _, t := range ts {
		status := " "
		if t.IsDone {
			status = "✓"
		}
		fmt.Printf("  [%s] %d. %s\n", status, t.ID, t.Title)
	}
}
