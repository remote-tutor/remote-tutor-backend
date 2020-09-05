package users

import (
	usersController "backend/controllers/users"

	"github.com/labstack/echo"
)

// InitializeRoutes initializes all the required routes for users
func InitializeRoutes(e *echo.Echo, adminRoute *echo.Group) {
	e.POST("/login", usersController.Login)
	e.POST("/register", usersController.Register)
	e.GET("/get-pending-students", usersController.GetPendingUsers)
}
