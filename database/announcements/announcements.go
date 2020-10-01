package database

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	announcementsModel "backend/models/announcements"
	"fmt"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// CreateAnnouncement inserts a new user to the database
func CreateAnnouncement(announcement *announcementsModel.Announcement) {
	dbInstance.GetDBConnection().Create(announcement)
}

// GetAnnouncementsByYear retrieves the announcements
func GetAnnouncementsByYear(c echo.Context, title, topic, content string, year int) ([]announcementsModel.Announcement, int64) {
	announcements := make([]announcementsModel.Announcement, 0)
	// to add the searched word inside '%' pairs, we use the Sprintf function
	// its normal use would be Sprintf("%s", variableName)
	// but as we need to escape the '%' character we put a pair of '%' to escape it
	query := dbInstance.GetDBConnection().Where("title LIKE ? AND topic LIKE ? AND content LIKE ? AND year = ?",
		fmt.Sprintf("%%%s%%", title), fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", content), year)
	numberOfRecords := countAnnouncements(query)
	query = query.Scopes(dbPagination.Paginate(c)).Find(&announcements)
	return announcements, numberOfRecords
}

// GetAnnouncementByID retrieves the announcement by the announcement id
func GetAnnouncementByID(id uint) announcementsModel.Announcement {
	var announcement announcementsModel.Announcement
	dbInstance.GetDBConnection().First(&announcement, id)
	return announcement
}

// UpdateAnnouncement updates the announcement
func UpdateAnnouncement(announcement *announcementsModel.Announcement) {
	dbInstance.GetDBConnection().Save(announcement)
}

// DeleteAnnouncement deletes the announcement
func DeleteAnnouncement(announcement *announcementsModel.Announcement) {
	dbInstance.GetDBConnection().Unscoped().Delete(announcement)
}

// countAnnouncements counts the number of records in the database by specific search
func countAnnouncements(db *gorm.DB) int64 {
	totalAnnouncements := int64(0)
	db.Model(&announcementsModel.Announcement{}).Count(&totalAnnouncements)
	return totalAnnouncements
}
