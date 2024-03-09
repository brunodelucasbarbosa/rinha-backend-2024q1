package repository

import (
	"context"
	"fmt"
	"time"

	errorapp "github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/error"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository/entities"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/request"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/response"
	"github.com/sirupsen/logrus"
)

func (c ClientRepository) CreateTransaction(request request.TransactionRequest) (response.TransactionResponse, *errorapp.Error) {
	var limit int
	tx, err := c.Db.Begin(context.Background())
	defer tx.Rollback(context.Background())

	if err != nil {
		panic(err)
	}

	err = tx.QueryRow(context.Background(), `SELECT limite FROM clientes WHERE id = $1`, request.ClientId).Scan(&limit)

	if err != nil {
		logrus.Error("error on CreateTransaction1: ", err)
		panic(err)
	}

	var t []entities.Transactions

	rows, err := tx.Query(context.Background(),
		`SELECT
    t.valor,
    t.tipo,
    t.descricao,
    t.realizada_em,
    c.limite
	FROM transacoes t
	JOIN clientes c ON t.cliente_id = c.id
	WHERE t.cliente_id = $1`, request.ClientId)

	if err != nil && err.Error() != "no rows in result set" {
		logrus.Error("error on CreateTransaction2: ", err)
		panic(err)
	}

	for rows.Next() {
		var transaction entities.Transactions
		err = rows.Scan(&transaction.Value, &transaction.Type, &transaction.Description, &transaction.CreatedAt, &limit)
		if err != nil {
			panic(err)
		}
		t = append(t, transaction)
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

	_, err = tx.Exec(context.Background(),
		`INSERT INTO transacoes 
	(cliente_id, valor, tipo, descricao, realizada_em) VALUES
	($1, $2, $3, $4, $5)`, request.ClientId, request.Value, request.Type, request.Description, time.Now())

	if err != nil {
		logrus.Error("error on CreateTransaction3: ", err)
		panic(err)
	}
	tx.Commit(context.Background())
	return response.TransactionResponse{
		Limit:   limit,
		Balance: newBalance,
	}, nil
}
