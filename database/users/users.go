package database

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	usersModel "backend/models/users"
	"fmt"

	"github.com/labstack/echo"
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
func CreateUser(user *usersModel.User) error {
	err := dbInstance.GetDBConnection().Create(user).Error
	return err
}

// UpdateUser updates the user information
func UpdateUser(user *usersModel.User) error {
	err := dbInstance.GetDBConnection().Save(user).Error
	return err
}

// DeleteUser deletes the user from the database
func DeleteUser(user *usersModel.User) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(user).Error
	return err
}

// GetUsers retrieve the non activated users from the database
func GetUsers(c echo.Context, searchByValue, searchByField string, pending bool) []usersModel.User {
	db := dbInstance.GetDBConnection()
	if searchByField == "username" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	} else if searchByField == "fullName" {
		db = db.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	}
	db = db.Scopes(dbPagination.Paginate(c))
	pendingUsers := make([]usersModel.User, 0)
	db.Where("activated = ?", !pending).Find(&pendingUsers)
	return pendingUsers
}

// GetTotalNumberOfUsers returns the number of total pending users in the database
func GetTotalNumberOfUsers(searchByValue, searchByField string, pending bool) int64 {
	db := dbInstance.GetDBConnection()
	if searchByField == "username" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	} else if searchByField == "fullName" {
		db = db.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	}
	var count int64
	db.Model(&usersModel.User{}).Where("activated = ?", !pending).Count(&count)
	return count
}

func GetAdminUsers() []usersModel.User {
	admins := make([]usersModel.User, 0)
	dbInstance.GetDBConnection().Where("admin = 1").Find(&admins)
	return admins
}