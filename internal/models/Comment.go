package models

type Comment struct {
	Writer      User
	CommentBlog Blog
	Content     string
}
