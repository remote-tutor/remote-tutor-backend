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

// GetPendingUsers retrieve the non activated users from the database
func GetPendingUsers(sortBy []string, sortDesc []bool, page, itemsPerPage int) []md.User {
	db := GetDBConnection()
	if sortBy != nil {
		for i := 0; i < len(sortBy); i++ {
			if sortDesc[i] {
				db.Order(sortBy[i] + " DESC")
			} else {
				db.Order(sortBy[i])
			}
		}
	}
	db = db.Offset((page - 1) * itemsPerPage)

	var pendingUsers []md.User
	db.Where("activated = 0").Find(&pendingUsers)
	return pendingUsers
}

// GetTotalNumberOfPendingUsers returns the number of total pending users in the database
func GetTotalNumberOfPendingUsers() int {
	db := GetDBConnection()
	var count int
	db.Model(&md.User{}).Count(&count)
	return count
}
