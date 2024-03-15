package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "joshua468:170821002@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
}

func SaveLoginInfo(username, password string) (string, error) {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := db.Exec(query, username, password)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return string(id), nil
}

func GetSubmission(submissionID string) (interface{}, error) {

	var submission interface{}

	query := "SELECT * FROM submissions WHERE id = ?"
	err := db.QueryRow(query, submissionID).Scan(&submission)
	if err != nil {
		return nil, err
	}

	return submission, nil
}
