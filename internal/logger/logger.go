package logger

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"os"
	"otus_project/internal/data"
	"otus_project/internal/repository"
	"sync"
	"time"
)

var (
	logger = log.NewLogfmtLogger(os.Stdout)

	userPath        = data.GetFinalFilePath(data.UsersFile)
	projectPath     = data.GetFinalFilePath(data.ProjectsFile)
	tagsPath        = data.GetFinalFilePath(data.TagsFile)
	tasksPath       = data.GetFinalFilePath(data.TasksFile)
	remindersPath   = data.GetFinalFilePath(data.RemindersFile)
	timeEntriesPath = data.GetFinalFilePath(data.TimeEntriesFile)

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

func StartSliceLogger(ctx context.Context, logger log.Logger) {
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
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

func LoadAll() error {
	if err := data.LoadDataFromFile(userPath, &repository.UsersMu, &repository.Users, &lastUsersLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(projectPath, &repository.ProjectsMu, &repository.Projects, &lastProjectsLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(tasksPath, &repository.TasksMu, &repository.Tasks, &lastTasksLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(tagsPath, &repository.TagsMu, &repository.Tags, &lastTagsLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(remindersPath, &repository.RemindersMu, &repository.Reminders, &lastRemindersLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(timeEntriesPath, &repository.TimeEntriesMu, &repository.TimeEntries, &lastTimeEntriesLen); err != nil {
		return err
	}
	return nil
}
