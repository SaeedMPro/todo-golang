package controller

import (
	"errors"
	"fmt"

	"github.com/SaeedMPro/todo-list/model"
)

func DeleteTaskByID(tasks *[]model.Task, id int) error {
	if tasks == nil || len(*tasks) == 0 {
		return errors.New("no tasks available to delete")
	}

	for idx, t := range *tasks {
		if t.ID == id {
			*tasks = append((*tasks)[:idx], (*tasks)[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}
