package entities

import (
	"time"

	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/response"
)

type Transactions struct {
	Limit       int       `db:"limite"`
	Value       int       `db:"valor"`
	Type        string    `db:"tipo"`
	Description string    `db:"descricao"`
	CreatedAt   time.Time `db:"realizada_em"`
}

func ToExtractResponse(t []Transactions) response.ExtractResponse {
	lastTransactions := []response.Transaction{}
	if len(t) == 0 {
		return response.ExtractResponse{}
	}
	for i, v := range t {
		if i == 10 {
			break
		}
		lastTransactions = append(lastTransactions, response.Transaction{
			Value:       v.Value,
			Type:        v.Type,
			Description: v.Description,
			RealizedIn:  v.CreatedAt,
		})
	}
	return response.ExtractResponse{
		Amount: response.Amount{
			Total: GetBalance(t),
			Date:  time.Now(),
			Limit: t[0].Limit,
		},
		LastTransactions: lastTransactions,
	}
}

func GetBalance(t []Transactions) int {
	total := 0
	for _, v := range t {
		if v.Type == "d" {
			total -= v.Value
		} else {
			total += v.Value
		}

	}
	return total
}
