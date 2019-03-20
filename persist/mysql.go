package persist

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
        "fmt"
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

func InitDb() {
	db, err := sql.Open("mysql", "root:password@tcp(tourney-mysql:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connection created successfully")
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS tourney_db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database.")
	}

	_, err = db.Exec("USE tourney_db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully.")
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS players ( id bigint(20) unsigned not null auto_increment, name varchar(20) not null, knickname varchar(20), primary key (id) )")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully.")
	}

	defer db.Close()
}
