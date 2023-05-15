package models

import "gorm.io/gorm"

// TODO: Add comment_id as UUID
type Comment struct {
	gorm.Model
	Writer      User
	CommentBlog Blog
	Content     string
}
