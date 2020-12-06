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
		databaseConnection.AutoMigrate(&usersModel.User{},
			&announcementsModel.Announcement{},
			&quizzesModel.Quiz{},
			&quizzesModel.MCQ{},
			&quizzesModel.LongAnswer{},
			&quizzesModel.Choice{},
			&quizzesModel.MCQSubmission{},
			&quizzesModel.LongAnswerSubmission{},
			&quizzesModel.QuizGrade{},
			&paymentsModel.Payment{},
			&assignmentsModel.Assignment{},
			&assignmentsModel.AssignmentSubmission{},
			&videosModel.Video{},
			&videosModel.VideoPart{},
			&videosModel.UserWatch{},
			&videosModel.Code{},
			&organizationsModel.Organization{},
			&organizationsModel.Class{},
			&organizationsModel.ClassUser{},
			)
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
