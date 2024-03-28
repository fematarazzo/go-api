package models

import (
	"errors"
	"net/mail"
	"strings"
	"time"

	"api/src/security"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	if error := user.format(step); error != nil {
		return error
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name is mandatory and cannot be empty")
	}

	if user.Nickname == "" {
		return errors.New("Nickname is mandatory and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("Email is mandatory and cannot be empty")
	}

	_, error := mail.ParseAddress(user.Email)
	if error != nil {
		return errors.New("Invalid email")
	}

	if user.Password == "" && step == "registration" {
		return errors.New("Password is mandatory and cannot be empty")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if step == "registration" {
		passwordHash, error := security.Hash(user.Password)
		if error != nil {
			return error
		}

		user.Password = string(passwordHash)
	}

	return nil
}
