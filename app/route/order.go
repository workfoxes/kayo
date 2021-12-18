package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workfoxes/calypso/utils"
	"github.com/workfoxes/kayo/app/service"
	"github.com/workfoxes/kayo/app/session"
)

func RegisterOrder(app fiber.Router) {
	route := app.Group("/order")

	route.Use(func(c *fiber.Ctx) error {
		s := c.Locals("session").(*session.BaseSession)
		c.Locals("service", &service.Order{Session: s})
		return c.Next()
	})

	route.Get("/", getOrders)
	route.Get("/:orderId", getUser)
	route.Get("/history", getUser)
}

func getOrders(c *fiber.Ctx) error {
	page, limit := utils.GetPageAndLimit(c)
	_service := c.Locals("service").(*service.Order)
	_orders := _service.GetOrders(page, limit)
	return c.JSON(_orders)
}
