package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository/entities"
	"github.com/sirupsen/logrus"
)

func (c ClientRepository) ClientExists(id int) bool {
	client := entities.Client{}

	err := c.Db.QueryRow(context.Background(), "SELECT * FROM clientes WHERE id = $1", fmt.Sprint(id)).Scan(&client.Id, &client.Name, &client.Limit)
	if err != nil {
		logrus.Error("error on ClientExists: ", err)
		if strings.Contains(err.Error(), "no rows in result set") {
			return false
		}
		panic(err)
	}

	return true
}
