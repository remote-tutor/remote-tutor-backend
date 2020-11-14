package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeGradeRoutes initializes all question routes
func InitializeGradeRoutes(quizzes *echo.Group, adminQuizzes *echo.Group) {
	quizzes.GET("/grades", quizzesController.GetGradesByQuiz)
	quizzes.GET("/grades/time", quizzesController.GetStudentRemainingTime)
	quizzes.POST("/grades", quizzesController.CreateQuizGrade)

	adminQuizzes.GET("/grades", quizzesController.GetGradesByQuizForAllUsers)
	// adminQuizzes.PUT("/grades", quizzesController.UpdateQuizTotalMark)
}
