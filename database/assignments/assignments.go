package assignments

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	assignmentsDiagnostics "backend/diagnostics/database/assignments"
	assignmentsModel "backend/models/assignments"
	"gorm.io/gorm"
)

// CreateAssignment inserts a new assignment to the database
func CreateAssignment(assignment *assignmentsModel.Assignment) error {
	err := dbInstance.GetDBConnection().Create(assignment).Error
	assignmentsDiagnostics.WriteAssignmentErr(err, "Create", assignment)
	return err
}

// GetAssignmentsByClass gets an array of assignments to display to the user
func GetAssignmentsByClass(paginationData *dbPagination.PaginationData, class string) ([]assignmentsModel.Assignment, int64) {
	assignments := make([]assignmentsModel.Assignment, 0)
	db := dbInstance.GetDBConnection().Where("class_hash = ?", class)
	totalAssignments := countAssignments(db)
	db.Scopes(dbPagination.Paginate(paginationData)).Find(&assignments)
	return assignments, totalAssignments
}

// countAssignments counts the total number of assignments for a specific user (year)
func countAssignments(db *gorm.DB) int64 {
	totalAssignments := int64(0)
	db.Model(&assignmentsModel.Assignment{}).Count(&totalAssignments)
	return totalAssignments
}

// GetAssignmentByID returns the assignment with the specific ID
func GetAssignmentByID(id uint) assignmentsModel.Assignment {
	var assignment assignmentsModel.Assignment
	dbInstance.GetDBConnection().First(&assignment, id)
	return assignment
}

// UpdateAssignment updates the given assignment in the database
func UpdateAssignment(assignment *assignmentsModel.Assignment) error {
	err := dbInstance.GetDBConnection().Save(assignment).Error
	assignmentsDiagnostics.WriteAssignmentErr(err, "Update", assignment)
	return err
}

// DeleteAssignment deletes the given assignment from the database
func DeleteAssignment(assignment *assignmentsModel.Assignment) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(assignment).Error
	assignmentsDiagnostics.WriteAssignmentErr(err, "Delete", assignment)
	return err
}