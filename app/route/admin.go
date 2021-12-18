package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workfoxes/calypso/utils"
	"github.com/workfoxes/calypso/utils/response"
	"github.com/workfoxes/kayo/app/session"
	"github.com/workfoxes/kayo/internal/model"
	"net/http"
	"strconv"
)

func RegisterUser(app fiber.Router) {
	route := app.Group("/user")

	route.Get("/", getUsers)
	route.Post("/", createUser)
	route.Put("/:userId", updateUser)
	route.Delete("/:userId", deleteUser)
	route.Get("/:userId", getUser)
}

// getUsers : Will get all the user from the user management system in kayo web application
func getUsers(c *fiber.Ctx) error {
	var users []*model.User
	var count *int64
	s := c.Locals("session").(*session.BaseSession)

	page, limit := utils.GetPageAndLimit(c)
	s.TX.Find(&users).Count(count).Limit(int(limit)).Offset(int(page * limit))
	c.Response().Header.Set("X-Total-Count", strconv.FormatInt(*count, 10))
	return c.JSON(users)
}

// createUser : create user in the user management system in kayo web application
func createUser(c *fiber.Ctx) error {
	var user model.User
	s := c.Locals("session").(*session.BaseSession)
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	s.TX.Create(&user)
	c.Response().SetStatusCode(http.StatusCreated)
	return response.SuccessResponse(c, user, "User created successfully")
}

// updateUser : will update the user by userId in the user management system in kayo web application
func updateUser(c *fiber.Ctx) error {
	var user model.User
	s := c.Locals("session").(*session.BaseSession)

	userId := c.Params("userId")
	u64, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	s.TX.First(&model.User{}, uint(u64)).Updates(&user)
	return response.SuccessResponse(c, user, "User Updated successfully")
}

// deleteUser : will delete user from user management system in kayo web application
func deleteUser(c *fiber.Ctx) error {
	s := c.Locals("session").(*session.BaseSession)

	userId := c.Params("userId")
	u64, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return err
	}
	s.TX.Delete(&model.User{}, uint(u64))
	c.Response().SetStatusCode(http.StatusAccepted)
	return response.SuccessResponse(c, &model.User{}, "User Deleted")
}

// getUser : will get user by userId from user management system in kayo web application
func getUser(c *fiber.Ctx) error {
	var user model.User
	s := c.Locals("session").(*session.BaseSession)

	userId := c.Params("userId")
	u64, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return err
	}
	s.TX.First(&user, uint(u64))
	return c.JSON(user)
}
