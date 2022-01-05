module github.com/workfoxes/kayo

go 1.16

require (
	github.com/gofiber/fiber/v2 v2.21.0
	github.com/influxdata/influxdb-client-go/v2 v2.6.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/workfoxes/calypso v1.1.1
	github.com/workfoxes/tripwire v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f
	golang.org/x/sys v0.0.0-20211113001501-0c823b97ae02 // indirect
	golang.org/x/tools v0.1.8-0.20211029000441-d6a9af8af023 // indirect
	gorm.io/gorm v1.22.2
)

replace (
	github.com/workfoxes/calypso => ../calypso
	github.com/workfoxes/tripwire => ../tripwire
)
