package entity

import (
	"time"
)

type StatementBalance struct {
	Total         int       `json:"total"`
	StatementDate time.Time `json:"data_extrato"`
	Limit         int       `json:"limite"`
}

type StatementResponse struct {
	Balance          StatementBalance `json:"saldo"`
	LastTransactions []Transaction    `json:"ultimas_transacoes"`
}
