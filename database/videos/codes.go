package videos

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	codesDiagnostics "backend/diagnostics/database/videos"
	codesModel "backend/models/videos"
	"fmt"
	"gorm.io/gorm"
)

func GetCodesByVideo(paginationData *dbPagination.PaginationData, search string, videoID uint) ([]codesModel.Code, int64) {
	codes := make([]codesModel.Code, 0)
	query := dbInstance.GetDBConnection().Where("video_id = ?", videoID).
		Where("value LIKE ? OR used_by_user.full_name LIKE ? OR created_by_user.full_name LIKE ?",
			fmt.Sprintf("%s%%", search), fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search)).
		Joins("LEFT JOIN users AS used_by_user ON used_by_user_id = used_by_user.id").
		Joins("LEFT JOIN users AS created_by_user ON created_by_user_id = created_by_user.id")
	numberOfRecords := countCodes(query)
	query.Scopes(dbPagination.Paginate(paginationData)).
		Preload("UsedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, full_name") // id must be selected for a valid custom preloading
		}).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, full_name") // id must be selected for a valid custom preloading
		}).
		Find(&codes)
	return codes, numberOfRecords
}

func GenerateCodes(codes []codesModel.Code) error {
	err := dbInstance.GetDBConnection().Omit("used_by_user_id").Create(&codes).Error
	codesDiagnostics.WriteCodesErr(err, "Create", codes)
	return err
}

func countCodes(db *gorm.DB) int64 {
	totalCodes := int64(0)
	db.Model(&codesModel.Code{}).Count(&totalCodes)
	return totalCodes
}
