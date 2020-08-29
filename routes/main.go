package routes

import (
	announcementRouter "backend/routes/announcements"
	quizRouter "backend/routes/quizzes"
	userRouter "backend/routes/users"

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

	userRouter.InitializeRoutes(e)
	announcementRouter.InitializeRoutes(e)
	quizRouter.InitializeRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success")
	}, middleware.JWT([]byte("secret")))

}
