package repository

import (
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository/entities"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/request"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/response"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type IClientRepository interface {
	ClientExists(id int) bool
	GetExtractByClientId(id int) response.ExtractResponse
	CreateCredit(transaction request.TransactionRequest)
	GetBalance(id int) int
}

type clientRepository struct {
	db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) IClientRepository {
	return clientRepository{db}
}

func (c clientRepository) GetBalance(id int) int {
	var limit int
	err := c.db.Get(&limit, `SELECT limite FROM clientes WHERE id = $1`, id)
	if err != nil {
		panic(err)
	}

	var t []entities.Transactions

	err = c.db.Select(&t,
	`SELECT
    t.valor,
    t.tipo,
    t.descricao,
    t.realizada_em,
    c.limite
	FROM transacoes t
	JOIN clientes c ON t.cliente_id = c.id
	WHERE t.cliente_id = $1`, id)

	if err != nil {
		panic(err)
	}
	logrus.Info("balance: ", entities.GetBalance(t))
	return entities.GetBalance(t)
}
