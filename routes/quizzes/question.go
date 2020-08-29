package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeQuestionRoutes initializes all question routes
func InitializeQuestionRoutes(quizzes *echo.Group) {
	quizzes.POST("questions/mcq", quizzesController.CreateMCQ)
	quizzes.POST("questions/longanswer", quizzesController.CreateLongAnswer)
	quizzes.PUT("questions/mcq", quizzesController.UpdateMCQ)
	quizzes.PUT("questions/longanswer", quizzesController.UpdateLongAnswer)
}
