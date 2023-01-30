package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	username = "root"
	password = "nihan"
	hostname = "127.0.0.1:3306"
	dbName   = "gotodo"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

var (
	db  *sql.DB
	err error
)

func Connect() (db *sql.DB) {
	db, err = sql.Open(dbDriver, dsn(""))

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(`USE gotodo`)

	if err != nil {
		fmt.Println(err)
	}
	return db
}

func CreateDB() (db *sql.DB) {
	db, err = sql.Open(dbDriver, dsn(""))

	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS gotodo.todos(
			ID INT AUTO_INCREMENT,
			Item TEXT NOT NULL,
			Completed BOOLEAN DEFAULT FALSE,

			PRIMARY KEY(ID)
		);
	`)

	if err != nil {
		fmt.Println(err, "Can't create table!")
	}

	return nil
}
