module github.com/workfoxes/kayo

go 1.16

require (
	github.com/adshao/go-binance/v2 v2.3.3 // indirect
	github.com/alpacahq/alpaca-trade-api-go v1.9.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/gofiber/fiber/v2 v2.21.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/markcheno/go-quote v0.0.0-20210728005305-f6d452d96e34
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/workfoxes/calypso v1.1.1
	github.com/workfoxes/tripwire v0.0.1
	go.oneofone.dev/ta v0.0.7 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2
	golang.org/x/sys v0.0.0-20211113001501-0c823b97ae02 // indirect
	golang.org/x/tools v0.1.8-0.20211029000441-d6a9af8af023 // indirect
	gonum.org/v1/gonum v0.9.3 // indirect
	gorm.io/gorm v1.22.2 // indirect
)

replace github.com/workfoxes/calypso => ../calypso

replace github.com/workfoxes/tripwire => ../tripwire
