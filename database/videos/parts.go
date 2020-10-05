package videos

import (
	dbInstance "backend/database"
	videoParts "backend/models/videos"
)

func GetPartsByVideo(videoID uint) []videoParts.VideoPart {
	parts := make([]videoParts.VideoPart, 0)
	dbInstance.GetDBConnection().Where("video_id = ?", videoID).Order("number").Find(&parts)
	return parts
}

func GetPartByID(id uint) videoParts.VideoPart {
	var part videoParts.VideoPart
	dbInstance.GetDBConnection().First(&part, id)
	return part
}

func CreatePart(part *videoParts.VideoPart) {
	dbInstance.GetDBConnection().Create(part)
}

func UpdatePart(part *videoParts.VideoPart) {
	dbInstance.GetDBConnection().Save(part)
}

func DeletePart(part *videoParts.VideoPart) {
	dbInstance.GetDBConnection().Unscoped().Delete(part)
}