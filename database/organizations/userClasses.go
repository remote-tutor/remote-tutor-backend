package organizations

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	classUsersModel "backend/models/organizations"
)

func GetClassesByUser(userID uint) []classUsersModel.ClassUser {
	userClasses := make([]classUsersModel.ClassUser, 0)
	dbInstance.GetDBConnection().Where("user_id = ?", userID).Preload("Class.Organization").Find(&userClasses)
	return userClasses
}

func getClassesIDsByUser(userID uint) []string {
	hashes := make([]string, 0)
	dbInstance.GetDBConnection().Model(&classUsersModel.ClassUser{}).
		Where("user_id = ?", userID).Pluck("class_hash", &hashes)
	return hashes
}


func EnrollUser(classUser *classUsersModel.ClassUser) error {
	err := dbInstance.GetDBConnection().Create(classUser).Error
	diagnostics.WriteError(err, "EnrollUser")
	return err
}