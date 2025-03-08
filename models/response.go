package models

import (
	"time"

	"github.com/kelaasor-quiz/db"
)

type Response struct {
	Id           int64     `json:"id" `
	UserId       int64     `json:"user_id" `
	QuestionId   int64     `json:"question_id" binding:"required"`
	AnswerId     int64     `json:"answer_id" binding:"required"`
	ResponseTime time.Time `json:"response_time" `
}

func (r *Response) Submit() (int64, error) {
	query := `INSERT INTO users_response (user_id , question_id , answer_id , response_time) VALUES (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(r.UserId, r.QuestionId, r.AnswerId, r.ResponseTime)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	return id, err
}
