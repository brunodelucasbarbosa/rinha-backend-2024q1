package main

import (
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/config"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/services"
)

func main() {
	envs := config.LoadEnvsConfigs()

	db := config.ConnectDatabase(envs.DatabaseCredentials)
	repository := repository.NewClientRepository(db)

	service := services.NewTransactionsService(repository)
	routes.StartRoutes(service)
}
