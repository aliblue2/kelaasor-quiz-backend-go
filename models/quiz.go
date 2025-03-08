package models

import (
	"time"

	"github.com/kelaasor-quiz/db"
)

type Quiz struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	State       int       `json:"state"`
	CreatedAt   time.Time `json:"created_at"`
}

func GetAllQuizzes() (*[]Quiz, error) {
	quizzes := []Quiz{}

	query := `SELECT * FROM quizzes`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		quiz := Quiz{}
		err := rows.Scan(&quiz.Id, &quiz.Title, &quiz.Description, &quiz.State, &quiz.CreatedAt)

		if err != nil {
			return nil, err
		}

		quizzes = append(quizzes, quiz)
	}

	return &quizzes, nil
}

func GetQuizWithId(id int64) (*Quiz, error) {
	quiz := Quiz{}
	query := `SELECT * FROM quizzes WHERE id = ?`

	err := db.DB.QueryRow(query, id).Scan(&quiz.Id, &quiz.Title, &quiz.Description, &quiz.State, &quiz.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &quiz, nil

}

func (q *Quiz) AddNewQuiz() (int64, error) {

	query := `INSERT INTO quizzes (title, description, state , created_at) VALUES (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, err
	}

	resutl, err := stmt.Exec(q.Title, q.Description, 1, q.CreatedAt)

	if err != nil {
		return -1, err
	}

	quizId, err := resutl.LastInsertId()

	return quizId, err
}
