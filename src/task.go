package main

import (
	"time"
)

type Task struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Tasks []Task

func (tasks *Tasks) add(title string) {
	task := Task{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*tasks = append(*tasks, task)
}
