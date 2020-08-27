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
		databaseConnection.AutoMigrate(&quizzes.MCQ{}).AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.LongAnswer{}).AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.Choice{}).AddForeignKey("mcq_id", "mcqs(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.MCQSubmission{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("question_id", "mcqs(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.LongAnswerSubmission{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("question_id", "long_answers(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzes.QuizGrade{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")
	}
}

// GetDBConnection returns the DB connection
func GetDBConnection() *gorm.DB {
	return databaseConnection
}
