package model

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

// var dbSession *mgo.Session

// Database obj containing db credentials
type Database struct {
	DBURL        string
	DBName       string
	DBCollection string
}

// Init initiates the db
func (db *Database) Init() {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DBName).C(db.DBCollection).EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}

}

// IDE: ha egne admin enpoints som bruker alle db funksjonaliteten, som for eksempel
// Add puts a user in the db
func (db *Database) Add(t User) error {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	err = session.DB(db.DBName).C(db.DBCollection).Insert(t)
	if err != nil {
		fmt.Printf("Error in insert(): %v", err.Error())
		return err
	}
	return nil
}

// Get retrieves a user from the db by the id
func (db *Database) Get(keyID string) (User, error) {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	user := User{}

	err = session.DB(db.DBName).C(db.DBCollection).Find(bson.M{"id": keyID}).One(&user)
	if err != nil {
		fmt.Printf("Error %v", err.Error())
		return User{}, err
	}

	return user, nil
}

// GetAll gets all the user from db
func (db *Database) GetAll() ([]User, error) {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	var users []User

	err = session.DB(db.DBName).C(db.DBCollection).Find(bson.M{}).All(&users)
	if err != nil {
		fmt.Printf("Error %v", err.Error())
		return []User{}, err
	}

	return users, nil
}
