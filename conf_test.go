package main

import "testing"

func TestInitializeMongoHost(t *testing.T) {
	a := AppConfig{}
	a.Initialize()
	if a.MongoHost != "localhost" {
		t.Error("Initialize should have set MongoHost to localhost")
	}
}

func TestInitializeMongoPort(t *testing.T) {
	a := AppConfig{}
	a.Initialize()
	if a.MongoPort != 27017 {
		t.Error("Initialize should have set MongoPort to 27017")
	}
}

func TestInitializeMongoDatabase(t *testing.T) {
	a := AppConfig{}
	a.Initialize()
	if a.MongoDatabase != "surplus" {
		t.Error("Initialize should have set MongoDatabase to surplus")
	}
}

func TestInitializeServerAddr(t *testing.T) {
	a := AppConfig{}
	a.Initialize()
	if a.ServerAddr != "localhost" {
		t.Error("Initialize should have set ServerAddr to localhost")
	}
}

func TestInitializeServerPort(t *testing.T) {
	a := AppConfig{}
	a.Initialize()
	if a.ServerPort != 8080 {
		t.Error("Initialize should have set ServerPort to 8080")
	}
}
