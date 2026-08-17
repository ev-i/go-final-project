package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ev-i/go-final-project/pkg/db"
)

type Response struct {
	ID int `json:"id"`
}

func addTaskHandler(res http.ResponseWriter, req *http.Request) {

	var task db.Task

	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(res, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		respondWithError(res, 400, "Title cannot be empty")
		return
	}

	err = checkDate(&task)

	if err != nil {
		respondWithError(res, 500, err.Error())
		return
	}

	id, err := db.AddTask(&task)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(res).Encode(Response{ID: int(id)}); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func checkDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(layout)
	}
	t, err := time.Parse(layout, task.Date)
	if err != nil {
		return err
	}
	next := now.Format(layout)
	if task.Repeat != "" && task.Date != now.Format(layout) {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}
	if !afterNow(now, t) {
		if len(task.Repeat) == 0 {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format(layout)
		} else {
			// в противном случае, берём вычисленную ранее следующую дату
			task.Date = next
		}
	}
	return nil
}
