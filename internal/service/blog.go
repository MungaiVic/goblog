package service

import "github.com/gofiber/fiber/v2"

type blogImpl struct {
}

func NewBlogService() BlogService {
	return &blogImpl{}
}

func (blog blogImpl) GetAllBlogs(c *fiber.Ctx) error {
	return c.JSON("Here are all blogs")
}
func (blog blogImpl) GetBlog(c *fiber.Ctx) error {
	return c.JSON("Here is a blog")
}
func (blog blogImpl) CreateBlog(c *fiber.Ctx) error {
	return c.JSON("Here you create blogs")
}
func (blog blogImpl) EditBlog(c *fiber.Ctx) error {
	return c.JSON("Here you edit blogs")
}
func (blog blogImpl) DeleteBlog(c *fiber.Ctx) error {
	return c.JSON("Here you delete blogs")
}
