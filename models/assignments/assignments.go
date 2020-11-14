package assignments

import (
	classesModel "backend/models/organizations"
	hashUtils "backend/utils/hash"
	"gorm.io/gorm"
	"os"
	"time"
)

// Assignment struct to hold the data of the assignment
type Assignment struct {
	gorm.Model
	Title             string             `json:"title"`
	Deadline          time.Time          `json:"deadline"`
	Questions         string             `json:"questions"`
	ModelAnswer       string             `json:"modelAnswer"`
	TotalMark         int                `json:"totalMark" gorm:"default:10"`
	ModelAnswerPeriod int                `json:"modelAnswerPeriod" gorm:"default:0"`
	ClassHash         string             `json:"classHash" gorm:"size:25"`
	Class             classesModel.Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
	Hash              string             `json:"hash" gorm:"size:25"`
}

// this function generates the hash then update the Class created
func (assignment *Assignment) AfterCreate(tx *gorm.DB) (err error) {
	hash := hashUtils.GenerateHash([]uint{assignment.ID}, os.Getenv("ASSIGNMENTS_SALT"))
	tx.Model(assignment).UpdateColumn("hash", hash)
	return
}
