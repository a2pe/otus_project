package handler

import (
	"otus_project/internal/model"
	"otus_project/internal/model/common"
)

func getEmptyItem(itemType string) common.Item {
	switch itemType {
	case "user":
		return &model.User{}
	case "project":
		return &model.Project{}
	case "task":
		return &model.Task{}
	case "reminder":
		return &model.Reminder{}
	case "tag":
		return &model.Tag{}
	case "time_entry":
		return &model.TimeEntry{}
	default:
		return nil
	}
}

//func findByID[T any](slice []*T, id int) *T {
//	for _, item := range slice {
//		if i, ok := any(item).(common.Item); ok && int(i.GetItem()) == id {
//			return item
//		}
//	}
//	return nil
//}

//func getLockedByID(itemType string, id int) (any, bool) {
//	switch itemType {
//	case "user":
//		repository.UsersMu.RLock()
//		defer repository.UsersMu.RUnlock()
//		item := findByID(repository.Users, id)
//		return item, item != nil
//
//	case "project":
//		repository.ProjectsMu.RLock()
//		defer repository.ProjectsMu.RUnlock()
//		item := findByID(repository.Projects, id)
//		return item, item != nil
//
//	case "task":
//		repository.TasksMu.RLock()
//		defer repository.TasksMu.RUnlock()
//		item := findByID(repository.Tasks, id)
//		return item, item != nil
//
//	case "reminder":
//		repository.RemindersMu.RLock()
//		defer repository.RemindersMu.RUnlock()
//		item := findByID(repository.Reminders, id)
//		return item, item != nil
//
//	case "tag":
//		repository.TagsMu.RLock()
//		defer repository.TagsMu.RUnlock()
//		item := findByID(repository.Tags, id)
//		return item, item != nil
//
//	case "time_entry":
//		repository.TimeEntriesMu.RLock()
//		defer repository.TimeEntriesMu.RUnlock()
//		item := findByID(repository.TimeEntries, id)
//		return item, item != nil
//	}
//	return nil, false
//}

//func updateItemWithLock(itemType string, id int, updated common.Item) bool {
//	switch itemType {
//	case "user":
//		repository.UsersMu.Lock()
//		defer repository.UsersMu.Unlock()
//		ok := updateByID(&repository.Users, id, updated.(*model.User))
//		if ok {
//			_ = repository.SaveAllItems("user")
//		}
//		return ok
//
//	case "project":
//		repository.ProjectsMu.Lock()
//		defer repository.ProjectsMu.Unlock()
//		ok := updateByID(&repository.Projects, id, updated.(*model.Project))
//		if ok {
//			_ = repository.SaveAllItems("project")
//		}
//		return ok
//
//	case "task":
//		repository.TasksMu.Lock()
//		defer repository.TasksMu.Unlock()
//		ok := updateByID(&repository.Tasks, id, updated.(*model.Task))
//		if ok {
//			_ = repository.SaveAllItems("task")
//		}
//		return ok
//
//	case "reminder":
//		repository.RemindersMu.Lock()
//		defer repository.RemindersMu.Unlock()
//		ok := updateByID(&repository.Reminders, id, updated.(*model.Reminder))
//		if ok {
//			_ = repository.SaveAllItems("reminder")
//		}
//		return ok
//
//	case "tag":
//		repository.TagsMu.Lock()
//		defer repository.TagsMu.Unlock()
//		ok := updateByID(&repository.Tags, id, updated.(*model.Tag))
//		if ok {
//			_ = repository.SaveAllItems("tag")
//		}
//		return ok
//
//	case "time_entry":
//		repository.TimeEntriesMu.Lock()
//		defer repository.TimeEntriesMu.Unlock()
//		ok := updateByID(&repository.TimeEntries, id, updated.(*model.TimeEntry))
//		if ok {
//			_ = repository.SaveAllItems("time_entry")
//		}
//		return ok
//	default:
//		return false
//	}
//}

//func updateByID[T any](slice *[]*T, id int, updated *T) bool {
//	for i, item := range *slice {
//		if iItem, ok := any(item).(common.Item); ok && int(iItem.GetItem()) == id {
//			(*slice)[i] = updated
//			return true
//		}
//	}
//	return false
//}

//func deleteItemWithLock(itemType string, id int) bool {
//	switch itemType {
//	case "user":
//		repository.UsersMu.Lock()
//		defer repository.UsersMu.Unlock()
//		ok := deleteByID(&repository.Users, id)
//		if ok {
//			_ = repository.SaveAllItems("user")
//		}
//		return ok
//
//	case "project":
//		repository.ProjectsMu.Lock()
//		defer repository.ProjectsMu.Unlock()
//		ok := deleteByID(&repository.Projects, id)
//		if ok {
//			_ = repository.SaveAllItems("project")
//		}
//		return ok
//
//	case "task":
//		repository.TasksMu.Lock()
//		defer repository.TasksMu.Unlock()
//		ok := deleteByID(&repository.Tasks, id)
//		if ok {
//			_ = repository.SaveAllItems("task")
//		}
//		return ok
//
//	case "reminder":
//		repository.RemindersMu.Lock()
//		defer repository.RemindersMu.Unlock()
//		ok := deleteByID(&repository.Reminders, id)
//		if ok {
//			_ = repository.SaveAllItems("reminder")
//		}
//		return ok
//
//	case "tag":
//		repository.TagsMu.Lock()
//		defer repository.TagsMu.Unlock()
//		ok := deleteByID(&repository.Tags, id)
//		if ok {
//			_ = repository.SaveAllItems("tag")
//		}
//		return ok
//
//	case "time_entry":
//		repository.TimeEntriesMu.Lock()
//		defer repository.TimeEntriesMu.Unlock()
//		ok := deleteByID(&repository.TimeEntries, id)
//		if ok {
//			_ = repository.SaveAllItems("time_entry")
//		}
//		return ok
//	default:
//		return false
//	}
//}

//func deleteByID[T any](slice *[]*T, id int) bool {
//	log.Println("Got DELETE request for id: ", id)
//
//	for i, item := range *slice {
//		if iItem, ok := any(item).(common.Item); ok && int(iItem.GetItem()) == id {
//			*slice = append((*slice)[:i], (*slice)[i+1:]...)
//			return true
//		}
//	}
//	return false
//}
