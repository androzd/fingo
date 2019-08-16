package model

import (
	"upper.io/db.v3"
)

type User struct {
	Username string   `json:"username" bson:"_id,omitempty"`
	Password string   `json:"-" bson:"password,omitempty"`
	Roles    []string `json:"roles" bson:"roles"`
}

type UserInterface interface {
	ModelInterface
	Find(login string) error
	GetAccounts() ([]Account, error)
}

func (m User) table() db.Collection {
	return mongoDatabase.Collection("users")
}
func (m User) Create() error {
	_, err := m.table().Insert(m)
	return err
}

func (m User) Update() error {
	//filter := bson.D{{"_id", u.Username}}
	err := m.table().Find(m).Update(m)
	return err
}

func (m User) Delete() error {
	//filter := bson.D{{"_id", u.Username}}
	err := m.table().Find(m).Delete()
	return err
}

func (m *User) Find(login string) error {
	//filter := bson.D{{"_id", login}}
	return m.table().Find(m).One(&m)
}

func (m User) GetAccounts() (accounts []Account, error error) {
	//filter := bson.D{{"user_id", u.Username}}
	err := Account{}.table().Find(m).All(&accounts)

	if err != nil {
		error = err
		panic(err)
		return
	}

	return
}