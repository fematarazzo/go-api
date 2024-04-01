package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	userInDB, error := repository.SearchByEmail(user.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.CheckPassword(userInDB.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, error := authentication.GenerateToken(userInDB.ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	userID := strconv.FormatUint(userInDB.ID, 10)

	responses.JSON(w, http.StatusOK, models.DataAuthentication{ID: userID, Token: token})
}
