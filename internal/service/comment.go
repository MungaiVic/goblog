package service

import "github.com/gofiber/fiber/v2"

type commentImpl struct{}

func NewCommentService() CommentService {
	return &commentImpl{}
}

func (comment commentImpl) CreateComment(c *fiber.Ctx) error {
	return c.JSON("Add comment")
}
func (comment commentImpl) EditComment(c *fiber.Ctx) error {
	return c.JSON("Edit comment")
}
func (comment commentImpl) DeleteComment(c *fiber.Ctx) error {
	return c.JSON("Delete comment")
}
