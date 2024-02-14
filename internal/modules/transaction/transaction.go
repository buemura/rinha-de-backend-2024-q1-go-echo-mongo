package transaction

import "time"

type Transaction struct {
	Amount      int       `json:"amount"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTransactionRequest struct {
	Amount      int    `json:"valor" validate:"required,min=1"`
	Type        string `json:"tipo" validate:"required"`
	Description string `json:"descricao" validate:"required,min=1,max=10"`
}

type CreateTransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}
