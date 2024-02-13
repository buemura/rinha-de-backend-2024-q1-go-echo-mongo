package customer

import (
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/shared/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCustomerBalance(customerID int) (*CustomerBalance, error) {
	var customerBalance CustomerBalance
	err := database.CustColl.FindOne(database.Ctx, bson.D{{Key: "cliente_id", Value: customerID}}).Decode(&customerBalance)
	if err == mongo.ErrNoDocuments {
		return nil, ErrCustomerNotFound
	}
	if err != nil {
		return nil, err
	}
	return &customerBalance, nil
}
