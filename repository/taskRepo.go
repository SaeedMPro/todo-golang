package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/SaeedMPro/todo-list/model"
)

func LoadTasks(tasks *[]model.Task) {
	file, err := os.Open("./repository/task.json")
	if err != nil {
		panic(fmt.Sprintf("Error in read json file : %v", err))
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshaling tasks: %v", err))
	}
}

func SaveTasks(tasks []model.Task) {
	file, err := os.OpenFile("./repository/task.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("Error opening tasks file: %v", err))
	}
	defer file.Close()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("Error marshaling tasks: %v", err))
	}

	_, err = file.Write(data)
	if err != nil {
		panic(fmt.Sprintf("Error writing to file: %v", err))
	}
}
