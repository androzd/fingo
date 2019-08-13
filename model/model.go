package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mongoDatabase *mongo.Database

func Initialize() {
	fmt.Println("Starting connection to mongo")
	c, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	// Create connect
	err = c.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	mongoDatabase = c.Database("finance")
}

func modelContext() context.Context {
	return context.TODO()
}

type ModelInterface interface {
	Create() error
	Update() error
	Delete() error
	table() string
}