package main

type AppConfig struct {
	MongoHost     string
	MongoPort     int
	MongoDatabase string

	ServerAddr string
	ServerPort int
}

func (a *AppConfig) Initialize() {
	a.MongoHost = "localhost"
	a.MongoPort = 27017
	a.MongoDatabase = "surplus"

	a.ServerAddr = "localhost"
	a.ServerPort = 8080
}
