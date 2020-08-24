package routes

import (
	"backend/controllers"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for the application
func InitializeRoutes(e *echo.Echo) {
	// to enable sending requests from the frontend application
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
	e.POST("/announcement", controllers.CreateAnnouncement)

	e.PUT("/announcement", controllers.UpdateAnnouncement)

	e.GET("/get-pending-students", controllers.GetPendingUsers)
	e.GET("/announcements", controllers.GetAnnouncements)
	e.GET("/isAdmin", controllers.GetAnnouncements)

	e.DELETE("./announcement", controllers.DeleteAnnouncement)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success")
	})
}
