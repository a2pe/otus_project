package repository

import (
	"errors"
	"fmt"
	"otus_project/internal/model"
)

type Item interface {
	GetItem() uint
}

var (
	users       []*model.User
	projects    []*model.Project
	reminders   []*model.Reminder
	tags        []*model.Tag
	timeEntries []*model.TimeEntry
	tasks       []*model.Task
)

func SaveItem(item Item) error {
	switch v := item.(type) {
	case model.User:
		users = append(users, &v)
		fmt.Println("User saved:", v)
	case model.Project:
		projects = append(projects, &v)
		fmt.Println("Project saved:", v)
	case model.Task:
		tasks = append(tasks, &v)
		fmt.Println("Task saved:", v)
	case model.Reminder:
		reminders = append(reminders, &v)
		fmt.Println("Reminder saved:", v)
	case model.Tag:
		tags = append(tags, &v)
		fmt.Println("Tag saved:", v)
	case model.TimeEntry:
		timeEntries = append(timeEntries, &v)
		fmt.Println("TimeEntry saved:", v)
	default:
		return errors.New(fmt.Sprintf("Error saving: %v", v.GetItem()))
	}
	return nil
}
