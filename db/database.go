package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	pgHost                                               = "localhost"
	pgPort                                               = 5432
	pgUser                                               = "root"
	pgPassword                                           = "root"
	pgDatabase                                           = "root"
	studentsTable, studentsNameColumn, studentsAgeColumn = "students", "name", "age"
)

var (
	DB  *sql.DB
	err error
)

func ConnectToPostgres() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", pgUser, pgPassword, pgHost, pgPort, pgDatabase)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	createTableIfNotExists()
}

func CloseConnection() {
	sqlDB := DB
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func createTableIfNotExists() {
	sqlStatement := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id SERIAL,
		%s TEXT,
		%s INT
	);`, studentsTable, studentsNameColumn, studentsAgeColumn)

	_, err = DB.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}
