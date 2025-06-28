package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/log"
	"os"
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

	logger = log.NewLogfmtLogger(os.Stdout)

	lastUsersLen       int
	lastProjectsLen    int
	lastTasksLen       int
	lastRemindersLen   int
	lastTagsLen        int
	lastTimeEntriesLen int
)

func Init(ctx context.Context) error {
	if err := LoadAll(); err != nil {
		return fmt.Errorf("failed to load initial data: %w", err)
	}

	updateMaxIDs()

	StartSliceLogger(ctx, logger)
	return nil
}

func updateMaxIDs() {
	for _, u := range Users {
		if id := u.GetItem(); uint64(id) > userCounter {
			userCounter = uint64(id)
		}
	}
	for _, p := range Projects {
		if id := p.GetItem(); uint64(id) > projectCounter {
			projectCounter = uint64(id)
		}
	}
	for _, t := range Tasks {
		if id := t.GetItem(); uint64(id) > taskCounter {
			taskCounter = uint64(id)
		}
	}
	for _, r := range Reminders {
		if id := r.GetItem(); uint64(id) > reminderCounter {
			reminderCounter = uint64(id)
		}
	}
	for _, tag := range Tags {
		if id := tag.GetItem(); uint64(id) > tagCounter {
			tagCounter = uint64(id)
		}
	}
	for _, te := range TimeEntries {
		if id := te.GetItem(); uint64(id) > timeEntriesCounter {
			timeEntriesCounter = uint64(id)
		}
	}
}

func checkAndLogNewItems[T any](label string, mu *sync.RWMutex, slice *[]*T, lastLen *int) {
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
				checkAndLogNewItems("User", &UsersMu, &Users, &lastUsersLen)
				checkAndLogNewItems("Project", &ProjectsMu, &Projects, &lastProjectsLen)
				checkAndLogNewItems("Task", &TasksMu, &Tasks, &lastTasksLen)
				checkAndLogNewItems("Reminder", &RemindersMu, &Reminders, &lastRemindersLen)
				checkAndLogNewItems("Tag", &TagsMu, &Tags, &lastTagsLen)
				checkAndLogNewItems("TimeEntry", &TimeEntriesMu, &TimeEntries, &lastTimeEntriesLen)

			}
		}
	}()
}

func LoadAll() error {
	var (
		userPath        = data.GetFinalFilePath(DataRegistry["user"].FileName)
		projectPath     = data.GetFinalFilePath(DataRegistry["project"].FileName)
		tagsPath        = data.GetFinalFilePath(DataRegistry["tag"].FileName)
		tasksPath       = data.GetFinalFilePath(DataRegistry["task"].FileName)
		remindersPath   = data.GetFinalFilePath(DataRegistry["reminder"].FileName)
		timeEntriesPath = data.GetFinalFilePath(DataRegistry["time_entry"].FileName)
	)

	if err := data.LoadDataFromFile(userPath, &Users, &lastUsersLen); err != nil {
		return err

	}
	if err := data.LoadDataFromFile(projectPath, &Projects, &lastProjectsLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(tasksPath, &Tasks, &lastTasksLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(tagsPath, &Tags, &lastTagsLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(remindersPath, &Reminders, &lastRemindersLen); err != nil {
		return err
	}
	if err := data.LoadDataFromFile(timeEntriesPath, &TimeEntries, &lastTimeEntriesLen); err != nil {
		return err
	}
	return nil
}

func GetByID(itemType string, id int) (common.Item, bool) {
	reg, ok := DataRegistry[itemType]
	if !ok {
		return nil, false
	}

	reg.Mutex.RLock()
	defer reg.Mutex.RUnlock()

	switch data := reg.Data.(type) {
	case *[]*model.User:
		item := findItem(*data, id)
		return item, item != nil
	case *[]*model.Project:
		item := findItem(*data, id)
		return item, item != nil
	case *[]*model.Task:
		item := findItem(*data, id)
		return item, item != nil
	case *[]*model.Reminder:
		item := findItem(*data, id)
		return item, item != nil
	case *[]*model.Tag:
		item := findItem(*data, id)
		return item, item != nil
	case *[]*model.TimeEntry:
		item := findItem(*data, id)
		return item, item != nil
	default:
		return nil, false
	}
}

func findItem[T any](items []*T, id int) common.Item {
	for _, item := range items {
		if i, ok := any(item).(common.Item); ok && int(i.GetItem()) == id {
			return i
		}
	}
	return nil
}

func GetAllItems(itemType string) (any, error) {
	reg, ok := DataRegistry[itemType]
	if !ok {
		return nil, fmt.Errorf("unknown item type: %s", itemType)
	}

	reg.Mutex.RLock()
	defer reg.Mutex.RUnlock()

	switch data := reg.Data.(type) {
	case *[]*model.User:
		return data, nil
	case *[]*model.Project:
		return data, nil
	case *[]*model.Task:
		return data, nil
	case *[]*model.Reminder:
		return data, nil
	case *[]*model.Tag:
		return data, nil
	case *[]*model.TimeEntry:
		return data, nil
	default:
		return nil, fmt.Errorf("unsupported data type for: %s", itemType)
	}
}

func UpdateItem(itemType string, updated common.Item) bool {
	reg, ok := DataRegistry[itemType]
	if !ok {
		return false
	}

	reg.Mutex.Lock()
	defer reg.Mutex.Unlock()

	switch data := reg.Data.(type) {
	case *[]*model.User:
		return updateItemInSlice(data, updated)
	case *[]*model.Project:
		return updateItemInSlice(data, updated)
	case *[]*model.Task:
		return updateItemInSlice(data, updated)
	case *[]*model.Reminder:
		return updateItemInSlice(data, updated)
	case *[]*model.Tag:
		return updateItemInSlice(data, updated)
	case *[]*model.TimeEntry:
		return updateItemInSlice(data, updated)
	default:
		return false
	}
}

func updateItemInSlice[T any](slice *[]*T, updated common.Item) bool {
	for i, item := range *slice {
		if orig, ok := any(item).(common.Item); ok && orig.GetItem() == updated.GetItem() {
			if newVal, ok := any(updated).(*T); ok {
				(*slice)[i] = newVal
				return true
			}
		}
	}
	return false
}

func DeleteItem(itemType string, id int) bool {
	reg, ok := DataRegistry[itemType]
	if !ok {
		return false
	}

	reg.Mutex.Lock()
	defer reg.Mutex.Unlock()

	switch data := reg.Data.(type) {
	case *[]*model.User:
		return deleteItemFromSlice(data, id)
	case *[]*model.Project:
		return deleteItemFromSlice(data, id)
	case *[]*model.Task:
		return deleteItemFromSlice(data, id)
	case *[]*model.Reminder:
		return deleteItemFromSlice(data, id)
	case *[]*model.Tag:
		return deleteItemFromSlice(data, id)
	case *[]*model.TimeEntry:
		return deleteItemFromSlice(data, id)
	default:
		return false
	}
}

func deleteItemFromSlice[T any](slice *[]*T, id int) bool {
	for i, item := range *slice {
		if v, ok := any(item).(common.Item); ok && int(v.GetItem()) == id {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return true
		}
	}
	return false
}

func appendWithLock[T any](mu *sync.RWMutex, slice *[]*T, item *T) {
	mu.Lock()
	defer mu.Unlock()
	*slice = append(*slice, item)
}

func prepareNewItem[T common.Item](item T, counter *uint64) {
	item.SetID(uint(atomic.AddUint64(counter, 1)))
	item.SetCreatedAt(time.Now())
}

func SaveItem(item common.Item) error {
	switch v := item.(type) {
	case *model.User:
		prepareNewItem(v, &userCounter)
		appendWithLock(&UsersMu, &Users, v)
		return SaveAllItems("user")
	case *model.Project:
		prepareNewItem(v, &projectCounter)
		appendWithLock(&ProjectsMu, &Projects, v)
		return SaveAllItems("project")
	case *model.Task:
		prepareNewItem(v, &taskCounter)
		appendWithLock(&TasksMu, &Tasks, v)
		return SaveAllItems("task")
	case *model.Reminder:
		prepareNewItem(v, &reminderCounter)
		appendWithLock(&RemindersMu, &Reminders, v)
		return SaveAllItems("reminder")
	case *model.Tag:
		prepareNewItem(v, &tagCounter)
		appendWithLock(&TagsMu, &Tags, v)
		return SaveAllItems("tag")
	case *model.TimeEntry:
		prepareNewItem(v, &timeEntriesCounter)
		appendWithLock(&TimeEntriesMu, &TimeEntries, v)
		return SaveAllItems("time_entry")
	default:
		return errors.New("unsupported type in SaveItem")

	}
}

func SaveAllItems(itemType string) error {
	reg, ok := DataRegistry[itemType]
	if !ok {
		return fmt.Errorf("unknown item type: %s", itemType)
	}

	fmt.Printf("Saving all items of type %s to %s\n", itemType, reg.FileName)

	switch info := reg.Data.(type) {
	case *[]*model.User:
		return data.SaveSliceToFile(reg.FileName, reg.Mutex, *info)
	case *[]*model.Project:
		return data.SaveSliceToFile(reg.FileName, reg.Mutex, *info)
	case *[]*model.Task:
		return data.SaveSliceToFile(reg.FileName, reg.Mutex, *info)
	case *[]*model.Reminder:
		return data.SaveSliceToFile(reg.FileName, reg.Mutex, *info)
	case *[]*model.Tag:
		return data.SaveSliceToFile(reg.FileName, reg.Mutex, *info)
	case *[]*model.TimeEntry:
		return data.SaveSliceToFile(reg.FileName, reg.Mutex, *info)
	default:
		return fmt.Errorf("unsupported data type for: %s", itemType)
	}
}
