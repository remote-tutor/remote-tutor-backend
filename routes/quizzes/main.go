package quizzes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for quizzes section
func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	quizzes := e.Group("/quizzes", middleware.JWT([]byte("secret")))
	adminQuizzes := adminRouter.Group("/quizzes")

	InitializeQuizRoutes(quizzes, adminQuizzes)
	InitializeQuestionRoutes(quizzes, adminQuizzes)
	InitializeChoiceRoutes(quizzes, adminQuizzes)
}
