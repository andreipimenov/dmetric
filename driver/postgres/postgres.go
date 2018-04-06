package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(url string) (*Postgres, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &Postgres{
		DB: db,
	}, nil
}

func (p *Postgres) Close() error {
	return p.DB.Close()
}

func (p *Postgres) Ping() error {
	return p.DB.Ping()
}

func (p *Postgres) Exec(q string, args ...interface{}) (sql.Result, error) {
	return p.DB.Exec(q, args...)
}

func (p *Postgres) Query(q string, args ...interface{}) (*sql.Rows, error) {
	return p.DB.Query(q, args...)
}

func (p *Postgres) QueryRow(q string, args ...interface{}) *sql.Row {
	return p.DB.QueryRow(q, args...)
}
