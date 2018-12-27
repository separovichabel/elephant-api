package dao

import (
	"log"

	. "github.com/separovichabel/monkey-api/models/user"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserDAO is a DAO DB
type UserDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION users
	COLLECTION = "users"
)

// Connect to Connect
func (m *UserDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// GetAll to GetAll
func (m *UserDAO) GetAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// GetByID to GetByID
func (m *UserDAO) GetByID(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Create to Create
func (m *UserDAO) Create(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// Delete to Delete
func (m *UserDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update to Update
func (m *UserDAO) Update(id string, user User) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &user)
	return err
}
