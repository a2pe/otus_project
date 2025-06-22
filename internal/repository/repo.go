package repository

import (
	"errors"
	"fmt"
	"otus_project/internal/data"
	"otus_project/internal/model"
	"otus_project/internal/model/common"
	"sync"
)

var (
	UsersMu       sync.Mutex
	ProjectsMu    sync.Mutex
	TasksMu       sync.Mutex
	RemindersMu   sync.Mutex
	TagsMu        sync.Mutex
	TimeEntriesMu sync.Mutex

	Users       []*model.User
	Projects    []*model.Project
	Reminders   []*model.Reminder
	Tags        []*model.Tag
	TimeEntries []*model.TimeEntry
	Tasks       []*model.Task
)

func appendWithLock[T any](mu *sync.Mutex, slice *[]*T, item *T) {
	mu.Lock()
	defer mu.Unlock()
	*slice = append(*slice, item)
}

func SaveItem(item common.Item) error {
	switch v := item.(type) {
	case model.User:
		appendWithLock(&UsersMu, &Users, &v)
		return data.AppendToFile(data.UsersFile, &v)
	case model.Project:
		appendWithLock(&ProjectsMu, &Projects, &v)
		return data.AppendToFile(data.ProjectsFile, &v)
	case model.Task:
		appendWithLock(&TasksMu, &Tasks, &v)
		return data.AppendToFile(data.TasksFile, &v)
	case model.Reminder:
		appendWithLock(&RemindersMu, &Reminders, &v)
		return data.AppendToFile(data.RemindersFile, &v)
	case model.Tag:
		appendWithLock(&TagsMu, &Tags, &v)
		return data.AppendToFile(data.TagsFile, &v)
	case model.TimeEntry:
		appendWithLock(&TimeEntriesMu, &TimeEntries, &v)
		return data.AppendToFile(data.TimeEntriesFile, &v)
	default:
		return errors.New(fmt.Sprintf("Error saving: %v", v.GetItem()))
	}
}
