package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func respondWithError(res http.ResponseWriter, code int, message string) {
	res.Header().Set("Content-Type", "application/json")
	switch code {
	case 400:
		res.WriteHeader(http.StatusBadRequest)
	case 404:
		res.WriteHeader(http.StatusNotFound)
	case 500:
		res.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(res).Encode(ErrorResponse{Error: message})
}

func Init() {
	http.HandleFunc("GET /api/nextdate", nextDayHandler)
	http.HandleFunc("POST /api/task", addTaskHandler)
	http.HandleFunc("GET /api/tasks", tasksHandler)
	http.HandleFunc("GET /api/task", getTaskHandler)
	http.HandleFunc("PUT /api/task", putTaskHandler)
	http.HandleFunc("DELETE /api/task", deleteTaskHandler)
	http.HandleFunc("POST /api/task/done", makeDoneTaskHandler)
}
