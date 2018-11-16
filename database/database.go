package database

import (
	"github.com/globalsign/mgo"
	"os"
)

var dbSession *mgo.Session

type Database struct {
	dbURL        string
	dbName       string
	dbCollection string
}

func Init() {
	var err error
	dbSession, err = mgo.Dial(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
}
