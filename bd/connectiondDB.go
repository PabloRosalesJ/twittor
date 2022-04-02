package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = connect()
var clientOptions = options.Client().ApplyURI("mongodb://127.0.0.1:27017")

func connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("[DB] - Conected")
	return client
}

func CheckConnetion() bool {
	err := MongoConnection.Ping(context.TODO(), nil)
	return err == nil
}
