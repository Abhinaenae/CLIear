package main

import (
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type indexedTask struct {
	Index int
	Task  Task
}

func (tasks *Tasks) printTaskTable() {

	indexedTasks := make([]indexedTask, len(*tasks))
	for i, t := range *tasks {
		indexedTasks[i] = indexedTask{Index: i, Task: t}
	}

	// Sort by priority in descending order
	sort.SliceStable(indexedTasks, func(i, j int) bool {
		return indexedTasks[i].Task.Priority > indexedTasks[j].Task.Priority
	})

	// Create and populate the table
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At", "Priority")

	for _, it := range indexedTasks {
		t := it.Task
		originalIndex := it.Index
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✔️"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(originalIndex), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt, strconv.Itoa(int(t.Priority)))
	}

	table.Render()
}
