package models

import "github.com/kelaasor-quiz/db"

type Answer struct {
	Id              int64  `json:"id" `
	QuestionId      int64  `json:"question_id" `
	AnswerText      string `json:"answer_text" binding:"required"`
	PersonalityType string `json:"personality_type" binding:"required"`
}

func (a *Answer) AddNewAnswerToQuestions() (int64, error) {
	query := `INSERT INTO answers (question_id,answer_text,personality_type) VALUES (?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(a.QuestionId, a.AnswerText, a.PersonalityType)

	if err != nil {
		return -1, err
	}

	answerId, err := result.LastInsertId()

	return answerId, err

}

func GetAllAnswersByQuestionId(questionId int64) (*[]Answer, error) {
	answers := []Answer{}
	query := `SELECT * FROM answers WHERE question_id = ?`

	rows, err := db.DB.Query(query, questionId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tempAnswer := Answer{}
		err := rows.Scan(&tempAnswer.Id, &tempAnswer.QuestionId, &tempAnswer.AnswerText, &tempAnswer.PersonalityType)

		if err != nil {
			return nil, err
		}

		answers = append(answers, tempAnswer)
	}

	return &answers, nil

}
