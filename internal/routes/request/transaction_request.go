package request

import errorapp "github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/error"

type TransactionRequest struct {
	ClientId    int
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func (tr TransactionRequest) ValidateBody() *errorapp.Error {
	if tr.Value < 0 || tr.ClientId < 0 {
		return &errorapp.Error{Code: 400, Message: "invalid value"}
	}

	if tr.Type != "c" && tr.Type != "d" {
		return &errorapp.Error{Code: 400, Message: "invalid type"}
	}

	if len(tr.Description) > 10 {
		return &errorapp.Error{Code: 400, Message: "invalid description"}
	}
	return nil
}
