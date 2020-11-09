package quizzes

import (
	classesModel "backend/models/organizations"
	hashUtils "backend/utils/hash"
	"os"
	"time"

	"gorm.io/gorm"
)

// Quiz struct to store the quiz data
type Quiz struct {
	gorm.Model
	Title     string             `json:"title"`
	StartTime time.Time          `json:"startTime"`
	EndTime   time.Time          `json:"endTime"`
	TotalMark int                `json:"totalMark"`
	ClassHash string             `json:"classHash" gorm:"size:25"`
	Class     classesModel.Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
	Test      bool               `json:"test" gorm:"default:0"`
	Hash      string             `json:"hash" gorm:"size:25"`
}

// this function generates the hash then update the Quiz created
func (quiz *Quiz) AfterCreate(tx *gorm.DB) (err error) {
	hash := hashUtils.GenerateHash([]uint{quiz.ID}, os.Getenv("QUIZZES_SALT"))
	tx.Model(quiz).UpdateColumn("hash", hash)
	return
}