package service

import "github.com/gofiber/fiber/v2"

func GetAllBlogs(c *fiber.Ctx) error {
	return c.JSON("Here are all blogs")
}
func GetBlog(c *fiber.Ctx) error {
	return c.JSON("Here is a blog")
}
func CreateBlog(c *fiber.Ctx) error {
	return c.JSON("Here you create blogs")
}
func EditBlog(c *fiber.Ctx) error {
	return c.JSON("Here you edit blogs")
}
func DeleteBlog(c *fiber.Ctx) error {
	return c.JSON("Here you delete blogs")
}
