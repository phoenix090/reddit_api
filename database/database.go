package database

import "github.com/globalsign/mgo"

var dbSession *mgo.Session

type Database struct {
	dbURL        string
	dbName       string
	dbCollection string
}

func Init(db *Database) {
	var err error
	dbSession, err = mgo.Dial(db.dbURL)
	if err != nil {
		panic(err)
	}
}
