package main

import (
	"fmt"

	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
)

func DB() martini.Handler {
	conn_string := fmt.Sprintf(
		"mongodb://%s:%d", Config.MongoHost, Config.MongoPort)
	session, err := mgo.Dial(conn_string)
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(Config.MongoDatabase))
		defer s.Close()
		c.Next()
	}
}
