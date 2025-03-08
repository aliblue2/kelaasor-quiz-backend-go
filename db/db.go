package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func DatabaseConnection() {
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("database connection is not established.!")
		panic(err)
	}

	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTablesDB()

	if err != nil {
		fmt.Println("failed to create tables")
		panic(err)
	}

	fmt.Println("database successfully created and connection established.!")
}

func createTablesDB() error {
	usersTableQuery := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    phone VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := DB.Exec(usersTableQuery)

	if err != nil {
		return err
	}

	quizzesTableQuery := `CREATE TABLE  IF NOT EXISTS quizzes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(100) NOT NULL,
    description TEXT,
	state INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

	_, err = DB.Exec(quizzesTableQuery)

	if err != nil {
		return err
	}

	questionsTableQuery := `CREATE TABLE IF NOT EXISTS questions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    quiz_id INTEGER,
    question_text TEXT NOT NULL,
    FOREIGN KEY (quiz_id) REFERENCES Quizzes(id) ON DELETE CASCADE
);`

	_, err = DB.Exec(questionsTableQuery)

	if err != nil {
		return err
	}

	answersTableQuery := `CREATE TABLE IF NOT EXISTS answers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question_id INTEGER,
    answer_text TEXT NOT NULL,
    personality_type VARCHAR(50) NOT NULL,
    FOREIGN KEY (question_id) REFERENCES Questions(id) ON DELETE CASCADE
);`

	_, err = DB.Exec(answersTableQuery)

	if err != nil {
		return err
	}

	userResponseTableQuery := `CREATE TABLE IF NOT EXISTS users_response (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    question_id INTEGER,
    answer_id INTEGER,
    response_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (question_id) REFERENCES Questions(id) ON DELETE CASCADE,
    FOREIGN KEY (answer_id) REFERENCES Answers(id) ON DELETE CASCADE
);`

	_, err = DB.Exec(userResponseTableQuery)

	if err != nil {
		return err
	}
	userResultTableQuery := `CREATE TABLE IF NOT EXISTS results (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    quiz_id INTEGER,
    personality_type VARCHAR(50) NOT NULL,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (quiz_id) REFERENCES Quizzes(id) ON DELETE CASCADE
);`

	_, err = DB.Exec(userResultTableQuery)

	if err != nil {
		return err
	}

	return nil

}
