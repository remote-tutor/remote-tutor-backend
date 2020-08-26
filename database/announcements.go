package database

import (
	md "backend/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

// CreateAnnouncement inserts a new user to the database
func CreateAnnouncement(announcement *md.Announcement) {
	GetDBConnection().Create(announcement)
}

// GetAnnouncements retrieves the announcements
func GetAnnouncements(title, topic, content string, length, currentPage int) ([]md.Announcement, int) {
	var announcements []md.Announcement
	// to add the searched word inside '%' pairs, we use the Sprintf function
	// its normal use would be Sprintf("%s", variableName)
	// but as we need to escape the '%' character we put a pair of '%' to escape it
	query := GetDBConnection().Where("title LIKE ? AND topic LIKE ? AND content LIKE ?",
		fmt.Sprintf("%%%s%%", title), fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", content))
	numberOfRecords := countAnnouncements(query)
	query.Offset(length * (currentPage - 1)).Limit(length).
		Order("created_at DESC").Find(&announcements)
	return announcements, numberOfRecords
}

// GetAnnouncementById retrieves the announcement by the announcement id
func GetAnnouncementById(id uint) md.Announcement {
	var announcement md.Announcement
	GetDBConnection().First(&announcement, id)
	return announcement
}

// SaveAnnouncement saves the announcement
func SaveAnnouncement(announcement *md.Announcement) {
	GetDBConnection().Save(announcement)
}

// SaveAnnouncement saves the announcement
func DeleteAnnouncement(announcement *md.Announcement) {
	GetDBConnection().Delete(announcement)
}

// countAnnouncements counts the number of records in the database by specific search
func countAnnouncements(db *gorm.DB) int {
	totalAnnouncements := 0
	db.Model(&md.Announcement{}).Count(&totalAnnouncements)
	return totalAnnouncements
}
