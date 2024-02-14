package transaction

import (
	"context"
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/modules/customer"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/shared/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTransactions(customerID int) ([]Transaction, error) {
	filter := bson.M{"customer_id": customerID}
	options := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(10)

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
	if transactions == nil {
		transactions = []Transaction{}
	}

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

	trxRes, err := insertTransaction(customerID, customerBalance.Limit, customerBalance.Balance, trx.Amount, string(trx.Type), trx.Description)
	if err != nil {
		return nil, err
	}
	return trxRes, nil
}

func insertTransaction(customerID, limit, balance, trxAmount int, trxType, description string) (*CreateTransactionResponse, error) {
	var balanceInc int
	if trxType == "c" {
		balanceInc = trxAmount
	}
	if trxType == "d" {
		if (balance-trxAmount)*-1 > limit {
			return nil, customer.ErrCustomerNoLimit
		}
		balanceInc = -trxAmount
	}

	updateCustomerFilter := bson.M{"customer_id": customerID}
	updateCustomer := bson.M{"$inc": bson.M{"balance": balanceInc}}
	err := database.CustColl.FindOneAndUpdate(database.Ctx, updateCustomerFilter, updateCustomer).Err()
	if err != nil {
		return nil, err
	}

	transaction := bson.M{
		"customer_id": customerID,
		"amount":      trxAmount,
		"type":        trxType,
		"description": description,
		"created_at":  time.Now(),
	}
	_, err = database.TrxColl.InsertOne(database.Ctx, transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionResponse{
		Balance: balance + balanceInc,
		Limit:   limit,
	}, nil
}
