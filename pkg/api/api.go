package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func respondWithError(res http.ResponseWriter, message string) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(res).Encode(ErrorResponse{Error: message})
}

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("POST /api/task", addTaskHandler)
}

func nextDayHandler(res http.ResponseWriter, req *http.Request) {
	now := req.FormValue("now")
	n, err := time.Parse(layout, now)
	if err != nil {
		http.Error(res, "Invalid \"now\" format", http.StatusBadRequest)
		return

	}
	date := req.FormValue("date")
	if date == "" {
		http.Error(res, "date cannot be empty", http.StatusBadRequest)
		return

	}

	repeat := req.FormValue("repeat")
	if repeat == "" {
		http.Error(res, "repeat cannot be empty", http.StatusBadRequest)
		return

	}

	nextDate, err := NextDate(n, date, repeat)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(nextDate))

}
