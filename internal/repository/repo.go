package repository

import (
	"errors"
	"fmt"
	"log"
	"otus_project/internal/data"
	"otus_project/internal/model"
	"otus_project/internal/model/common"
	"sync"
	"sync/atomic"
	"time"
)

var (
	userCounter        uint64
	projectCounter     uint64
	taskCounter        uint64
	reminderCounter    uint64
	tagCounter         uint64
	timeEntriesCounter uint64

	UsersMu       sync.RWMutex
	ProjectsMu    sync.RWMutex
	TasksMu       sync.RWMutex
	RemindersMu   sync.RWMutex
	TagsMu        sync.RWMutex
	TimeEntriesMu sync.RWMutex

	Users       []*model.User
	Projects    []*model.Project
	Reminders   []*model.Reminder
	Tags        []*model.Tag
	TimeEntries []*model.TimeEntry
	Tasks       []*model.Task
)

func appendWithLock[T any](mu *sync.RWMutex, slice *[]*T, item *T) {
	mu.Lock()
	defer mu.Unlock()
	*slice = append(*slice, item)
}

func SaveItem(item common.Item) error {
	switch v := item.(type) {
	case *model.User:
		v.ID = uint(atomic.AddUint64(&userCounter, 1))
		v.CreatedAt = time.Now()
		appendWithLock(&UsersMu, &Users, v)
		return data.AppendToFile(data.UsersFile, v)
	case *model.Project:
		v.ID = uint(atomic.AddUint64(&projectCounter, 1))
		v.CreatedAt = time.Now()
		appendWithLock(&ProjectsMu, &Projects, v)
		return data.AppendToFile(data.ProjectsFile, v)
	case *model.Task:
		v.ID = uint(atomic.AddUint64(&taskCounter, 1))
		v.CreatedAt = time.Now()
		appendWithLock(&TasksMu, &Tasks, v)
		return data.AppendToFile(data.TasksFile, v)
	case *model.Reminder:
		v.ID = uint(atomic.AddUint64(&reminderCounter, 1))
		v.CreatedAt = time.Now()
		appendWithLock(&RemindersMu, &Reminders, v)
		return data.AppendToFile(data.RemindersFile, v)
	case *model.Tag:
		v.ID = uint(atomic.AddUint64(&tagCounter, 1))
		v.CreatedAt = time.Now()
		appendWithLock(&TagsMu, &Tags, v)
		return data.AppendToFile(data.TagsFile, v)
	case *model.TimeEntry:
		v.ID = uint(atomic.AddUint64(&timeEntriesCounter, 1))
		v.CreatedAt = time.Now()
		appendWithLock(&TimeEntriesMu, &TimeEntries, v)
		return data.AppendToFile(data.TimeEntriesFile, v)
	default:
		return errors.New(fmt.Sprintf("Error saving: %v", v.GetItem()))
	}

}

func SaveAllItems(itemType string) error {
	log.Printf("Saving %s with a new id", itemType)
	switch itemType {
	case "user":
		return data.SaveSliceToFile(data.UsersFile, Users)
	case "project":
		return data.SaveSliceToFile(data.ProjectsFile, Projects)
	case "task":
		return data.SaveSliceToFile(data.TasksFile, Tasks)
	case "reminder":
		return data.SaveSliceToFile(data.RemindersFile, Reminders)
	case "tag":
		return data.SaveSliceToFile(data.TagsFile, Tags)
	case "time_entry":
		return data.SaveSliceToFile(data.TimeEntriesFile, TimeEntries)
	default:
		return fmt.Errorf("unknown item type: %s", itemType)
	}
}
