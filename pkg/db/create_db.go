package db

import (
	"database/sql"
	"errors"
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const updateFile = "schemas/001_update.sql"

func CreateDB(path string) error {
	sqlString, err := readSQLFile()
	if err != nil {
		return err
	}

	database, err := sql.Open("sqlite3", path)

	if err != nil {
		return errors.New("db dont open: " + err.Error())
	}

	statment, err := database.Prepare(sqlString)

	if err != nil {
		return errors.New("db dont connect: " + err.Error())
	}

	statment.Exec()

	return nil
}

func readSQLFile() (string, error) {
	file, err := os.Open(updateFile)
	if err != nil {
		return "", errors.New("sql file not exist: " + err.Error())
	}

	defer file.Close()

	data := make([]byte, 64)

	var value string
	for {
		r, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		value += string(data[:r])
	}
	return value, nil
}
