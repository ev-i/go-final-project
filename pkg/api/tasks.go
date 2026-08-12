package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ev-i/go-final-project/pkg/db"
)

type TaskResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		respondWithError(w, err.Error())
		return
	}

	response := TaskResp{Tasks: tasks}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
