package videos

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
	"time"
)

type UserWatch struct {
	UserID      uint            `json:"userID" gorm:"primary_key;autoIncrement:false"`
	User        usersModel.User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VideoPartID uint            `json:"videoPartID" gorm:"primary_key;autoIncrement:false"`
	VideoPart   VideoPart       `json:"videoPart" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartAt     time.Time       `json:"startAt"`
	ValidTill   time.Time       `json:"validTill"`
	DeletedAt   gorm.DeletedAt
}
