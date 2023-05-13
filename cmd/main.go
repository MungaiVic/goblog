package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Println("Hello bloggers!")
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeZone:   "Africa/Nairobi",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Hello there!")
	})
	log.Fatal(app.Listen(":3000"))
}
