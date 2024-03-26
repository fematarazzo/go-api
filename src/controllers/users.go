package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created\n"))
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
