package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseCredentials struct {
	Server   string
	User     string
	Password string
	Port     string
	Database string
}

func ConnectDatabase(e DatabaseCredentials) *sqlx.DB {
	conn, err := sqlx.Open("postgres", e.ConnectionString())
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	return conn
}

func (db DatabaseCredentials) ConnectionString() string {
	return fmt.Sprintf("user=%s password=%s port=%s database=%s sslmode=disable", db.User, db.Password, db.Port, db.Database)
}
