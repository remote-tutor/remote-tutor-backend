package organizations

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	classesModel "backend/models/organizations"
	"fmt"
	"gorm.io/gorm"
)

func GetAllClasses(paginationData *dbPagination.PaginationData, className, subject, teacherName string, year int, userID uint) ([]classesModel.Class, int64) {
	currentEnrollment := getClassesIDsByUser(userID)
	classes := make([]classesModel.Class, 0)
	query := dbInstance.GetDBConnection().Joins("JOIN organizations ON organizations.hash = classes.organization_hash").
		Where("name LIKE ? AND subject LIKE ? AND teacher_name LIKE ? AND year = ?",
		fmt.Sprintf("%%%s%%", className), fmt.Sprintf("%%%s%%", subject), fmt.Sprintf("%%%s%%", teacherName), year)
	if len(currentEnrollment) != 0 {
		query = query.Not("classes.hash IN ?", currentEnrollment)
	}
	numberOfRecords := countClasses(query)
	query.Scopes(dbPagination.Paginate(paginationData)).Preload("Organization").
		Order("classes.created_at DESC").Find(&classes)
	return classes, numberOfRecords
}

func countClasses(db *gorm.DB) int64 {
	totalClasses := int64(0)
	db.Model(&classesModel.Class{}).Count(&totalClasses)
	return totalClasses
}