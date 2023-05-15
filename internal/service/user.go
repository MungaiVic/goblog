package service

import "github.com/gofiber/fiber/v2"

type userImpl struct{}

func NewUserService() UserService {
	return &userImpl{}
}

func (user userImpl) GetUsers(c *fiber.Ctx) error {
	return c.JSON("Get users")
}
func (user userImpl) GetUser(c *fiber.Ctx) error {
	return c.JSON("Get User")
}
func (user userImpl) Register(c *fiber.Ctx) error {
	return c.JSON("Create/Register User")
}
func (user userImpl) ChangePassword(c *fiber.Ctx) error {
	return c.JSON("change User Password")
}
func (user userImpl) UpdateUser(c *fiber.Ctx) error {
	return c.JSON("Update User details")
}
func (user userImpl) DeleteUser(c *fiber.Ctx) error {
	return c.JSON("Delete User")
}
