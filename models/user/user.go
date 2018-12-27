package user

import (
	"gopkg.in/mgo.v2/bson"
)

// User is a model of user
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string        `bson:"name" json:"name"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Phone    string        `bson:"phone,omitempty" json:"phone,omitempty"`
	Address  *Address      `bson:"address,omitempty" json:"address,omitempty"`
}

// Address is a model of users Address
type Address struct {
	Street       string `bson:"street,omitempty" json:"street,omitempty"`
	Number       *int   `bson:"number,omitempty" json:"number,omitempty"`
	Complement   string `bson:"complement,omitempty" json:"complement,omitempty"`
	Neighborhood string `bson:"neighborhood,omitempty" json:"neighborhood,omitempty"`
	City         string `bson:"city,omitempty" json:"city,omitempty"`
	State        string `bson:"state,omitempty" json:"state,omitempty"`
	Zipcode      string `bson:"zipcode,omitempty" json:"zipcode,omitempty"`
}
