package services

import (
	errorapp "github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/error"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/repository"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/request"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/response"
)

type TransactionsService struct {
	Repository repository.ClientRepository
}

func (ts TransactionsService) CreateTransaction(transaction request.TransactionRequest) (response.TransactionResponse, *errorapp.Error) {
	error := transaction.ValidateBody()
	if error != nil {
		return response.TransactionResponse{}, error
	}

	exists := ts.Repository.ClientExists(transaction.ClientId)
	if !exists {
		return response.TransactionResponse{}, &errorapp.Error{Code: 404, Message: "client not found"}
	}

	return ts.Repository.CreateTransaction(transaction)
}

func (ts TransactionsService) GetExtract(clientId int) (response.ExtractResponse, *errorapp.Error) {
	exists := ts.Repository.ClientExists(clientId)
	if !exists {
		return response.ExtractResponse{}, &errorapp.Error{Code: 404, Message: "client not found"}
	}

	return ts.Repository.GetExtractByClientId(clientId), nil
}
