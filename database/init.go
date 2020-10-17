package database

import (
	announcementsModel "backend/models/announcements"
	assignmentsModel "backend/models/assignments"
	organizationsModel "backend/models/organizations"
	paymentsModel "backend/models/payments"
	quizzesModel "backend/models/quizzes"
	usersModel "backend/models/users"
	videosModel "backend/models/videos"
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
	databaseConnection *gorm.DB = nil
	err error
)

// MigrateTables makes sure that the tables are migrated at the start of the application
func MigrateTables() {
	initializeDBConnection()
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

		databaseConnection.AutoMigrate(&paymentsModel.Payment{})

		databaseConnection.AutoMigrate(&assignmentsModel.Assignment{})
		databaseConnection.AutoMigrate(&assignmentsModel.AssignmentSubmission{})

		databaseConnection.AutoMigrate(&videosModel.Video{})
		databaseConnection.AutoMigrate(&videosModel.VideoPart{})
		databaseConnection.AutoMigrate(&videosModel.UserWatch{})

		databaseConnection.AutoMigrate(&organizationsModel.Organization{})
		databaseConnection.AutoMigrate(&organizationsModel.Class{})
		databaseConnection.AutoMigrate(&organizationsModel.ClassUser{})
	}
}

func initializeDBConnection() {
	dsn := os.ExpandEnv("${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Africa%2FCairo")
	connection, connectionErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if connectionErr == nil {
		databaseConnection = connection
	}
	err = connectionErr
}

// GetDBConnection returns the DB connection
func GetDBConnection() *gorm.DB {
	if databaseConnection == nil {
		initializeDBConnection()
	}
	return databaseConnection
}
