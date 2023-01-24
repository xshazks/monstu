package helper

import (
	"database/sql"
	"fmt"

	"boilerplate/config"

	_ "github.com/go-sql-driver/mysql"
)

func MariaConnect(dbname string) (db *sql.DB) {
	db, err := sql.Open("mysql", config.MariaStringAkademik+dbname)
	if err != nil {
		fmt.Printf("MariaConnect: %v\n", err)
	}
	return db

}
