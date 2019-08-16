package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"upper.io/db.v3"
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

func (m Account) table() db.Collection {
	return mongoDatabase.Collection("accounts")
}

func (m Account) Create() error {
	_, err := m.table().Insert(m)
	return err
}

func (m Account) Update() error {
	//filter := bson.D{{"_id", a.Id}}
	//id must
	err := m.table().Find(m).Update(m)
	return err
}

func (m Account) Delete() error {
	err := m.table().Find(m).Delete()
	return err
}

func (m Account) Currency() (currency Currency, error error) {
	error = Currency{}.table().Find(m.CurrencyId).One(currency)
	return
}

func (m Account) All() (accounts []Account, error error) {
	//filter := bson.D{}
	err := m.table().Find(m).All(accounts)

	if err != nil {
		error = err
		panic(err)
		return
	}

	return
}
