package main

import (
	"context"

	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	db, err := pgxpool.New(context.Background(), "postgres://admin:admin@db:5432/rinha")
	if err != nil {
		panic(err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	repository := repository.ClientRepository{Db: db}

	service := services.TransactionsService{Repository: repository}
	routes.StartRoutes(service)
}
