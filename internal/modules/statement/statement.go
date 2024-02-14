package statement

import (
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/modules/transaction"
)

type StatementBalance struct {
	Total         int       `json:"total"`
	StatementDate time.Time `json:"data_extrato"`
	Limit         int       `json:"limite"`
}

type StatementResponse struct {
	Balance          StatementBalance          `json:"saldo"`
	LastTransactions []transaction.Transaction `json:"ultimas_transacoes"`
}
