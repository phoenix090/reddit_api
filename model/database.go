package model

import (
	"log"

	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

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

// Add puts a user in the db
func (db *Database) Add(t User) error {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	err = session.DB(db.DBName).C(db.DBCollection).Insert(t)
	if err != nil {
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
		return []User{}, err
	}

	return users, nil
}

// Count returns the nr of users in the collection
func (db *Database) Count() int {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	count, err := session.DB(db.DBName).C(db.DBCollection).Count()
	if err != nil {
		return -1
	}

	return count
}

// DeleteUser removes one user from db
func (db *Database) DeleteUser(keyID string) error {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DBName).C(db.DBCollection).Remove(bson.M{"id": keyID})
	if err != nil {
		return err
	}
	return nil
}

// DeleteAll deletes all the users from db
func (db *Database) DeleteAll() error {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	_, err = session.DB(db.DBName).C(db.DBCollection).RemoveAll(bson.M{})
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Upsert(t User) {
	session, err := mgo.Dial(db.DBURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.DB(db.DBName).C(db.DBCollection).Upsert(bson.M{"id": t.ID}, bson.M{"url": t.URL})
}
