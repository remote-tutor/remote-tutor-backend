package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeGradeRoutes initializes all question routes
func InitializeGradeRoutes(quizzes *echo.Group, adminQuizzes *echo.Group) {
	quizzes.GET("/grades", quizzesController.GetGradesByQuiz)
	quizzes.GET("/grades/month", quizzesController.GetGradesByMonthAndUser)
	quizzes.GET("/grades/time", quizzesController.GetStudentRemainingTime)
	quizzes.POST("/grades", quizzesController.CreateQuizGrade)

	adminQuizzes.GET("/grades/month", quizzesController.GetGradesByMonthForAllUsers)
	adminQuizzes.GET("/grades/month/pdf", quizzesController.GenerateGradesPDF)
	// adminQuizzes.PUT("/grades", quizzesController.UpdateQuizTotalMark)
}
