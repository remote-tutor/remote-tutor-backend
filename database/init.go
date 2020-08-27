package database

import (
	md "backend/models"
	"backend/models/quizzes"

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
		databaseConnection.AutoMigrate(&md.Announcement{})
		databaseConnection.AutoMigrate(&quizzes.Quiz{})
		databaseConnection.AutoMigrate(&quizzes.MCQ{}).AddForeignKey("quiz_id", "quiz(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.LongAnswer{}).AddForeignKey("quiz_id", "quiz(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.Choice{}).AddForeignKey("mcq_id", "mcq(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.MCQSubmission{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE").AddForeignKey("mcq_id", "mcq(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.LongAnswerSubmission{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE").AddForeignKey("long_answer_id", "long_answer(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.QuizGrade{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE").AddForeignKey("quiz_id", "quiz(id)", "CASCADE", "CASCADE")
	}
}

// GetDBConnection returns the DB connection
func GetDBConnection() *gorm.DB {
	return databaseConnection
}
