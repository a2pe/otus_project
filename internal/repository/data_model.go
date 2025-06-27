package repository

import (
	"otus_project/internal/data"
	"sync"
)

type DataModel struct {
	Mutex    *sync.RWMutex
	Data     any
	FileName string
}

var DataRegistry = map[string]*DataModel{
	"user": {
		Mutex:    &UsersMu,
		Data:     &Users,
		FileName: data.UsersFile,
	},
	"project": {
		Mutex:    &ProjectsMu,
		Data:     &Projects,
		FileName: data.ProjectsFile,
	},
	"task": {
		Mutex:    &TasksMu,
		Data:     &Tasks,
		FileName: data.TasksFile,
	},
	"reminder": {
		Mutex:    &RemindersMu,
		Data:     &Reminders,
		FileName: data.RemindersFile,
	},
	"tag": {
		Mutex:    &TagsMu,
		Data:     &Tags,
		FileName: data.TagsFile,
	},
	"time_entry": {
		Mutex:    &TimeEntriesMu,
		Data:     &TimeEntries,
		FileName: data.TimeEntriesFile,
	},
}
