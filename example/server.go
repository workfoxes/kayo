//+build !test

package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/workfoxes/kayo/pkg/log"
)

func logger() {
	app := fiber.New()
	// spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	// dema, _ := ta.New(spy.Close).Dema(10)
	// fmt.Println(dema)

	app.Get("/", func(c *fiber.Ctx) error {
		log.Print("Default Server")
		log.Trace("Something very low level.")
		log.Debug("Useful debugging information.")
		log.Info("Something noteworthy happened!")
		log.Warn("You should probably take a look at this.")
		log.Error("Something failed but I'm not quitting.")
		// Calls os.Exit(1) after logging
		//log.Fatal("Bye.")
		//// Calls panic() after logging
		//log.Panic("I'm bailing.")
		return c.SendString("Hello, World 👋!")
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
