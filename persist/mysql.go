package persist

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(tourney-mysql:3306)/tourney_db")

	if err != nil {
		panic(err)
	}
	return db
}

func QueryDatabase(statement string) *sql.Rows {
	db := dbConnection()
	defer db.Close()

	rows, err := db.Query(statement)

	if err != nil {
		panic(err)
	}

	return rows
}

func InsertDatabase(statement string) {
	rows := QueryDatabase(statement)
	rows.Close()
}
