package models

// TODO: Add user_id as UUID
type User struct {
	Fname    string
	Lname    string
	Email    string
	Password string
	IsActive bool
	Username string
}
