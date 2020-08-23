package models
import (
	"github.com/jinzhu/gorm"
)

// User struct to store the announcement data
type Announcement struct {
	gorm.Model
	User  	User
	UserID 	uint
	Title 	string
	Topic 	string
	Content string
}