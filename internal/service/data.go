package service

import (
	"fmt"
	"log"
	"math/rand"
	"otus_project/internal/model"
	"otus_project/internal/model/common"
	"otus_project/internal/repository"
	"sync"
	"time"
)

func GenerateData() {
	dataChan := make(chan common.Item, 100) // buffered to avoid blocking
	var wg sync.WaitGroup

	// data generator to run for 2 minutes
	wg.Add(1)
	go func() {
		defer wg.Done()
		timeout := time.After(2 * time.Minute)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-timeout:
				close(dataChan)
				return
			case <-ticker.C:
				num := uint(rand.Intn(1001))
				user := model.User{
					ID:        num,
					Name:      fmt.Sprintf("user-%d", num),
					Timezone:  "Asia/Shanghai",
					CreatedAt: time.Now(),
				}
				user.SetEmail(fmt.Sprintf("new_email-%d", num))
				user.SetPassword("123456")
				dataChan <- user

				project := model.Project{
					ID:          num,
					UserID:      user.ID,
					Name:        fmt.Sprintf("project-%d", num),
					CreatedAt:   time.Now(),
					Description: fmt.Sprintf("project-%d", num),
				}
				dataChan <- project

				task := model.Task{
					ID:          num,
					ProjectID:   project.ID,
					Title:       fmt.Sprintf("new task %d", num),
					Status:      "new",
					CreatedAt:   time.Now(),
					Description: fmt.Sprintf("task-%d", num),
					DueDate:     time.Now().Add(24 * time.Hour * 10),
				}
				dataChan <- task

				reminder := model.Reminder{
					ID:        num,
					UserID:    user.ID,
					TaskID:    task.ID,
					RemindAt:  time.Now().Add(24 * time.Hour * 10),
					CreatedAt: time.Now(),
				}
				dataChan <- reminder

				tag := model.Tag{
					ID:     num,
					UserID: user.ID,
					Name:   fmt.Sprintf("tag-%d", num),
				}
				dataChan <- tag

				timeEntry := model.TimeEntry{
					ID:        num,
					UserID:    user.ID,
					TaskID:    task.ID,
					StartTime: time.Now(),
					CreatedAt: time.Now(),
				}
				dataChan <- timeEntry
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for item := range dataChan {
			if err := repository.SaveItem(item); err != nil {
				log.Println(err)
			}
		}
	}()

	wg.Wait()
}
