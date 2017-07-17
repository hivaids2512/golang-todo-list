package db

import (
	"../config/db"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var mgoSession *mgo.Session

// Creates a new session if mgoSession is nil i.e there is no active mongo session.
//If there is an active mongo session it will return a Clone
func GetMongoSession() *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{db.DB_HOST},
		Timeout:  60 * time.Second,
		Database: db.DB_NAME,
		Username: db.DB_USER,
		Password: db.DB_PASS,
	}
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	return mgoSession.Clone()
}
