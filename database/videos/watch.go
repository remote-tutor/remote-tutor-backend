package videos

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	videosDiagnostics "backend/diagnostics/database/videos"
	watchesModel "backend/models/videos"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateUserWatch(userWatch *watchesModel.UserWatch) error {
	// try to create a nw user watch, if it fails (record already there but deleted), then set deleted at to null
	err := dbInstance.GetDBConnection().Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"deleted_at"}),
	}).FirstOrCreate(userWatch).Error
	videosDiagnostics.WriteWatchErr(err, "Create", userWatch)
	return err
}

func DeleteUserWatch(userWatch *watchesModel.UserWatch) error {
	err := dbInstance.GetDBConnection().Delete(userWatch).Error // soft delete (set deleted at record)
	videosDiagnostics.WriteWatchErr(err, "Delete", userWatch)
	return err
}

func GetPartWatchesForAllUsers(partID uint, paginationData *dbPagination.PaginationData) ([]watchesModel.UserWatch, int64) {
	userWatches := make([]watchesModel.UserWatch, 0)
	query := dbInstance.GetDBConnection().Where("video_part_id = ?", partID)
	total := countPartWatches(query)
	query.Scopes(dbPagination.Paginate(paginationData)).
		Joins("User").Joins("VideoPart").Find(&userWatches)
	return userWatches, total
}

func GetUserWatchByUserAndPart(userID, partID uint) watchesModel.UserWatch {
	var userWatch watchesModel.UserWatch
	dbInstance.GetDBConnection().Where("user_id = ? AND video_part_id = ?", userID, partID).Find(&userWatch)
	return userWatch
}

func countPartWatches(db *gorm.DB) int64 {
	totalWatches := int64(0)
	db.Model(&watchesModel.UserWatch{}).Count(&totalWatches)
	return totalWatches
}
