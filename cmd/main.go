package main

import (
	"fmt"
	"goblog/internal/handlers"
	"goblog/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Println("Hello bloggers!")
	app := fiber.New()
	v1 := app.Group("api/v1").(*fiber.Group)

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeZone:   "Africa/Nairobi",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	// Setup services
	blogSVC := service.NewBlogService()

	// Set up routes
	handlers.SetupBlogRoutes(v1, blogSVC)
	log.Fatal(app.Listen(":3000"))
}
