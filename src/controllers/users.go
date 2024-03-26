package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User

	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching all users\n"))
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching a user\n"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a user\n"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user\n"))
}
