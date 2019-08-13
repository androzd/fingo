package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Password string   `json:"password,omitempty" bson:"password,omitempty"`
	Roles    []string `json:"roles" bson:"roles"`
}

type UserInterface interface {
	ModelInterface
	Find(login string) error
	GetAccounts() ([]Account, error)
}

func (u *User) table() string {
	return "users"
}
func (u *User) Create() error {
	_, err := mongoDatabase.Collection(u.table()).InsertOne(modelContext(), u)
	return err
}

func (u *User) Update() error {
	filter := bson.D{{"_id", u.Username}}
	_, err := mongoDatabase.Collection(u.table()).UpdateOne(modelContext(), filter, u)
	return err
}

func (u *User) Delete() error {
	filter := bson.D{{"_id", u.Username}}
	_, err := mongoDatabase.Collection(u.table()).DeleteOne(modelContext(), filter)
	return err
}

func (u *User) Find(login string) error {
	filter := bson.D{{"_id", login}}
	return mongoDatabase.Collection(u.table()).FindOne(modelContext(), filter).Decode(&u)
}

func (u *User) GetAccounts() (accounts []Account, error error) {
	filter := bson.D{{"user_id", u.Username}}
	cursor, err := mongoDatabase.Collection("accounts").Find(modelContext(), filter)

	if err != nil {
		error = err
		panic(err)
		return
	}

	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
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