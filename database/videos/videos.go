package videos

import (
	dbInstance "backend/database"
	videosModel "backend/models/videos"
)

func CreateVideo(video *videosModel.Video) {
	dbInstance.GetDBConnection().Create(video)
}
