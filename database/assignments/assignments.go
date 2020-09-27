package assignments

import (
	dbInstance "backend/database"
	assignmentsModel "backend/models/assignments"
)

// CreateAssignment inserts a new assignment to the database
func CreateAssignment(assignment *assignmentsModel.Assignment) {
	dbInstance.GetDBConnection().Create(assignment)
}

// GetAssignmentByID returns the assignment with the specific ID
func GetAssignmentByID(id uint) assignmentsModel.Assignment {
	var assignment assignmentsModel.Assignment
	dbInstance.GetDBConnection().First(&assignment, id)
	return assignment
}

// UpdateAssignment updates the given assignment in the database
func UpdateAssignment(assignment *assignmentsModel.Assignment) {
	dbInstance.GetDBConnection().Save(assignment)
}

// DeleteAssignment deletes the given assignment from the database
func DeleteAssignment(assignment *assignmentsModel.Assignment) {
	dbInstance.GetDBConnection().Unscoped().Delete(assignment)
}