package application

import (
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/entity"
)

func GetStatement(customerID int) (*entity.StatementResponse, error) {
	customerBalance, err := GetCustomerBalance(customerID)
	if err != nil {
		if customerBalance == nil {
			return nil, entity.ErrCustomerNotFound
		}
		return nil, err
	}

	transactions, err := GetTransactions(customerID)
	if err != nil {
		return nil, err
	}

	return &entity.StatementResponse{
		Balance: entity.StatementBalance{
			Total:         customerBalance.Balance,
			Limit:         customerBalance.Limit,
			StatementDate: time.Now(),
		},
		LastTransactions: transactions,
	}, nil
}
