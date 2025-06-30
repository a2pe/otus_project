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

func ensureFileWithEmptyArray(path string) error {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte("[]"), 0644); err != nil {
			return fmt.Errorf("failed to create and write to file %s: %w", path, err)
		}
		return nil
	}

	if err != nil {
		return fmt.Errorf("failed to stat file %s: %w", path, err)
	}

	if info.Size() == 0 {
		if err := os.WriteFile(path, []byte("[]"), 0644); err != nil {
			return fmt.Errorf("failed to write empty array to file %s: %w", path, err)
		}
	}

	return nil
}

func LoadDataFromFile[T any](filePath string, data *[]*T, lastLen *int) error {
	if err := ensureFileWithEmptyArray(filePath); err != nil {
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

func SaveSliceToFile[T any](path string, mu *sync.RWMutex, slice []*T) error {
	mu.Lock()
	defer mu.Unlock()

	filePath := GetFinalFilePath(path)

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(slice); err != nil {
		return fmt.Errorf("failed to encode slice to file %s: %w", filePath, err)
	}
	return nil
}
