package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ClientRepository struct {
	Db *pgxpool.Pool
}
