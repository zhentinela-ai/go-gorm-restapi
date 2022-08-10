package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhentinela-ai/go-gorm-restapi/db"
	"github.com/zhentinela-ai/go-gorm-restapi/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) // 200
	w.Write([]byte("Task deleted"))
}
