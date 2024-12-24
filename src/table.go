package main

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

func (tasks *Tasks) printTaskTable() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At", "Priority")
	for index, t := range *tasks {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✔️"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt, strconv.Itoa(int(t.Priority)))
	}

	table.Render()

}
