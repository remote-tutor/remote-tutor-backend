package database

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	usersModel "backend/models/users"
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
	diagnostics.WriteError(err, "CreateUser")
	return err
}

// UpdateUser updates the user information
func UpdateUser(user *usersModel.User) error {
	err := dbInstance.GetDBConnection().Save(user).Error
	diagnostics.WriteError(err, "UpdateUser")
	return err
}

// DeleteUser deletes the user from the database
func DeleteUser(user *usersModel.User) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(user).Error
	diagnostics.WriteError(err, "DeleteUser")
	return err
}

func GetAdminUsers() []usersModel.User {
	admins := make([]usersModel.User, 0)
	dbInstance.GetDBConnection().Where("admin = 1").Find(&admins)
	return admins
}