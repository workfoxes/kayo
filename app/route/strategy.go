package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workfoxes/calypso/utils"
	"github.com/workfoxes/kayo/app/service"
	"github.com/workfoxes/kayo/app/session"
	"strconv"
)

func RegisterStrategy(app fiber.Router) {
	route := app.Group("/strategy")

	route.Use(func(c *fiber.Ctx) error {
		s := c.Locals("session").(*session.BaseSession)
		c.Locals("service", &service.Strategy{Session: s})
		return c.Next()
	})

	route.Get("/", getOrders)
	route.Get("/:strategyId", getStrategy)
	route.Get("/history", getUser)
}

func getStrategies(c *fiber.Ctx) error {
	page, limit := utils.GetPageAndLimit(c)
	_service := c.Locals("service").(*service.Strategy)
	_orders := _service.GetStrategies(page, limit)
	return c.JSON(_orders)
}

func getStrategy(c *fiber.Ctx) error {
	strategyId := c.Params("strategyId")
	u64, err := strconv.ParseUint(strategyId, 10, 32)
	if err != nil {
		return err
	}
	_service := c.Locals("service").(*service.Strategy)
	_strategy := _service.GetStrategy(uint(u64))
	return c.JSON(_strategy)
}
