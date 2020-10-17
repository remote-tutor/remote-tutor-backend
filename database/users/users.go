package database

import (
	dbInstance "backend/database"
	classUsersDBInteractions "backend/database/organizations"
	"backend/diagnostics"
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
	diagnostics.WriteError(err, "database.log", "CreateUser")
	return err
}

// UpdateUser updates the user information
func UpdateUser(user *usersModel.User) error {
	err := dbInstance.GetDBConnection().Save(user).Error
	diagnostics.WriteError(err, "database.log", "UpdateUser")
	return err
}

// DeleteUser deletes the user from the database
func DeleteUser(user *usersModel.User) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(user).Error
	diagnostics.WriteError(err, "database.log", "DeleteUser")
	return err
}

func GetAdminUsers(userID uint) []usersModel.User {
	userClasses := classUsersDBInteractions.GetClassesHashesByUserID(userID)
	admins := make([]usersModel.User, 0)
	dbInstance.GetDBConnection().Joins("JOIN class_users ON users.id = class_users.user_id").
		Where("admin = 1 AND class_hash IN (?)", userClasses).Find(&admins)
	return admins
}
