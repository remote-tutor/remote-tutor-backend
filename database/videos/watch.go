package videos

import (
	dbInstance "backend/database"
	"backend/diagnostics"
	watchesModel "backend/models/videos"
)

func CreateUserWatch(userWatch *watchesModel.UserWatch) error {
	err := dbInstance.GetDBConnection().Create(userWatch).Error
	diagnostics.WriteError(err, "database.log", "CreateUserWatch")
	return err
}

func GetUserWatchByUserAndPart(userID, partID uint) watchesModel.UserWatch {
	var userWatch watchesModel.UserWatch
	dbInstance.GetDBConnection().Where("user_id = ? AND video_part_id = ?", userID, partID).Find(&userWatch)
	return userWatch
}