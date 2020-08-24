package database

import (
	md "backend/models"
)

// CreateAnnouncement inserts a new user to the database
func CreateAnnouncement(announcement *md.Announcement) {
	GetDBConnection().Create(announcement)
}

// GetAnnouncements retrieves the announcements
func GetAnnouncements() []md.Announcement {
	var announcements []md.Announcement
	GetDBConnection().Find(&announcements)
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