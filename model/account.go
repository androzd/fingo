package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CurrencyId string             `json:"currency_id,omitempty" bson:"currency_id,omitempty"`
	Balance    int64
}

type AccountInterface interface {
	ModelInterface
	All() ([]Account, error)
	Currency() (Currency, error)
}

func (a *Account) table() string {
	return "accounts"
}

func (a *Account) Create() error {
	_, err := mongoDatabase.Collection(a.table()).InsertOne(modelContext(), a)
	return err
}

func (a *Account) Update() error {
	filter := bson.D{{"_id", a.Id}}
	_, err := mongoDatabase.Collection(a.table()).UpdateOne(modelContext(), filter, a)
	return err
}

func (a *Account) Delete() error {
	filter := bson.D{{"_id", a.Id}}
	_, err := mongoDatabase.Collection(a.table()).DeleteOne(modelContext(), filter)
	return err
}

func (a *Account) Currency() (currency Currency, error error) {
	currency, error = Currency{}.Find(a.CurrencyId)
	return
}

func (a Account) All() (accounts []Account, error error) {
	filter := bson.D{}
	cursor, err := mongoDatabase.Collection(a.table()).Find(modelContext(), filter)

	if err != nil {
		error = err
		panic(err)
		return
	}

	for cursor.Next(context.TODO()) {
		var elem Account
		err := cursor.Decode(&elem)
		if err != nil {
			error = err
			panic(err)
			return
		}
		accounts = append(accounts, elem)
	}

	return
}
