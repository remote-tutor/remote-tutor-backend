package users

import (
	usersController "backend/controllers/users"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for users
func InitializeRoutes(e *echo.Echo, adminRoute *echo.Group) {
	e.POST("/login", usersController.Login)
	e.POST("/register", usersController.Register)

	adminRoute.GET("/students", usersController.GetUsers)
	e.PUT("/students", usersController.UpdateUser)

	e.PUT("/change-password", usersController.ChangePassword, middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
}
