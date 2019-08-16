package model

import (
	"fmt"
	"log"
	"upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

var settings = mongo.ConnectionURL{
	User:     "",
	Password: "",
	Host:     "127.0.0.1:27017",
	Database: "finance",
	Options:  nil,
}

var mongoDatabase db.Database

func Initialize() {
	fmt.Println("Starting connection to mongo")
	//c, err := mongo.ParseURL("mongodb://localhost:27017")
	//fmt.Println(c)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c.Database = "finance"
	var err error
	mongoDatabase, err = mongo.Open(settings)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

//func modelContext() context.Context {
//	return context.TODO()
//}

type ModelInterface interface {
	Create() error
	Update() error
	Delete() error
	table() db.Collection
}