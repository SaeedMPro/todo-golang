package controller

import (
	"errors"
	"fmt"

	"github.com/SaeedMPro/todo-list/model"
)

func CompleteTaskByID(tasks *[]model.Task, id int) error {
	if tasks == nil || len(*tasks) == 0 {
		return errors.New("no tasks available to marked as completed")
	}

	for idx, t := range *tasks {
		if t.ID == id {
			if (*tasks)[idx].IsDone {
				return fmt.Errorf("task with id <%d> is already completed", id)
			}

			(*tasks)[idx].IsDone = true
			fmt.Printf("Task with id <%d> marked as completed\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with id <%d> not found", id)
}
