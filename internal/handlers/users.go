package handlers

import (
	"goblog/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(group *fiber.Group, svc service.UserService) {
	userRoutes := group.Group("/users")

	userRoutes.Get("/", func(c *fiber.Ctx) error {
		return svc.GetUsers(c)
	})
	userRoutes.Get("/:id", func(c *fiber.Ctx) error {
		return svc.GetUser(c)
	})

	userRoutes.Post("/", func(c *fiber.Ctx) error {
		return svc.Register(c)
	})

	userRoutes.Patch("/:id", func(c *fiber.Ctx) error {
		return svc.UpdateUser(c)
	})

	userRoutes.Delete("/:id", func(c *fiber.Ctx) error {
		return svc.DeleteUser(c)
	})
}
