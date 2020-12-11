package videos

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	videosDiagnostics "backend/diagnostics/database/videos"
	classesModel "backend/models/organizations"
	videosModel "backend/models/videos"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func GetVideosByClassAndMonthAndYear(class string, startOfMonth, endOfMonth time.Time) []videosModel.Video {
	videos := make([]videosModel.Video, 0)
	dbInstance.GetDBConnection().Where("class_hash = ? AND available_from >= ? AND available_from <= ?",
		class, startOfMonth, endOfMonth).Order("available_from").Find(&videos)
	return videos
}

func GetNonAccessedStudents(paginationData *dbPagination.PaginationData, classHash, search string, videoID uint) ([]classesModel.ClassUser, int64) {
	students := make([]classesModel.ClassUser, 0)
	subQuery := dbInstance.GetDBConnection().Select("used_by_user_id").
		Where("video_id = ? AND used_by_user_id IS NOT NULL", videoID).Table("codes")
	query := dbInstance.GetDBConnection().Where("class_hash = ? AND admin = 0 AND activated = 1", classHash).
		Where("(users.username LIKE ? OR users.full_name LIKE ? OR users.phone_number LIKE ?)",
			fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%s%%", search)).
		Where("user_id NOT IN (?)", subQuery).
		Joins("JOIN users ON users.id = class_users.user_id")
		//Where("codes.used_by_user_id IS NULL")
	numberOfRecords := countClassUsers(query)
	query.Scopes(dbPagination.Paginate(paginationData)).
		Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password") // omit 'password' column
	}).Find(&students)
	return students, numberOfRecords
}

func CreateVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Create(video).Error
	videosDiagnostics.WriteVideoErr(err, "Create", video)
	return err
}

func GetVideoByID(id uint) videosModel.Video {
	var video videosModel.Video
	dbInstance.GetDBConnection().First(&video, id)
	return video
}

func GetVideoByPartID(id uint) videosModel.Video {
	var video videosModel.Video
	subQuery := dbInstance.GetDBConnection().Select("video_id").
		Where("id = ?", id).Table("video_parts")
	dbInstance.GetDBConnection().Where("id = (?)", subQuery).First(&video)
	return video
}

func GetVideoByHash(hash string) videosModel.Video {
	var video videosModel.Video
	dbInstance.GetDBConnection().Where("hash = ?", hash).First(&video)
	return video
}

func UpdateVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Save(video).Error
	videosDiagnostics.WriteVideoErr(err, "Update", video)
	return err
}

func DeleteVideo(video *videosModel.Video) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(video).Error
	videosDiagnostics.WriteVideoErr(err, "Delete", video)
	return err
}

func countClassUsers(db *gorm.DB) int64 {
	totalClassUsers := int64(0)
	db.Model(&classesModel.ClassUser{}).Count(&totalClassUsers)
	return totalClassUsers
}