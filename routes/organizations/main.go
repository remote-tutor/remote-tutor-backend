package organizations

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	classes := e.Group("/classes", middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
	adminClasses := adminRouter.Group("/classes")

	InitializeUerClassesRoutes(classes, adminClasses)
}

