package routes

import (
	"backend/controllers"
	quizzesController "backend/controllers/quizzes"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for the application
func InitializeRoutes(e *echo.Echo) {
	// to enable sending requests from the frontend application
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)

	e.GET("/get-pending-students", controllers.GetPendingUsers)

	announcements := e.Group("/announcements")
	announcements.Use(middleware.JWT([]byte("secret")))
	announcements.GET("", controllers.GetAnnouncements)
	announcements.POST("", controllers.CreateAnnouncement)
	announcements.PUT("", controllers.UpdateAnnouncement)
	announcements.DELETE("", controllers.DeleteAnnouncement)

	quizzes := e.Group("/quizzes")
	quizzes.Use(middleware.JWT([]byte("secret")))
	quizzes.GET("/past", quizzesController.GetPastQuizzes)
	quizzes.GET("/future", quizzesController.GetFutureQuizzes)
	quizzes.GET("/current", quizzesController.GetCurrentQuizzes)
	quizzes.POST("", quizzesController.CreateQuiz)

	e.GET("/isAdmin", controllers.GetAnnouncements)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success")
	}, middleware.JWT([]byte("secret")))

}
