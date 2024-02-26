package repository

import (
	"time"

	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/request"
)

func (c clientRepository) CreateCredit(transaction request.TransactionRequest) {

	_, err := c.db.MustExec(`
		INSERT INTO
		transacoes (valor, tipo, descricao, realizada_em, cliente_id)
		VALUES ($1, $2, $3, $4, $5)`, transaction.Value, transaction.Type, transaction.Description, time.Now(), transaction.ClientId).RowsAffected()

	if err != nil {
		panic(err)
	}
}
