package model

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"explanation"`
	IsDone      bool   `json:"isDone"`
}

func (t Task) GetTitle() string {
	return t.Title
}

func (t Task) GetExplanation() string {
	return t.Description
}
