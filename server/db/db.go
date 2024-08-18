package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	// *sql.DB is a connection to the db
	db *sql.DB
}

// making a new db connection
func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/go-chat?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

// closing conn
func (d *Database) Close() {
	d.db.Close()
}

// getter func
func (d *Database) GetDB() *sql.DB {
	// here d is struct and db is propery of 'd' struct
	return d.db
}
