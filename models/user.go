package models

import (
	"errors"
	"time"

	"github.com/kelaasor-quiz/db"
	"github.com/kelaasor-quiz/utils"
)

type User struct {
	Id        int64     `json:"id" `
	Phone     string    `json:"phone" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at" `
}

func (u *User) Signup() (int64, error) {
	query := `INSERT INTO users (phone , password_hash, created_at) VALUES (?,?,?)`
	hahedPassword, err := utils.GenerateHashPassword(u.Password)

	if err != nil {
		return -1, err
	}

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, err
	}

	resutl, err := stmt.Exec(u.Phone, hahedPassword, time.Now())

	if err != nil {
		return -1, err
	}

	userId, err := resutl.LastInsertId()

	return userId, err
}

func ValidateUserCreadentials(phone, password string) (string, error) {
	user := User{}
	query := `SELECT * FROM users WHERE phone = ?`

	err := db.DB.QueryRow(query, phone).Scan(&user.Id, &user.Phone, &user.Password, &user.CreatedAt)

	if err != nil {
		return "", err
	}

	validPass := utils.ComparePassword(password, user.Password)

	if !validPass {
		return "", errors.New("password is invalid")
	}

	accessToken, err := utils.GenerateToken(phone, user.Password, user.Id)

	if err != nil {
		return "", errors.New("cant generate new token")

	}

	return accessToken, nil

}
