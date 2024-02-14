package database

import (
	"context"
	"log"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Ctx      = context.TODO()
	CustColl *mongo.Collection
	TrxColl  *mongo.Collection
)

func Connect() {
	clientOptions := options.Client().ApplyURI(config.DATABASE_URL)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	CustColl = client.Database("rinha").Collection("customers")
	TrxColl = client.Database("rinha").Collection("transactions")
}
