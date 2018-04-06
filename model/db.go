package model

import (
	"database/sql"
)

type DB interface {
	Ping() error
	Exec(q string, args ...interface{}) (sql.Result, error)
	Query(q string, args ...interface{}) (*sql.Rows, error)
	QueryRow(q string, args ...interface{}) *sql.Row
}
