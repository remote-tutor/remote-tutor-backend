package database

import (
	md "backend/models"
)

// GetUserByUsername searches the database for the user with the given username to be used in the login action
func GetUserByUsername(username string) md.User {
	db := GetDBConnection()
	var user md.User
	db.Where("username = ?", username).First(&user)
	return user
}

// CreateUser inserts a new user to the database
func CreateUser(user *md.User) {
	db := GetDBConnection()
	db.Create(user)
}
