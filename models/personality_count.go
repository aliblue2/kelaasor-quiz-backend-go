package models

import (
	"database/sql"
	"math/rand"

	"github.com/kelaasor-quiz/db"
)

type PersonalityType struct {
	Type  string
	Count int
}

func GetUserResult(userId int64) (string, error) {
	query := `SELECT a.personality_type FROM users_response ur JOIN answers a ON ur.answer_id = a.id WHERE ur.user_id = ?`

	rows, err := db.DB.Query(query, userId)

	if err != nil {
		return "", err
	}

	defer rows.Close()
	personalityCounts := make(map[string]int)

	for rows.Next() {
		var personType string
		err := rows.Scan(&personType)

		if err != nil {
			return "", err
		}
		personalityCounts[personType]++
	}

	if len(personalityCounts) == 0 {
		return "", sql.ErrNoRows
	}

	var maxCount int
	var topTypes []string

	for personality, count := range personalityCounts {
		if count > maxCount {
			maxCount = count
			topTypes = []string{personality}
		} else if count == maxCount {
			topTypes = append(topTypes, personality)
		}
	}

	selectedType := topTypes[rand.Intn(len(topTypes))]

	return selectedType, nil

}
