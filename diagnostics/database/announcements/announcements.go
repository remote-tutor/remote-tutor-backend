package announcements

import (
	"backend/diagnostics"
	announcementsModel "backend/models/announcements"
)

func WriteAnnouncementErr(err error, errorType string, announcement *announcementsModel.Announcement) {
	filepath := "database/announcements/announcements.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, announcement)
}
