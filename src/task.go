package main

import (
	"time"
)

type Task struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	Priority    uint
}

type Tasks []Task
