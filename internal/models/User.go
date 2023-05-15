package models

import "gorm.io/gorm"

// TODO: Add user_id as UUID
type User struct {
	gorm.Model
	Fname    string
	Lname    string
	Email    string
	Password string
	IsActive bool
	Username string
}
