package models

import "github.com/kelaasor-quiz/db"

type Question struct {
	Id           int64  `json:"id" `
	QuizId       int64  `json:"quiz_id"`
	QuestionText string `json:"question_text" binding:"required" `
}

func GetQuestionByQuizId(quizId int64) (*[]Question, error) {
	questions := []Question{}
	query := `SELECT * FROM questions WHERE quiz_id = ?`
	rows, err := db.DB.Query(query, quizId)

	if err != nil {
		return nil, err

	}

	for rows.Next() {
		tempQuestion := Question{}

		err := rows.Scan(&tempQuestion.Id, &tempQuestion.QuizId, &tempQuestion.QuestionText)

		if err != nil {
			return nil, err
		}

		questions = append(questions, tempQuestion)
	}

	return &questions, nil

}

func GetQuestionById(questionId int64) (*Question, error) {
	query := `SELECT * FROM questions WHERE id = ?`
	question := Question{}
	err := db.DB.QueryRow(query, questionId).Scan(&question.Id, &question.QuizId, &question.QuestionText)

	if err != nil {
		return nil, err
	}

	return &question, nil

}

func (q *Question) AddQuestionToQuiz() (int64, error) {
	query := `INSERT INTO questions (quiz_id ,question_text) VALUES (?,?) `

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(q.QuizId, q.QuestionText)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	return id, err

}
