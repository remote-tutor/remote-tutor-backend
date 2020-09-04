package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeChoiceRoutes initializes all choice routes
func InitializeChoiceRoutes(quizzes *echo.Group, adminQuizzes *echo.Group) {
	adminQuizzes.POST("/choices", quizzesController.CreateChoice)
	adminQuizzes.PUT("/choices", quizzesController.UpdateChoice)
	adminQuizzes.DELETE("/choices", quizzesController.DeleteChoice)
}
