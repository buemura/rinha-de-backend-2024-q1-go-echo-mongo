package transaction

import (
	"context"
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/modules/customer"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/shared/database"
	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTransactions(customerID int) ([]Transaction, error) {
	filter := bson.M{"cliente_id": customerID}
	options := options.Find().SetSort(bson.D{{Key: "realizada_em", Value: -1}}).SetLimit(10)

	cursor, err := database.TrxColl.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var transactions []Transaction
	for cursor.Next(context.Background()) {
		var transaction Transaction
		err := cursor.Decode(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	// if transactions == nil {
	// 	transactions = []Transaction{}
	// }

	return transactions, nil
}

func CreateTransaction(customerID int, trx *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	customerBalance, err := customer.GetCustomerBalance(customerID)
	if err != nil {
		if customerBalance == nil {
			return nil, customer.ErrCustomerNotFound
		}
		return nil, err
	}

	trxRes, err := insertTransaction(customerID, customerBalance.Limite, trx.Valor, string(trx.Tipo), trx.Descricao)
	if err != nil {
		return nil, err
	}
	return trxRes, nil
}

func insertTransaction(customerID, limit, trxAmount int, trxType, description string) (*CreateTransactionResponse, error) {
	tx, err := database.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	var balance int
	err = tx.QueryRow(context.Background(), "SELECT valor FROM saldos WHERE cliente_id = $1 FOR UPDATE", customerID).Scan(&balance)
	if err != nil {
		return nil, err
	}

	if trxType == "c" {
		balance += trxAmount
	}
	if trxType == "d" {
		if (balance-trxAmount)*-1 > limit {
			return nil, customer.ErrCustomerNoLimit
		}
		balance -= trxAmount
	}

	batch := &pgx.Batch{}
	batch.Queue(`
        INSERT INTO transacoes (cliente_id, valor, tipo, descricao, realizada_em) 
        VALUES ($1, $2, $3, $4, $5)
    `, customerID, trxAmount, trxType, description, time.Now())
	batch.Queue(`
        UPDATE saldos SET valor = $1 WHERE cliente_id = $2
    `, balance, customerID)

	bRes := tx.SendBatch(context.Background(), batch)
	if _, err := bRes.Exec(); err != nil {
		return nil, err
	}
	if err := bRes.Close(); err != nil {
		return nil, err
	}
	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return &CreateTransactionResponse{
		Saldo:  balance,
		Limite: limit,
	}, nil
}
