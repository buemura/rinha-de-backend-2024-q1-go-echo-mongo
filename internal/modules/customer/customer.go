package customer

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrCustomerNotFound = errors.New("customer not found")
var ErrCustomerNoLimit = errors.New("customer has no limit")

type CustomerBalance struct {
	ID           primitive.ObjectID `bson:"_id"`
	CustomerId   int                `bson:"cliente_id"`
	Limite 		 int                `json:"limite"`
	Saldo  		 int                `json:"saldo"`
}
