package logger

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"os"
	"otus_project/internal/repository"
	"sync"
	"time"
)

var (
	logger = log.NewLogfmtLogger(os.Stdout)

	lastUsersLen       int
	lastProjectsLen    int
	lastTasksLen       int
	lastRemindersLen   int
	lastTagsLen        int
	lastTimeEntriesLen int
)

func checkAndLogNewItems[T any](label string, mu *sync.Mutex, slice *[]*T, lastLen *int) {
	mu.Lock()
	defer mu.Unlock()

	currentLen := len(*slice)
	if currentLen > *lastLen {
		newItems := (*slice)[*lastLen:]
		for _, item := range newItems {
			err := logger.Log("msg", "New item added:", "label", label, "item", fmt.Sprintf("%+v", *item))
			if err != nil {
				return
			}
		}
		*lastLen = currentLen
	}
}

func StartSliceLogger(context context.Context) {
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-context.Done():
				err := logger.Log("msg", "context done: slice logger is stopped")
				if err != nil {
					fmt.Println(err)
				}
				return
			case <-ticker.C:
				checkAndLogNewItems("User", &repository.UsersMu, &repository.Users, &lastUsersLen)
				checkAndLogNewItems("Project", &repository.ProjectsMu, &repository.Projects, &lastProjectsLen)
				checkAndLogNewItems("Task", &repository.TasksMu, &repository.Tasks, &lastTasksLen)
				checkAndLogNewItems("Reminder", &repository.RemindersMu, &repository.Reminders, &lastRemindersLen)
				checkAndLogNewItems("Tag", &repository.TagsMu, &repository.Tags, &lastTagsLen)
				checkAndLogNewItems("TimeEntry", &repository.TimeEntriesMu, &repository.TimeEntries, &lastTimeEntriesLen)

			}
		}
	}()
}
