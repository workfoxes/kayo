package app

import (
	"github.com/workfoxes/calypso"
	"github.com/workfoxes/kayo/app/route"
)

func RigisterBlueprints(app *calypso.ApplicationServer) {
	blueprint := app.Server.Group("/api/v1")

	route.RegisterUser(blueprint)
}

func StartKayoServer() {
	kayoServer := calypso.New(&calypso.ApplicationConfig{
		Name: "kayo",
		Port: 8000,
	})
	RigisterBlueprints(kayoServer)
	kayoServer.Start()
}
