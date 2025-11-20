package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
)
