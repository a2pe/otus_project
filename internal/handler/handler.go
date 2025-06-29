package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"otus_project/internal/model"
	"otus_project/internal/notification"
	"otus_project/internal/repository"
	"strconv"
	"time"
)

var (
	ErrInvalidID   = errors.New("invalid ID format")
	ErrInvalidJSON = errors.New("invalid JSON body")
	ErrUnknownType = errors.New("unknown item type")
	ErrInternal    = errors.New("internal server error")
)

// GetItemByIDHandler godoc
// @Summary Получить сущность по ID
// @Description Возвращает сущность по типу и ID
// @Tags items
// @Produce json
// @Param type path string true "Тип сущности"
// @Param id path int true "ID сущности"
// @Success 200 {object} interface{}
// @Failure 400,404 {string} string
// @Router /api/{type}/{id} [get]
// @Security BearerAuth
func GetItemByIDHandler(itemType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}

		item, ok := repository.GetByID(itemType, id)
		if !ok {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
		}
	}
}

// GetAllHandler godoc
// @Summary Получить список сущностей
// @Description Возвращает список сущностей по типу: user, project, task, reminder, tag, time_entry
// @Tags items
// @Produce json
// @Param type path string true "Тип сущности (user, project, task...)"
// @Success 200 {array} interface{}
// @Failure 400 {string} string "unknown item type"
// @Router /api/{type} [get]
// @Security BearerAuth
func GetAllHandler(itemType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := repository.GetAllItems(itemType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		}
	}
}

// CreateItemHandler godoc
// @Summary Создание сущности
// @Description Создаёт сущность указанного типа: user, project, task, reminder, tag, time_entry
// @Tags items
// @Accept json
// @Produce json
// @Param item body interface{} true "Любая модель: user/project/task/etc"
// @Success 201
// @Failure 400 {string} string "invalid json"
// @Failure 500 {string} string "failed to save item"
// @Router /api/{type} [post]
// @Security BearerAuth
func CreateItemHandler(itemType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item := getEmptyItem(itemType)
		if item == nil {
			http.Error(w, ErrUnknownType.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(item); err != nil {
			http.Error(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
			return
		}

		if err := repository.SaveItem(item); err != nil {
			http.Error(w, "Failed to save item", http.StatusInternalServerError)
			return
		}

		if itemType == "task" {
			raw := repository.DataRegistry["task"].Data
			taskSlicePtr, ok := raw.(*[]*model.Task)
			if ok && len(*taskSlicePtr) > 0 {
				thisTask := (*taskSlicePtr)[len(*taskSlicePtr)-1]

				reminderOffsets := []time.Duration{
					-24 * time.Hour,
					-8 * time.Hour,
					-1 * time.Hour,
				}
				messages := []string{
					"Reminder: task \"" + thisTask.Title + "\" is due in 24 hours!",
					"Reminder: task \"" + thisTask.Title + "\" is due in 8 hours!",
					"Reminder: task \"" + thisTask.Title + "\" is due in 1 hour!",
				}

				for i, offset := range reminderOffsets {
					reminderTime := thisTask.DueDate.Add(offset).Format(time.RFC3339)
					msg := messages[i]

					go func(remindAt, message string) {
						err := notification.ScheduleReminder(uint32(thisTask.ID), remindAt, message)
						if err != nil {
							fmt.Printf("Failed to schedule reminder: %v\n", err)
						}
					}(reminderTime, msg)
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode("Data successfully added! :)"); err != nil {
			http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		}
	}
}

// UpdateItemHandler godoc
// @Summary Обновить сущность
// @Description Обновляет сущность по ID
// @Tags items
// @Accept json
// @Produce json
// @Param type path string true "Тип сущности"
// @Param id path int true "ID сущности"
// @Param item body interface{} true "Обновлённая сущность"
// @Success 200
// @Failure 400,404,500 {string} string
// @Router /api/{type}/{id} [put]
// @Security BearerAuth
func UpdateItemHandler(itemType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}

		item := getEmptyItem(itemType)
		if item == nil {
			http.Error(w, ErrUnknownType.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(item); err != nil {
			http.Error(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
			return
		}

		item.SetID(uint(id))

		if !repository.UpdateItem(itemType, item) {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteItemHandler godoc
// @Summary Удалить сущность
// @Description Удаляет сущность по ID и типу
// @Tags items
// @Param type path string true "Тип сущности"
// @Param id path int true "ID сущности"
// @Success 204
// @Failure 400,404 {string} string
// @Router /api/{type}/{id} [delete]
// @Security BearerAuth
func DeleteItemHandler(itemType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}

		if !repository.DeleteItem(itemType, id) {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
