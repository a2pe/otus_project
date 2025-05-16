package service

import (
	"fmt"
	"math/rand"
	"otus_project/internal/model"
	"otus_project/internal/repository"
	"time"
)

func GenerateData() {
	for i := 0; i < 10; i++ {
		num := uint(rand.Intn(101))
		user := model.User{ID: num, Name: fmt.Sprintf("user-%d", num), Timezone: "Asia/Shanghai", CreatedAt: time.Now()}
		user.SetEmail(fmt.Sprintf("new_email-%d", num))
		user.SetPassword("123456")
		project := model.Project{ID: num, UserID: user.ID, Name: fmt.Sprintf("project-%d", num), CreatedAt: time.Now(), Description: fmt.Sprintf("project-%d", num)}
		task := model.Task{ID: num, ProjectID: project.ID, Title: fmt.Sprintf("new task %d", num), Status: "new", CreatedAt: time.Now(), Description: fmt.Sprintf("task-%d", num), DueDate: time.Now().Add(time.Hour * 24 * 10)}
		reminder := model.Reminder{ID: num, UserID: user.ID, TaskID: task.ID, RemindAt: time.Now().Add(time.Hour * 24 * 10), CreatedAt: time.Now()}
		tag := model.Tag{ID: num, UserID: user.ID, Name: fmt.Sprintf("tag-%d", num)}
		timeEntry := model.TimeEntry{ID: num, UserID: user.ID, TaskID: task.ID, StartTime: time.Now(), CreatedAt: time.Now()}

		_ = repository.SaveItem(user)
		_ = repository.SaveItem(task)
		_ = repository.SaveItem(project)
		_ = repository.SaveItem(reminder)
		_ = repository.SaveItem(tag)
		_ = repository.SaveItem(timeEntry)

		time.Sleep(5 * time.Second)
	}
}
