package assignments

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	dbPagination "backend/database/scopes"
	assignmentsModel "backend/models/assignments"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// CreateAssignment inserts a new assignment to the database
func CreateAssignment(assignment *assignmentsModel.Assignment) error {
	err := dbInstance.GetDBConnection().Create(assignment).Error
	diagnostics.WriteError(err, "CreateAssignment")
	return err
}

// GetAssignments gets an array of assignments to display to the user
func GetAssignments(c echo.Context, year int) ([]assignmentsModel.Assignment, int64) {
	assignments := make([]assignmentsModel.Assignment, 0)
	db := dbInstance.GetDBConnection().Where("year = ?", year)
	totalAssignments := countAssignments(db)
	db.Scopes(dbPagination.Paginate(c)).Find(&assignments)
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
	diagnostics.WriteError(err, "UpdateAssignment")
	return err
}

// DeleteAssignment deletes the given assignment from the database
func DeleteAssignment(assignment *assignmentsModel.Assignment) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(assignment).Error
	diagnostics.WriteError(err, "DeleteAssignment")
	return err
}