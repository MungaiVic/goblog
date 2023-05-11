package models

// TODO: Add comment_id as UUID
type Comment struct {
	Writer      User
	CommentBlog Blog
	Content     string
}
