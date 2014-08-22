package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danrjohnson/surplus_api/surplus"
	"github.com/go-martini/martini"
	"github.com/kelseyhightower/envconfig"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Config AppConfig

func main() {
	Config.Initialize()
	err := envconfig.Process("surplus", &Config)
	if err != nil {
		log.Fatal(err)
	}
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(DB())
	m.Get("/", func(r render.Render, db *mgo.Database) {
		var items []surplus.SurplusItem
		db.C("surplusitems").Find(nil).All(&items)
		r.JSON(200, items)
	})
	m.Get("/state/(?P<state>[A-Z][A-Z])", func(r render.Render, db *mgo.Database, params martini.Params) {
		var items []surplus.SurplusItem
		db.C("surplusitems").Find(bson.M{"state": params["state"]}).All(&items)
		r.JSON(200, items)
	})
	m.Get("/state/(?P<state>[A-Z][A-Z])/(?P<county>[A-Z]+)",
		func(r render.Render, db *mgo.Database, params martini.Params) {
			var items []surplus.SurplusItem
			db.C("surplusitems").Find(
				bson.M{"state": params["state"], "county": params["county"]},
			).All(&items)
			r.JSON(200, items)
		})
	listen_addr := fmt.Sprintf("%s:%d", Config.ServerAddr, Config.ServerPort)
	log.Fatal(http.ListenAndServe(listen_addr, m))
}
