package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeQuestionRoutes initializes all question routes
func InitializeQuestionRoutes(quizzes *echo.Group, adminQuizzes *echo.Group) {
	quizzes.GET("/questions", quizzesController.GetQuestionsByQuiz)

	adminQuizzes.POST("/questions/mcq", quizzesController.CreateMCQ)
	adminQuizzes.POST("/questions/longanswer", quizzesController.CreateLongAnswer)

	adminQuizzes.PUT("/questions/mcq", quizzesController.UpdateMCQ)
	adminQuizzes.PUT("/questions/longanswer", quizzesController.UpdateLongAnswer)

	adminQuizzes.DELETE("/questions/mcq", quizzesController.DeleteMCQ)
	adminQuizzes.DELETE("/questions/longanswer", quizzesController.DeleteLongAnswer)
}
