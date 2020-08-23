package database

import (
	md "backend/models"
)

// CreateAnnouncement inserts a new user to the database
func CreateAnnouncement(announcement *md.Announcement) {
	db := GetDBConnection()
	db.Create(announcement)
}

// GetAnnouncements retrieves the announcements
func GetAnnouncements() []md.Announcement {
	db := GetDBConnection()

	var announcements []md.Announcement
	db.Find(&announcements)
	return announcements
}

// GetAnnouncementById retrieves the announcement by the announcement id
func GetAnnouncementById(id uint) md.Announcement {
	db := GetDBConnection()

	var announcement md.Announcement
	db.First(&announcement, id)
	return announcement
}

// SaveAnnouncement saves the announcement
func SaveAnnouncement(announcement *md.Announcement) {
	db := GetDBConnection()
	db.Save(announcement)
}