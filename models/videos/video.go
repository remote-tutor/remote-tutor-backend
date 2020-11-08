package videos

import (
	classesModel "backend/models/organizations"
	hashUtils "backend/utils/hash"
	"gorm.io/gorm"
	"os"
	"time"
)

type Video struct {
	gorm.Model
	AvailableFrom time.Time          `json:"availableFrom"`
	Parts         []VideoPart        `json:"parts"`
	Year          int                `json:"year"`
	Title         string             `json:"title"`
	Hash          string             `json:"hash" gorm:"size:25"`
	ClassHash     string             `json:"classHash" gorm:"size:25"`
	Class         classesModel.Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
}

// this function generates the hash then update the Class created
func (video *Video) AfterCreate(tx *gorm.DB) (err error) {
	hash := hashUtils.GenerateHash([]uint{video.ID}, os.Getenv("VIDEOS_SALT"))
	tx.Model(video).UpdateColumn("hash", hash)
	return
}
