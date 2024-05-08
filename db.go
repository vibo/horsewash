package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "modernc.org/sqlite"
)


type DB struct {
	*sql.DB
}


func openDatabase(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil { 
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	initDatabase(db)

	return &DB{DB: db}, nil
}

func initDatabase(db *sql.DB) {
	initScript, err := os.ReadFile("./init.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(initScript))
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) GetUser(id string) (*User, error ) {
	user := &User{}

	err := db.QueryRow("SELECT id, username FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}