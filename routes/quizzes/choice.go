package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeChoiceRoutes initializes all choice routes
func InitializeChoiceRoutes(quizzes *echo.Group) {
	quizzes.POST("/choices", quizzesController.CreateChoice)
}
