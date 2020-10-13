package organizations

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	dbPagination "backend/database/scopes"
	classUsersModel "backend/models/organizations"
	"fmt"
	"gorm.io/gorm"
)

func GetClassesByUser(userID uint) []classUsersModel.ClassUser {
	classUsers := make([]classUsersModel.ClassUser, 0)
	dbInstance.GetDBConnection().Where("user_id = ?", userID).Preload("Class.Organization").Find(&classUsers)
	return classUsers
}

func GetStudentsByClass(paginationData *dbPagination.PaginationData,
	searchByValue, searchByField, class string, pending bool) ([]classUsersModel.ClassUser, int64) {
	students := make([]classUsersModel.ClassUser, 0)
	query := dbInstance.GetDBConnection().Joins("User").
		Where("class_hash = ? AND activated = ?", class, !pending)
	if searchByField == "username" {
		query = query.Where("username LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	} else if searchByField == "fullName" {
		query = query.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", searchByValue))
	}
	numberOfRecords := countClassStudents(query)
	query = query.Scopes(dbPagination.Paginate(paginationData))
	query.Scopes(dbPagination.Paginate(paginationData)).Find(&students)
	return students, numberOfRecords
}

func countClassStudents(db *gorm.DB) int64 {
	totalClasses := int64(0)
	db.Model(&classUsersModel.ClassUser{}).Count(&totalClasses)
	return totalClasses
}

func getClassesIDsByUser(userID uint) []string {
	hashes := make([]string, 0)
	dbInstance.GetDBConnection().Model(&classUsersModel.ClassUser{}).
		Where("user_id = ?", userID).Pluck("class_hash", &hashes)
	return hashes
}

func EnrollUser(classUser *classUsersModel.ClassUser) error {
	err := dbInstance.GetDBConnection().Create(classUser).Error
	diagnostics.WriteError(err, "EnrollUser")
	return err
}

func GetClassUserByID(id uint) classUsersModel.ClassUser {
	var classUser classUsersModel.ClassUser
	dbInstance.GetDBConnection().First(&classUser, id)
	return classUser
}

func UpdateClassUser(classUser *classUsersModel.ClassUser) error {
	err := dbInstance.GetDBConnection().Save(classUser).Error
	diagnostics.WriteError(err, "UpdateClassUser")
	return err
}

func DeleteClassUser(classUser *classUsersModel.ClassUser) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(classUser).Error
	diagnostics.WriteError(err, "DeleteClassUser")
	return err
}