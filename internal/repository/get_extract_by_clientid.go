package repository

import (
	"context"
	"fmt"

	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository/entities"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/response"
)

func (c ClientRepository) GetExtractByClientId(id int) response.ExtractResponse {
	transactions := []entities.Transactions{}

	rows, err := c.Db.Query(context.Background(),
		`SELECT
    t.valor,
    t.tipo,
    t.descricao,
    t.realizada_em,
    c.limite
	FROM transacoes t
	JOIN clientes c ON t.cliente_id = c.id
	WHERE t.cliente_id = $1
	ORDER BY t.realizada_em DESC;`, fmt.Sprint(id))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		transaction := entities.Transactions{}
		err = rows.Scan(&transaction.Value, &transaction.Type, &transaction.Description, &transaction.CreatedAt, &transaction.Limit)
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, transaction)
	}

	return entities.ToExtractResponse(transactions)
}
