package assignments

import (
	"gorm.io/gorm"
	"time"
)

// Assignment struct to hold the data of the assignment
type Assignment struct {
	gorm.Model
	Title       string    `json:"title"`
	Year        int       `json:"year"`
	Deadline    time.Time `json:"deadline"`
	Questions   string    `json:"questions"`
	ModelAnswer string    `json:"modelAnswer"`
	TotalMark   int       `json:"totalMark" gorm:"default:10"`
}
