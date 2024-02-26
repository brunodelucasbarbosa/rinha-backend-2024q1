package repository

import (
	"fmt"

	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository/entities"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/response"
)

func (c clientRepository) GetExtractByClientId(id int) response.ExtractResponse {
	transactions := []entities.Transactions{}

	err := c.db.Select(&transactions,
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

	return entities.ToExtractResponse(transactions)
}
