package repository

import (
	"fmt"
	"time"

	errorapp "github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/error"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository/entities"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/request"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/response"
	"github.com/jmoiron/sqlx"
)

type IClientRepository interface {
	ClientExists(id int) bool
	GetExtractByClientId(id int) response.ExtractResponse
	createCredit(transaction request.TransactionRequest)
	CreateTransaction(request request.TransactionRequest) (response.TransactionResponse, *errorapp.Error)
}

type clientRepository struct {
	db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) IClientRepository {
	return clientRepository{db}
}

func (c clientRepository) CreateTransaction(request request.TransactionRequest) (response.TransactionResponse, *errorapp.Error) {
	var limit int
	err := c.db.Get(&limit, `SELECT limite FROM clientes WHERE id = $1`, request.ClientId)
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
	WHERE t.cliente_id = $1`, request.ClientId)

	if err != nil {
		panic(err)
	}
	balance := entities.GetBalance(t)

	if request.Type == "c" {
		c.createCredit(request)
		return response.TransactionResponse{
			Limit:   limit,
			Balance: balance + request.Value,
		}, nil
	}

	newBalance := balance - request.Value

	if newBalance < -limit {
		return response.TransactionResponse{}, &errorapp.Error{Code: 422, Message: fmt.Sprintf("insufficient funds, limit: %d, balance: %d", limit, balance)}
	}

	_, err = c.db.MustExec(`
	INSERT INTO transacoes 
	(cliente_id, valor, tipo, descricao, realizada_em) VALUES
	($1, $2, $3, $4, $5)`, request.ClientId, request.Value, request.Type, request.Description, time.Now()).RowsAffected()

	if err != nil {
		panic(err)
	}

	return response.TransactionResponse{
		Limit:   limit,
		Balance: newBalance,
	}, nil
}
