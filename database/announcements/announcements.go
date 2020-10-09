package database

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	dbPagination "backend/database/scopes"
	announcementsModel "backend/models/announcements"
	"fmt"

	"gorm.io/gorm"
)

// CreateAnnouncement inserts a new user to the database
func CreateAnnouncement(announcement *announcementsModel.Announcement) error {
	err := dbInstance.GetDBConnection().Create(announcement).Error
	diagnostics.WriteError(err, "CreateAnnouncement")
	return err
}

// GetAnnouncementsByYear retrieves the announcements
func GetAnnouncementsByYear(paginationData *dbPagination.PaginationData, title, topic, content string, year int) ([]announcementsModel.Announcement, int64) {
	announcements := make([]announcementsModel.Announcement, 0)
	// to add the searched word inside '%' pairs, we use the Sprintf function
	// its normal use would be Sprintf("%s", variableName)
	// but as we need to escape the '%' character we put a pair of '%' to escape it
	query := dbInstance.GetDBConnection().Where("title LIKE ? AND topic LIKE ? AND content LIKE ? AND year = ?",
		fmt.Sprintf("%%%s%%", title), fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", content), year)
	numberOfRecords := countAnnouncements(query)
	query = query.Scopes(dbPagination.Paginate(paginationData)).Find(&announcements)
	return announcements, numberOfRecords
}

// GetAnnouncementByID retrieves the announcement by the announcement id
func GetAnnouncementByID(id uint) announcementsModel.Announcement {
	var announcement announcementsModel.Announcement
	dbInstance.GetDBConnection().First(&announcement, id)
	return announcement
}

// UpdateAnnouncement updates the announcement
func UpdateAnnouncement(announcement *announcementsModel.Announcement) error {
	err := dbInstance.GetDBConnection().Save(announcement).Error
	diagnostics.WriteError(err, "UpdateAnnouncement")
	return err
}

// DeleteAnnouncement deletes the announcement
func DeleteAnnouncement(announcement *announcementsModel.Announcement) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(announcement).Error
	diagnostics.WriteError(err, "DeleteAnnouncement")
	return err
}

// countAnnouncements counts the number of records in the database by specific search
func countAnnouncements(db *gorm.DB) int64 {
	totalAnnouncements := int64(0)
	db.Model(&announcementsModel.Announcement{}).Count(&totalAnnouncements)
	return totalAnnouncements
}
