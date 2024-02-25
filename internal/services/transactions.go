package services

import (
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/error"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/request"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/response"
)

type ITransactionsService interface {
	CreateTransaction(transaction request.TransactionRequest) (response.TransactionResponse, *errorapp.Error)
}

type TransactionsService struct {
	repository repository.IClientRepository
}

func NewTransactionsService(r repository.IClientRepository) ITransactionsService {
	return TransactionsService{r}
}

func (ts TransactionsService) CreateTransaction(transaction request.TransactionRequest) (response.TransactionResponse, *errorapp.Error) {
	error := transaction.ValidateBody()
	if error != nil {
		return response.TransactionResponse{}, error
	}

	exists := ts.repository.ClientExists(transaction.ClientId)
	if !exists {
		return response.TransactionResponse{}, &errorapp.Error{Code: 404, Message: "client not found"}
	}

	return response.TransactionResponse{}, nil
}
