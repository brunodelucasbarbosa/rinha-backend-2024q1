package response

import "time"

//ExtractResponse contem as ultimas 10 transacoes
type ExtractResponse struct {
	Amount           Amount        `json:"saldo"`
	LastTransactions []Transaction `json:"ultimas_transacoes"`
}

type Amount struct {
	Total int       `json:"total"`
	Date  time.Time `json:"data_extrato"`
	Limit int       `json:"limite"`
}

type Transaction struct {
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	RealizedIn  time.Time `json:"realizado_em"`
}
