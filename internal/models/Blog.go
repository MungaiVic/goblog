package models

// TODO: Add blog_id as UUID
type Blog struct {
	Title   string
	Content string
	Author  User
}
