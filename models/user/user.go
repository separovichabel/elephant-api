package user

import (
	"gopkg.in/mgo.v2/bson"
)

// User is a model of user
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Email    string        `bson:"email" json:"email"`
	Phone    string        `bson:"phone,omitempty" json:"phone,omitempty"`
	Password string        `bson:"password" json:"password"`
	Address  Address       `bson:"address,omitempty" json:"address,omitempty"`
}

// Address is a model of users Address
type Address struct {
	Street       string
	Number       int
	Complement   string
	Neighborhood string
	City         string
	State        string
	Zipcode      string
}
