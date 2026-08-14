package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ev-i/go-final-project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	task, err := db.GetTask(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, "Задача не найдена")
			return
		}
		respondWithError(w, err.Error())
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

	err := json.NewDecoder(r.Body).Decode(&task)

	if task.Title == "" {
		respondWithError(w, "Title cannot be empty")
		return
	}

	err = checkDate(&task)

	if err != nil {
		respondWithError(w, err.Error())
		return
	}

	err = db.UpdateTask(&task)

	if err != nil {
		respondWithError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
