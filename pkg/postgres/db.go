package postgres

import (
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

type Database struct {
	Db *sql.DB
}

func NewDb(dsn string) (*Database, error) {
	conn, err := newDbConn(dsn)
	if err != nil {
		return nil, err
	}
	return &Database{conn}, nil
}

func newDbConn(dsn string) (*sql.DB, error) {
	//url := postgres://username:password@localhost:5432/database
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err == nil {
		return conn, err
	}
	for i := 1; i < 5; i++ {
		time.Sleep(2 * time.Second)
		err = conn.Ping()
		if err == nil {
			return conn, err
		}
	}
	return conn, errors.New("couldn't ping database even after multiple tries")
}
