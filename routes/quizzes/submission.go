package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeSubmissionRoutes initializes all question routes
func InitializeSubmissionRoutes(quizzes *echo.Group, adminQuizzes *echo.Group) {
	quizzes.GET("/submissions", quizzesController.GetSubmissionsByQuizAndUser)
	quizzes.POST("/submissions/mcq", quizzesController.CreateMCQSubmission)
	quizzes.PUT("/submissions/mcq", quizzesController.UpdateMCQSubmission)
}
