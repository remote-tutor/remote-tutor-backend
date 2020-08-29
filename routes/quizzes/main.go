package quizzes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for the application
func InitializeRoutes(e *echo.Echo) {
	quizzes := e.Group("/announcements")
	quizzes.Use(middleware.JWT([]byte("secret")))
	InitializeQuizRoutes(quizzes)
	InitializeQuestionRoutes(quizzes)
	InitializeChoiceRoutes(quizzes)
}
