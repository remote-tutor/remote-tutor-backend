package database

import (
	md "backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // to import the gorm database wrapper
)

var (
	databaseConnection, err = gorm.Open("mysql", "root:password@/tutoring?charset=utf8&parseTime=True&loc=Local")
)

// MigrateTables makes sure that the tables are migrated at the start of the application
func MigrateTables() {
	if err == nil {
		databaseConnection.LogMode(true)
		databaseConnection.AutoMigrate(&md.User{})
	}
}

// GetDBConnection returns the DB connection
func GetDBConnection() *gorm.DB {
	return databaseConnection
}
