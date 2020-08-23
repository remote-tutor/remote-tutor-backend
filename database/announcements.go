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