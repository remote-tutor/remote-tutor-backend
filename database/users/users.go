package database

import (
	dbInstance "backend/database"
	usersModel "backend/models/users"
	"fmt"
)

// GetUserByUsername searches the database for the user with the given username to be used in the login action
func GetUserByUsername(username string) usersModel.User {
	db := dbInstance.GetDBConnection()
	var user usersModel.User
	db.Where("username = ?", username).First(&user)
	return user
}

// GetUserByUserID searches the database for the user with the given userid
func GetUserByUserID(userid uint) usersModel.User {
	db := dbInstance.GetDBConnection()
	var user usersModel.User
	db.Where("id = ?", userid).First(&user)
	return user
}

// CreateUser inserts a new user to the database
func CreateUser(user *usersModel.User) {
	db := dbInstance.GetDBConnection()
	db.Create(user)
}

// UpdateUser updates the user information
func UpdateUser(user *usersModel.User) {
	dbInstance.GetDBConnection().Save(user)
}

// GetPendingUsers retrieve the non activated users from the database
func GetPendingUsers(sortBy []string, sortDesc []bool, page, itemsPerPage int, searchByValue, searchByField string) []usersModel.User {
	db := dbInstance.GetDBConnection()
	if searchByField == "username" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	} else if searchByField == "fullName" {
		db = db.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	}
	if sortBy != nil {
		for i := 0; i < len(sortBy); i++ {
			if sortDesc[i] {
				db.Order(sortBy[i] + " DESC")
			} else {
				db.Order(sortBy[i])
			}
		}
	}
	db = db.Offset((page - 1) * itemsPerPage).Limit(itemsPerPage)

	var pendingUsers []usersModel.User
	db.Where("activated = 0").Find(&pendingUsers)
	return pendingUsers
}

// GetTotalNumberOfPendingUsers returns the number of total pending users in the database
func GetTotalNumberOfPendingUsers(searchByValue, searchByField string) int64 {
	db := dbInstance.GetDBConnection()
	if searchByField == "username" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	} else if searchByField == "fullName" {
		db = db.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	}
	var count int64
	db.Model(&usersModel.User{}).Where("activated = 0").Count(&count)
	return count
}
