package videos

import (
	dbInstance "backend/database"
	watchesModel "backend/models/videos"
)

func CreateUserWatch(userWatch *watchesModel.UserWatch) {
	dbInstance.GetDBConnection().Create(userWatch)
}

func GetUserWatchByUserAndPart(userID, partID uint) watchesModel.UserWatch {
	var userWatch watchesModel.UserWatch
	dbInstance.GetDBConnection().Where("user_id = ? AND video_part_id = ?", userID, partID).Find(&userWatch)
	return userWatch
}