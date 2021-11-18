package route

import "github.com/gofiber/fiber/v2"

func RegisterAuth(app fiber.Router) {
	route := app.Group("/auth")
	route.Get("/login", signin)
	route.Post("/signup", signup)
}

// signin : will collect username and password and check if it is correct
func signin(c *fiber.Ctx) error {
	return c.SendString("Hello, Community")
}

// signup : get username and password and create a new user
func signup(c *fiber.Ctx) error {
	return c.SendString("Hello, Community")
}
