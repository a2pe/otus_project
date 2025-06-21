package repository

import (
	"errors"
	"fmt"
	"github.com/go-kit/log"
	"os"
	"otus_project/internal/model"
	"otus_project/internal/model/common"
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

	usersMu       sync.Mutex
	projectsMu    sync.Mutex
	tasksMu       sync.Mutex
	remindersMu   sync.Mutex
	tagsMu        sync.Mutex
	timeEntriesMu sync.Mutex

	users       []*model.User
	projects    []*model.Project
	reminders   []*model.Reminder
	tags        []*model.Tag
	timeEntries []*model.TimeEntry
	tasks       []*model.Task
)

func appendWithLock[T any](mu *sync.Mutex, slice *[]*T, item *T) {
	mu.Lock()
	defer mu.Unlock()
	*slice = append(*slice, item)
}

func SaveItem(item common.Item) error {
	switch v := item.(type) {
	case model.User:
		appendWithLock(&usersMu, &users, &v)
	case model.Project:
		appendWithLock(&projectsMu, &projects, &v)
	case model.Task:
		appendWithLock(&tasksMu, &tasks, &v)
	case model.Reminder:
		appendWithLock(&remindersMu, &reminders, &v)
	case model.Tag:
		appendWithLock(&tagsMu, &tags, &v)
	case model.TimeEntry:
		appendWithLock(&timeEntriesMu, &timeEntries, &v)
	default:
		return errors.New(fmt.Sprintf("Error saving: %v", v.GetItem()))
	}
	return nil
}

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

func StartSliceLogger() {
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)

			checkAndLogNewItems("User", &usersMu, &users, &lastUsersLen)
			checkAndLogNewItems("Project", &projectsMu, &projects, &lastProjectsLen)
			checkAndLogNewItems("Task", &tasksMu, &tasks, &lastTasksLen)
			checkAndLogNewItems("Reminder", &remindersMu, &reminders, &lastRemindersLen)
			checkAndLogNewItems("Tag", &tagsMu, &tags, &lastTagsLen)
			checkAndLogNewItems("TimeEntry", &timeEntriesMu, &timeEntries, &lastTimeEntriesLen)
		}
	}()
}
