package main

import (
	"errors"
	"fmt"
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

func (tasks *Tasks) validateIndex(index int) error {
	if index < 0 || index >= len(*tasks) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (tasks *Tasks) delete(index int) error {
	t := *tasks

	if err := t.validateIndex(index); err != nil {
		return err
	}

	//Removes the item and replaces contents with a new list without that item
	*tasks = append(t[:index], t[index+1:]...)

	return nil
}

func (tasks *Tasks) toggle(index int) error {
	t := *tasks
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted
	return nil
}

func (tasks *Tasks) edit(index int, title string) error {
	t := *tasks
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}
