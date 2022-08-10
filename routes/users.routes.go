package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhentinela-ai/go-gorm-restapi/db"
	"github.com/zhentinela-ai/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
		return
	}

	// db.DB.Delete(&user)
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK) // 200
}
