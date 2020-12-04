package organizations

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	classUsersDiagnostics "backend/diagnostics/database/organizations"
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
	searchByValue, class string, pending bool) ([]classUsersModel.ClassUser, int64) {
	students := make([]classUsersModel.ClassUser, 0)
	query := dbInstance.GetDBConnection().Joins("User").
		Where("class_hash = ? AND activated = ?", class, !pending)
	query = query.Where("(username LIKE ? OR full_name LIKE ? OR phone_number LIKE ?)",
		fmt.Sprintf("%%%s%%", searchByValue), fmt.Sprintf("%%%s%%", searchByValue), fmt.Sprintf("%s%%", searchByValue))
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

func GetClassesHashesByUserID(userID uint) []string {
	hashes := make([]string, 0)
	dbInstance.GetDBConnection().Model(&classUsersModel.ClassUser{}).
		Where("user_id = ?", userID).Pluck("class_hash", &hashes)
	return hashes
}

func EnrollUser(classUser *classUsersModel.ClassUser) error {
	err := dbInstance.GetDBConnection().Create(classUser).Error
	classUsersDiagnostics.WriteClassUserErr(err, "Create", classUser)
	return err
}

func GetClassUserByID(id uint) classUsersModel.ClassUser {
	var classUser classUsersModel.ClassUser
	dbInstance.GetDBConnection().First(&classUser, id)
	return classUser
}

func GetClassUserByUserIDAndClass(userID uint, class string) classUsersModel.ClassUser {
	var classUser classUsersModel.ClassUser
	dbInstance.GetDBConnection().Where("user_id = ? AND class_hash = ?", userID, class).Find(&classUser)
	return classUser
}

func UpdateClassUser(classUser *classUsersModel.ClassUser) error {
	err := dbInstance.GetDBConnection().Save(classUser).Error
	classUsersDiagnostics.WriteClassUserErr(err, "Update", classUser)
	return err
}

func DeleteClassUser(classUser *classUsersModel.ClassUser) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(classUser).Error
	classUsersDiagnostics.WriteClassUserErr(err, "Delete", classUser)
	return err
}