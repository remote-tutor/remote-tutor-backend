package organizations

import (
	db "backend/database"
	organizationModel "backend/models/organizations"
)

func GetClassesByUser(userID uint) []organizationModel.UserClass {
	userClasses := make([]organizationModel.UserClass, 0)
	db.GetDBConnection().Where("user_id = ?", userID).Preload("Class.Organization").Find(&userClasses)
	return userClasses
}
