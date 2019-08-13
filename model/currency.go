package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Currency struct {
	Iso string `json:"_id,omitempty" bson:"_id,omitempty"`
}

type CurrencyInterface interface {
	ModelInterface
	All() ([]Currency, error)
	Find(iso string) (Currency, error)
}

func (c Currency) table() string {
	return "currencies"
}

func (c Currency) Create() error {
	_, err := mongoDatabase.Collection(c.table()).InsertOne(modelContext(), c)
	return err
}

func (c Currency) Update() error {
	filter := bson.D{{"_id", c.Iso}}
	_, err := mongoDatabase.Collection(c.table()).UpdateOne(modelContext(), filter, c)
	return err
}

func (c Currency) Delete() error {
	filter := bson.D{{"_id", c.Iso}}
	_, err := mongoDatabase.Collection(c.table()).DeleteOne(modelContext(), filter)
	return err
}

func (c Currency) Find(iso string) (currency Currency, error error) {
	filter := bson.D{{"_id", iso}}
	error = mongoDatabase.Collection(c.table()).FindOne(modelContext(), filter).Decode(&currency)

	return
}

func (c Currency) All() (currencies []Currency, error error) {
	filter := bson.D{}
	cursor, err := mongoDatabase.Collection(c.table()).Find(modelContext(), filter)

	if err != nil {
		error = err
		panic(err)
		return
	}

	for cursor.Next(context.TODO()) {
		var elem Currency
		err := cursor.Decode(&elem)
		if err != nil {
			error = err
			panic(err)
			return
		}
		currencies = append(currencies, elem)
	}

	return
}
