package main

import (
	"errors"
	"fmt"
	"time"
)

func (tasks *Tasks) add(title string, priority uint) error {
	if priority < 1 || priority > 5 {
		err := errors.New("invalid priority")
		fmt.Println(err)
		return err
	}
	task := Task{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		Priority:    priority,
	}

	*tasks = append(*tasks, task)

	return nil
}

func (tasks *Tasks) validateIndex(index int) error {

	if index == -5 {
		return nil
	}
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

	if index == -5 {
		*tasks = nil
		return nil
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
