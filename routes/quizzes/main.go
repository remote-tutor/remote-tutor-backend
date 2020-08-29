package quizzes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for quizzes section
func InitializeRoutes(e *echo.Echo) {
	quizzes := e.Group("/quizzes")
	quizzes.Use(middleware.JWT([]byte("secret")))
	InitializeQuizRoutes(quizzes)
	InitializeQuestionRoutes(quizzes)
	InitializeChoiceRoutes(quizzes)
}
