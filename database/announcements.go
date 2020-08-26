package database

import (
	md "backend/models"
	"fmt"
)

// CreateAnnouncement inserts a new user to the database
func CreateAnnouncement(announcement *md.Announcement) {
	GetDBConnection().Create(announcement)
}

// GetAnnouncements retrieves the announcements
func GetAnnouncements(title, topic, content string) []md.Announcement {
	var announcements []md.Announcement
	// to add the searched word inside '%' pairs, we use the Sprintf function
	// its normal use would be Sprintf("%s", variableName)
	// but as we need to escape the '%' character we put a pair of '%' to escape it
	GetDBConnection().Where("title LIKE ? AND topic LIKE ? AND content LIKE ?",
		fmt.Sprintf("%%%s%%", title), fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", content)).Order("created_at DESC").Find(&announcements)
	return announcements
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
