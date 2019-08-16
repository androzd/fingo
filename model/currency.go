package model

import (
	"upper.io/db.v3"
)

type Currency struct {
	Iso string `json:"_id,omitempty" bson:"_id,omitempty"`
}

type CurrencyInterface interface {
	ModelInterface
	All() ([]Currency, error)
	Find(iso string) (Currency, error)
}

func (m Currency) table() db.Collection {
	return mongoDatabase.Collection("currencies")
}

func (m Currency) Create() error {
	_, err := m.table().Insert(m)
	return err
}

func (m Currency) Update() error {
	//filter := bson.D{{"_id", c.Iso}}
	err := m.table().Find(m).Update(m)
	return err
}

func (m Currency) Delete() error {
	//filter := bson.D{{"_id", c.Iso}}
	err := m.table().Find(m).Delete()
	return err
}

func (m Currency) Find(iso string) (currency Currency, error error) {
	//filter := bson.D{{"_id", iso}}
	error = m.table().Find(m).One(&currency)

	return
}

func (m Currency) All() (currencies []Currency, error error) {
	//filter := bson.D{}
	//cursor, err := mongoDatabase.Collection(c.table()).Find(modelContext(), filter)

	err := m.table().Find().All(&currencies)
	if err != nil {
		error = err
		panic(err)
		return
	}

	return
}
