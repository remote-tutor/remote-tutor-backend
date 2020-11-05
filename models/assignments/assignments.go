package assignments

import (
	classesModel "backend/models/organizations"
	"gorm.io/gorm"
	"time"
)

// Assignment struct to hold the data of the assignment
type Assignment struct {
	gorm.Model
	Title             string             `json:"title"`
	Year              int                `json:"year"`
	Deadline          time.Time          `json:"deadline"`
	Questions         string             `json:"questions"`
	ModelAnswer       string             `json:"modelAnswer"`
	TotalMark         int                `json:"totalMark" gorm:"default:10"`
	ModelAnswerPeriod int                `json:"modelAnswerPeriod" gorm:"default:0"`
	ClassHash         string             `json:"classHash" gorm:"size:25"`
	Class             classesModel.Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
}
