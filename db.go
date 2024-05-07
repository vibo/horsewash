package main

import (
	"database/sql"
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

	return &DB{DB: db}, nil
}

func (db *DB) GetUser(id string) (*User, error ) {
	user := &User{}

	err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}