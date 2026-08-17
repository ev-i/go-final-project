package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ev-i/go-final-project/pkg/db"
)

func makeDoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим запрос
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	if id == "" {
		respondWithError(w, 400, "id cannot be empty")
		return
	}
	// Достаём задачу из базы
	task, err := db.GetTask(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, 404, "Задача не найдена")
			return
		}
		respondWithError(w, 500, err.Error())
		return
	}
	// Проверяем является ли задача повторяемой
	isRepeatable := false
	if task.Repeat != "" {
		isRepeatable = true
	}

	if isRepeatable {
		nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			respondWithError(w, 500, err.Error())
			return
		}
		task.Date = nextDate
		err = db.UpdateTask(task)
		if err != nil {
			respondWithError(w, 500, err.Error())
			return
		}

	} else {
		err = db.DeleteTask(id)
		if err != nil {
			respondWithError(w, 500, err.Error())
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	if id == "" {
		respondWithError(w, 400, "id cannot be empty")
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		if errors.Is(err, db.ErrNothingToDelete) {
			respondWithError(w, 404, err.Error())
			return
		}
		respondWithError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	if id == "" {
		respondWithError(w, 400, "id cannot be empty")
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, 404, "Задача не найдена")
			return
		}
		respondWithError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func putTaskHandler(w http.ResponseWriter, r *http.Request) {

	task := db.Task{}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if task.Title == "" {
		respondWithError(w, 400, "Title cannot be empty")
		return
	}

	err := checkDate(&task)

	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	err = db.UpdateTask(&task)

	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
