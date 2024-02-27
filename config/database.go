package config

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type DatabaseCredentials struct {
	Server   string
	User     string
	Password string
	Port     string
	Database string
}

func ConnectDatabase(e DatabaseCredentials) *sqlx.DB {
	logrus.Info("Connecting to database... : ", e.ConnectionString())
	conn, err := sqlx.Open("postgres", e.ConnectionString())

	if err != nil {
		panic(errors.New("Error connecting to database: " + err.Error()))
	}

	err = conn.Ping()
	if err != nil {
		panic(errors.New("Error pinging database: " + err.Error()))
	}
	return conn
}

func (db DatabaseCredentials) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.Server, db.Port, db.User, db.Password, db.Database)
}
