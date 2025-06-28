package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"otus_project/internal/repository"
	"strconv"
)

var (
	ErrInvalidID   = errors.New("invalid ID format")
	ErrInvalidJSON = errors.New("invalid JSON body")
	ErrUnknownType = errors.New("unknown item type")
	ErrInternal    = errors.New("internal server error")
)

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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode("Data successfully added! :)"); err != nil {
			http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		}
	}
}

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
