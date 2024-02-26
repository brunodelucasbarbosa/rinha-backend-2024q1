package repository

import (
	"fmt"
	"strings"

	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository/entities"
)

func (c clientRepository) ClientExists(id int) bool {
	client := entities.Client{}

	err := c.db.Get(&client, "SELECT * FROM clientes WHERE id = $1", fmt.Sprint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return false
		}
		panic(err)
	}

	return true
}
