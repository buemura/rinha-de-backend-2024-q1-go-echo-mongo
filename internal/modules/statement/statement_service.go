package statement

import (
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/modules/customer"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/modules/transaction"
)

func GetStatement(customerID int) (*StatementResponse, error) {
	customerBalance, err := customer.GetCustomerBalance(customerID)
	if err != nil {
		if customerBalance == nil {
			return nil, customer.ErrCustomerNotFound
		}
		return nil, err
	}

	transactions, err := transaction.GetTransactions(customerID)
	if err != nil {
		return nil, err
	}

	return &StatementResponse{
		Balance: StatementBalance{
			Total:         customerBalance.Balance,
			Limit:         customerBalance.Limit,
			StatementDate: time.Now(),
		},
		LastTransactions: transactions,
	}, nil
}
