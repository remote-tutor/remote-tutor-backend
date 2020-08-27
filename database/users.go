package database

import (
	md "backend/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

// GetUserByUsername searches the database for the user with the given username to be used in the login action
func GetUserByUsername(username string) md.User {
	db := GetDBConnection()
	var user md.User
	db.Where("username = ?", username).First(&user)
	return user
}

// GetUserByUserID searches the database for the user with the given userid
func GetUserByUserID(userid uint) md.User {
	db := GetDBConnection()
	var user md.User
	db.Where("userid = ?", userid).First(&user)
	return user
}

// CreateUser inserts a new user to the database
func CreateUser(user *md.User) {
	db := GetDBConnection()
	db.Create(user)
}

// GetStudents retrieves the students
func GetStudents(fullName, sortBy, searchByValue, searchByField string,
	page, itemsPerPage int, pending, sortDesc bool) ([]md.User, int) {

	var students []md.User
	query := GetDBConnection().Where(fmt.Sprintf("%s LIKE '%%%s%%'", searchByField, searchByValue))
	numberOfRecords := countStudents(query)
	sortByOrder := sortBy
	if sortDesc {
		sortByOrder += " DESC"
	}
	query.Offset(itemsPerPage * (page - 1)).Limit(itemsPerPage).
		Order(sortByOrder).Find(&students)
	return students, numberOfRecords
}

func countStudents(db *gorm.DB) int {
	totalStudents := 0
	db.Model(&md.User{}).Count(&totalStudents)
	return totalStudents
}
