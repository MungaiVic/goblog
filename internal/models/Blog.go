package models

import "gorm.io/gorm"

// TODO: Add blog_id as UUID
type Blog struct {
	gorm.Model
	Title   string
	Content string
	Author  User
}
