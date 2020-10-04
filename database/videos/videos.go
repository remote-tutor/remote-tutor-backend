package videos

import (
	dbInstance "backend/database"
	videosModel "backend/models/videos"
	"time"
)

func GetVideosByMonthAndYear(year int, startOfMonth, endOfMonth time.Time) []videosModel.Video {
	videos := make([]videosModel.Video, 0)
	dbInstance.GetDBConnection().Where("year = ? AND available_from >= ? AND available_from <= ?",
		year, startOfMonth, endOfMonth).Order("available_from").Find(&videos)
	return videos
}

func CreateVideo(video *videosModel.Video) {
	dbInstance.GetDBConnection().Create(video)
}

func GetVideoByID(id uint) videosModel.Video {
	var video videosModel.Video
	dbInstance.GetDBConnection().First(&video, id)
	return video
}

func UpdateVideo(video *videosModel.Video) {
	dbInstance.GetDBConnection().Save(video)
}