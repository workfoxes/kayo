package app

import (
	"github.com/workfoxes/calypso"
	"github.com/workfoxes/calypso/pkg/middleware/request"
	"github.com/workfoxes/kayo/app/route"
	"github.com/workfoxes/kayo/app/server"
	"github.com/workfoxes/kayo/internal/model"
	"gorm.io/gorm"
)

func RegisterBlueprints(app *calypso.ApplicationServer) {
	blueprint := app.Server.Group("/api/v1")

	route.RegisterUser(blueprint)
}

func StartKayoServer() {
	kayoServer := calypso.New(&calypso.ApplicationConfig{
		Name: "kayo",
		Port: 8000,
	})
	kayoServer.Use(request.New(request.Config{
		PreRequest:  server.PreRequest,
		PostRequest: server.PostRequest,
	}))
	kayoServer.Invoker(func(DB *gorm.DB) {
		model.AutoMigrateModel(DB)
	})
	RegisterBlueprints(kayoServer)
	kayoServer.Start()
}
