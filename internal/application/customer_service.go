package application

import (
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/entity"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/infra/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCustomerBalance(customerID int) (*entity.CustomerBalance, error) {
	var customerBalance entity.CustomerBalance
	findCustomerFilter := bson.D{{Key: "customer_id", Value: customerID}}
	err := database.CustColl.FindOne(database.Ctx, findCustomerFilter).Decode(&customerBalance)
	if err == mongo.ErrNoDocuments {
		return nil, entity.ErrCustomerNotFound
	}
	if err != nil {
		return nil, err
	}
	return &customerBalance, nil
}
