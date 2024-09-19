package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

func Connect() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	return session
}
