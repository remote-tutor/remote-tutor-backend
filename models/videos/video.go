package videos

import (
	classesModel "backend/models/organizations"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	AvailableFrom time.Time          `json:"availableFrom"`
	Parts         []VideoPart        `json:"parts"`
	Year          int                `json:"year"`
	Title         string             `json:"title"`
	ClassHash     string             `json:"classHash" gorm:"size:25"`
	Class         classesModel.Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
}
