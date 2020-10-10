package videos

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
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

func CreatePart(part *videoParts.VideoPart) error {
	err := dbInstance.GetDBConnection().Create(part).Error
	diagnostics.WriteError(err, "CreatePart")
	return err
}

func UpdatePart(part *videoParts.VideoPart) error {
	err := dbInstance.GetDBConnection().Save(part).Error
	diagnostics.WriteError(err, "UpdatePart")
	return err
}

func DeletePart(part *videoParts.VideoPart) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(part).Error
	diagnostics.WriteError(err, "DeletePart")
	return err
}