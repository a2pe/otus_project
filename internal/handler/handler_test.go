package handler_test

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"otus_project/internal/handler"
	"otus_project/internal/model"
	"otus_project/internal/repository"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Task Test OK: All Valid Params
func TestCreateTaskHandler_Success(t *testing.T) {
	body := `{
		"title": "Test Task",
		"description": "This is a test",
		"status": "new",
		"due_date": "2025-06-30T10:00:00Z",
		"estimate_hrs": 1.5
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.CreateItemHandler("task").ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)
	require.Contains(t, rr.Body.String(), "successfully")
}

// Task Test OK: Minimum Valid Params
func TestCreateTaskHandler_MinFields_Success(t *testing.T) {
	body := `{
		"title": "Test Task",
		"status": "in_progress",
		"due_date": "2025-07-01T15:04:05Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.CreateItemHandler("task").ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)
	require.Contains(t, rr.Body.String(), "Data successfully added")
}

// Task Test Fail: One Invalid Param
func TestCreateTaskHandler_NotValidFields_Fail(t *testing.T) {
	body := `{
		"some": "Test Task"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.CreateItemHandler("task").ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Contains(t, rr.Body.String(), "validation failed: "+
		"Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag\n"+
		"Key: 'Task.Status' Error:Field validation for 'Status' failed on the 'required' tag\n"+
		"Key: 'Task.DueDate' Error:Field validation for 'DueDate' failed on the 'required' tag\n")
}

// Task Test Fail: Invalid JSON
func TestCreateTaskHandler_InvalidJSON_Fail(t *testing.T) {
	body := `{bad json}`

	req := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.CreateItemHandler("task").ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Contains(t, rr.Body.String(), "invalid JSON")
}

// Task Test Fail: Empty JSON
func TestCreateItemHandler_EmptyBody_Fail(t *testing.T) {
	body := `{}`

	req := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.CreateItemHandler("task").ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Contains(t, rr.Body.String(), "validation failed: "+
		"Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag\n"+
		"Key: 'Task.Status' Error:Field validation for 'Status' failed on the 'required' tag\n"+
		"Key: 'Task.DueDate' Error:Field validation for 'DueDate' failed on the 'required' tag\n")
}

// Task Test Fail: Unknown Item Type
func TestCreateItemHandler_UnknownType_Fail(t *testing.T) {
	body := `{}`

	req := httptest.NewRequest(http.MethodPost, "/api/unknown", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.CreateItemHandler("unknown").ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Contains(t, rr.Body.String(), "unknown item type")
}

// Task Test Fail: Wrong Content Type
func TestCreateItemHandler_WrongContentType_Fail(t *testing.T) {
	body := `some=some`

	req := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler.CreateItemHandler("task").ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Contains(t, rr.Body.String(), "invalid JSON body")
}

func TestGetItemByIDHandler_Success(t *testing.T) {
	dueDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	item := &model.Task{Title: "demo", Status: "new", DueDate: dueDate}
	err := repository.SaveItem(item)
	require.NoError(t, err)

	r := chi.NewRouter()
	r.Get("/api/task/{id}", handler.GetItemByIDHandler("task"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/task/%d", item.ID), nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var result model.Task
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &result))
	require.Equal(t, item.ID, result.ID)
}

func TestGetItemByIDHandler_NotFound(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/api/task/{id}", handler.GetItemByIDHandler("task"))
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/task/%d", 1000000000), nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGetAllHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/task", nil)
	rr := httptest.NewRecorder()
	handler.GetAllHandler("task").ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var tasks []model.Task
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &tasks))
}

func TestUpdateItemHandler_Success(t *testing.T) {
	dueDate := time.Date(2025, time.September, 01, 0, 00, 1, 0, time.UTC)
	task := &model.Task{Title: "Old Title", Status: "new", DueDate: dueDate}
	err := repository.SaveItem(task)
	require.NoError(t, err)

	body := `{"title": "Updated Title", "status": "in_progress", "due_date": "2025-07-01T15:04:05Z"}`

	r := chi.NewRouter()
	r.Put("/api/task/{id}", handler.UpdateItemHandler("task"))
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/task/%d", task.ID), strings.NewReader(body))
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateItemHandler_InvalidID(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/api/task/abc", strings.NewReader(`{}`))
	rr := httptest.NewRecorder()
	handler.UpdateItemHandler("task").ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteItemHandler_Success(t *testing.T) {
	dueDate := time.Date(2025, time.July, 26, 10, 30, 0, 0, time.UTC)
	task := &model.Task{Title: "To delete", Status: "new", DueDate: dueDate}
	err := repository.SaveItem(task)
	require.NoError(t, err)

	r := chi.NewRouter()
	r.Delete("/api/task/{id}", handler.DeleteItemHandler("task"))

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/task/%d", task.ID), nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteItemHandler_NotFound(t *testing.T) {
	r := chi.NewRouter()
	r.Delete("/api/task/{id}", handler.DeleteItemHandler("task"))

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/task/%d", 10000), nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
}
