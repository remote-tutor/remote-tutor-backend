package organizations

import (
	userClassesController "backend/controllers/organizations"
	"github.com/labstack/echo"
)

func InitializeUerClassesRoutes(classes *echo.Group, adminClasses *echo.Group) {
	classes.GET("", userClassesController.GetClassesByUser)
	classes.POST("/enroll", userClassesController.Enroll)
}
