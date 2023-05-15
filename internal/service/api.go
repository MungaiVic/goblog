package service

import (
	"github.com/gofiber/fiber/v2"
)

type BlogService interface {
	GetAllBlogs(c *fiber.Ctx) error
	GetBlog(c *fiber.Ctx) error
	CreateBlog(c *fiber.Ctx) error
	EditBlog(c *fiber.Ctx)
	DeleteBlog(c *fiber.Ctx) error
}

type UserService interface {
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type CommentService interface {
	GetComments(c *fiber.Ctx) error
	CreateComment(c *fiber.Ctx) error
	EditComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
}
