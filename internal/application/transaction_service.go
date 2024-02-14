package application

import (
	"context"
	"log"
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/entity"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/infra/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTransactions(customerID int) ([]entity.Transaction, error) {
	filter := bson.M{"customer_id": customerID}
	options := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(10)

	cursor, err := database.TrxColl.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var transactions []entity.Transaction
	for cursor.Next(context.Background()) {
		var transaction entity.Transaction
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
		transactions = []entity.Transaction{}
	}

	return transactions, nil
}

func CreateTransaction(customerID int, trx *entity.CreateTransactionRequest) (*entity.CreateTransactionResponse, error) {
	customer, err := GetCustomerBalance(customerID)
	if err != nil {
		if customer == nil {
			return nil, entity.ErrCustomerNotFound
		}
		return nil, err
	}

	newBalance := customer.Balance
	if trx.Type == "d" {
		newBalance -= trx.Amount
	} else {
		newBalance += trx.Amount
	}
	if customer.Limit+newBalance < 0 {
		return nil, entity.ErrCustomerNoLimit
	}

	updateCustomerFilter := bson.M{"customer_id": customerID}
	updateCustomer := bson.M{"$set": bson.M{"balance": newBalance}}
	err = database.CustColl.FindOneAndUpdate(database.Ctx, updateCustomerFilter, updateCustomer).Err()
	if err != nil {
		return nil, err
	}

	transaction := bson.M{
		"customer_id": customerID,
		"amount":      trx.Amount,
		"type":        trx.Type,
		"description": trx.Description,
		"created_at":  time.Now(),
	}
	_, err = database.TrxColl.InsertOne(database.Ctx, transaction)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &entity.CreateTransactionResponse{
		Balance: newBalance,
		Limit:   customer.Limit,
	}, nil
}
