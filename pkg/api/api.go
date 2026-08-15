package api

import (
	"encoding/json"
	"net/http"
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
	http.HandleFunc("GET /api/tasks", tasksHandler)
	http.HandleFunc("GET /api/task", getTaskHandler)
	http.HandleFunc("PUT /api/task", putTaskHandler)
	http.HandleFunc("DELETE /api/task", deleteTaskHandler)
	http.HandleFunc("POST /api/task/done", makeDoneTaskHandler)
}
