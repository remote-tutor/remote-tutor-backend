package database

import (
	announcementsModel "backend/models/announcements"
	quizzesModel "backend/models/quizzes"
	usersModel "backend/models/users"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Millisecond, // Slow SQL threshold
			LogLevel:      logger.Info,      // Log level
			Colorful:      true,             // Disable color
		},
	)
	dsn                     = "root:password@tcp(127.0.0.1:3306)/tutoring?charset=utf8mb4&parseTime=True&loc=Local"
	databaseConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
)

// MigrateTables makes sure that the tables are migrated at the start of the application
func MigrateTables() {
	if err == nil {
		databaseConnection.AutoMigrate(&usersModel.User{})
		databaseConnection.AutoMigrate(&announcementsModel.Announcement{})

		databaseConnection.AutoMigrate(&quizzesModel.Quiz{})
		databaseConnection.AutoMigrate(&quizzesModel.MCQ{})
		databaseConnection.AutoMigrate(&quizzesModel.LongAnswer{})
		databaseConnection.AutoMigrate(&quizzesModel.Choice{})
		databaseConnection.AutoMigrate(&quizzesModel.MCQSubmission{})
		databaseConnection.AutoMigrate(&quizzesModel.LongAnswerSubmission{})
		databaseConnection.AutoMigrate(&quizzesModel.QuizGrade{})
	}
}

// GetDBConnection returns the DB connection
func GetDBConnection() *gorm.DB {
	return databaseConnection
}
