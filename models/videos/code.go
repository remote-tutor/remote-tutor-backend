package videos

import (
	usersModel "backend/models/users"
	"time"
)

type Code struct {
	CreatedAt       time.Time
	Value           string          `json:"value" gorm:"primaryKey"`
	Video           Video           `json:"video" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VideoID         uint            `json:"videoID" gorm:"uniqueIndex:idx_video_user"`
	CreatedByUser   usersModel.User `json:"createdByUser" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedByUserID uint            `json:"createdByUserID"`
	UsedByUser      usersModel.User `json:"usedByUser" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UsedByUserID    uint            `json:"usedByUserID" gorm:"uniqueIndex:idx_video_user"`
	Manual          bool            `json:"manual" gorm:"default:0"`
}
