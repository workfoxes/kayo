package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workfoxes/calypso/utils"
	"github.com/workfoxes/kayo/app/service"
	"github.com/workfoxes/kayo/app/session"
)

func RegisterUser(app fiber.Router) {
	route := app.Group("/user")

	route.Use(func(c *fiber.Ctx) error {
		s := c.Locals("session").(*session.BaseSession)
		c.Locals("service", &service.User{Session: s})
		return c.Next()
	})

	route.Get("/", getUsers)
	route.Post("/", createUser)
	route.Put("/:userId", updateUser)
	route.Delete("/:userId", deleteUser)
	route.Get("/:userId", getUser)
}

// getUsers : Will get all the user from the user management system in kayo web application
func getUsers(c *fiber.Ctx) error {
	page, limit := utils.GetPageAndLimit(c)
	_service := c.Locals("service").(*service.User)
	_user := _service.GetUsers(page, limit)
	return c.JSON(_user)
}

// createUser : create user in the user management system in kayo web application
func createUser(c *fiber.Ctx) error {
	return c.SendString("Hello, Community")
}

// updateUser : will update the user by userId in the user management system in kayo web application
func updateUser(c *fiber.Ctx) error {
	return c.SendString("Hello, Community")
}

// deleteUser : will delete user from user management system in kayo web application
func deleteUser(c *fiber.Ctx) error {
	return c.SendString("Hello, Community")
}

// getUser : will get user by userId from user management system in kayo web application
func getUser(c *fiber.Ctx) error {
	return c.SendString("Hello, Community")
}
