package repositories

import (
	"database/sql"
	"fmt"

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

func (repository users) Search(nameOrNickname string) ([]models.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname) // %%nameOrNickname%%

	lines, error := repository.db.Query("select id, name, nickname, email, createdAt from users where name LIKE ? or nickname LIKE ?",
		nameOrNickname,
		nameOrNickname,
	)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if error = lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchByID(ID uint64) (models.User, error) {
	lines, error := repository.db.Query("select id, name, nickname, email, createdAt from users where id = ?", ID)
	if error != nil {
		return models.User{}, error
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if error = lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

func (repository users) Update(ID uint64, user models.User) error {
	statement, error := repository.db.Prepare("update users set name = ?, nickname = ?, email = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nickname, user.Email, ID); error != nil {
		return error
	}

	return nil
}
