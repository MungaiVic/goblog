package handlers

import (
	"goblog/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupCommentRoutes(group *fiber.Group, svc service.CommentService) {
	commentRoutes := group.Group("/comments")

	commentRoutes.Post("/:blog_id", func(c *fiber.Ctx) error {
		return svc.CreateComment(c)
	})
	commentRoutes.Patch("/:id", func(c *fiber.Ctx) error {
		return svc.EditComment(c)
	})
	commentRoutes.Delete("/:id", func(c *fiber.Ctx) error {
		return svc.DeleteComment(c)
	})
}
