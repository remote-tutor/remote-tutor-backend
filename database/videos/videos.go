package videos

import (
	dbInstance "backend/database"
	videosDiagnostics "backend/diagnostics/database/videos"
	videosModel "backend/models/videos"
	"time"
)

func GetVideosByClassAndMonthAndYear(class string, startOfMonth, endOfMonth time.Time) []videosModel.Video {
	videos := make([]videosModel.Video, 0)
	dbInstance.GetDBConnection().Where("class_hash = ? AND available_from >= ? AND available_from <= ?",
		class, startOfMonth, endOfMonth).Order("available_from").Find(&videos)
	return videos
}

func CreateVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Create(video).Error
	videosDiagnostics.WriteVideoErr(err, "Create", video)
	return err
}

func GetVideoByID(id uint) videosModel.Video {
	var video videosModel.Video
	dbInstance.GetDBConnection().First(&video, id)
	return video
}

func GetVideoByHash(hash string) videosModel.Video {
	var video videosModel.Video
	dbInstance.GetDBConnection().Where("hash = ?", hash).First(&video)
	return video
}

func UpdateVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Save(video).Error
	videosDiagnostics.WriteVideoErr(err, "Update", video)
	return err
}

func DeleteVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(video).Error
	videosDiagnostics.WriteVideoErr(err, "Delete", video)
	return err
}