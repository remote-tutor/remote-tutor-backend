package organizations

import (
	classUsersController "backend/controllers/organizations"
	"github.com/labstack/echo"
)

func InitializeUerClassesRoutes(classes *echo.Group, adminClasses *echo.Group) {
	classes.GET("", classUsersController.GetClassesByUser)
	classes.POST("/enroll", classUsersController.Enroll)

	adminClasses.GET("/students", classUsersController.GetStudentsByClass)
	adminClasses.PUT("/students", classUsersController.AcceptStudents)
}
