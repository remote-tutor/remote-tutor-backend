package quizzes

import (
	classesModel "backend/models/organizations"
	"time"

	"gorm.io/gorm"
)

// Quiz struct to store the quiz data
type Quiz struct {
	gorm.Model
	Title     string             `json:"title"`
	Year      int                `json:"year"`
	StartTime time.Time          `json:"startTime"`
	EndTime   time.Time          `json:"endTime"`
	TotalMark int                `json:"totalMark"`
	ClassHash string             `json:"classHash" gorm:"size:255"`
	Class     classesModel.Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
	Test      bool               `json:"test" gorm:"default:0"`
}
