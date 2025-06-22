package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	UsersFile       = "users.json"
	ProjectsFile    = "projects.json"
	RemindersFile   = "reminders.json"
	TagsFile        = "tags.json"
	TasksFile       = "tasks.json"
	TimeEntriesFile = "time_entries.json"
)

func GetFinalFilePath(file string) string {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(b)
	finalPath := filepath.Join(projectRoot, "files", file)
	return finalPath
}

func ensureFileExists(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
		defer f.Close()
	}
	return nil
}

func LoadDataFromFile[T any](filePath string, mu *sync.Mutex, data *[]*T, lastLen *int) error {
	if err := ensureFileExists(filePath); err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer f.Close()

	var loaded []*T
	if err := json.NewDecoder(f).Decode(&loaded); err != nil {
		return fmt.Errorf("failed to decode file %s: %w", filePath, err)
	}

	mu.Lock()
	defer mu.Unlock()

	*data = loaded
	*lastLen = len(loaded)
	return nil
}

func AppendToFile[T any](path string, item *T) error {
	filePath := GetFinalFilePath(path)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to append to file %s: %w", filePath, err)
	}
	defer f.Close()

	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal values to file %s: %w", filePath, err)
	}

	_, err = f.Write(append(data, ',', '\n'))
	if err != nil {
		return fmt.Errorf("failed to append to file %s: %w", filePath, err)
	}
	return err
}
