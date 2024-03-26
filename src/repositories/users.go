package repositories

import (
	"database/sql"

	"api/src/models"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare(
		"insert into users (name, nickname, email, password) values (?, ?, ?, ?)",
	)

	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)

	if error != nil {
		return 0, error
	}

	lastInsertedID, error := result.LastInsertId()

	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedID), nil
}
