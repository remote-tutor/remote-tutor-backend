package quizzes

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Quiz struct to store the quiz data
type Quiz struct {
	gorm.Model
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
