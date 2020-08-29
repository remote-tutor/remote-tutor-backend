package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeQuizRoutes initializes all quiz routes
func InitializeQuizRoutes(quizzes *echo.Group) {
	quizzes.GET("/past", quizzesController.GetPastQuizzes)
	quizzes.GET("/future", quizzesController.GetFutureQuizzes)
	quizzes.GET("/current", quizzesController.GetCurrentQuizzes)
	quizzes.POST("", quizzesController.CreateQuiz)
}
