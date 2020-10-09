package videos

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	videosModel "backend/models/videos"
	"time"
)

func GetVideosByMonthAndYear(year int, startOfMonth, endOfMonth time.Time) []videosModel.Video {
	videos := make([]videosModel.Video, 0)
	dbInstance.GetDBConnection().Where("year = ? AND available_from >= ? AND available_from <= ?",
		year, startOfMonth, endOfMonth).Order("available_from").Find(&videos)
	return videos
}

func CreateVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Create(video).Error
	diagnostics.WriteError(err, "CreateVideo")
	return err
}

func GetVideoByID(id uint) videosModel.Video {
	var video videosModel.Video
	dbInstance.GetDBConnection().First(&video, id)
	return video
}

func UpdateVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Save(video).Error
	diagnostics.WriteError(err, "UpdateVideo")
	return err
}

func DeleteVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(video).Error
	diagnostics.WriteError(err, "DeleteVideo")
	return err
}