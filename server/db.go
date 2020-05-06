package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Store interface {
	Insert([]string) error
}

type SQLLiteStore struct {
	db *sql.DB
}

//Connect init db connection and config table
func (s *SQLLiteStore) Connect(dbName string) error {
	var err error

	s.db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return fmt.Errorf("Open db error: %v\n", err)
	}

	if err = s.db.Ping(); err != nil {
		return fmt.Errorf("Connection db error: %v\n", err)
	}

	_, err = s.db.Exec(`create table IF NOT EXISTS schedule (event TEXT, executor TEXT, count TEXT, 
created TIMESTAMP  DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		return fmt.Errorf("Init table error: %v\n", err)
	}
	return err
}

//Insert insert record to db
func (s *SQLLiteStore) Insert(rec []string) error {
	var err error
	if len(rec) <= 3 {
		_, err = s.db.Exec("insert into schedule (event, executor, count) values ($1, $2, $3)",
			rec[0], rec[1], rec[2])
	} else {
		err = fmt.Errorf("Incorrect record len\n")
	}

	return err
}
