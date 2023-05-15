package handlers

import (
	"goblog/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupBlogRoutes(group *fiber.Group, svc service.BlogService){
	blogRoutes := group.Group("/blog")

	blogRoutes.Get("/", func(c *fiber.Ctx) error {
		return svc.GetAllBlogs(c)
	})
}