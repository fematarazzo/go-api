package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"api/src/database"
	"api/src/models"
	"api/src/repositories"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := io.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User

	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	userID, error := repository.Create(user)
	if error != nil {
		log.Fatal(error)
	}

	w.Write([]byte(fmt.Sprintf("ID inserted: %d\n", userID)))
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
