package routes

import (
	authControllers "backend/controllers/auth"
	announcementRouter "backend/routes/announcements"
	assignmentsRouter "backend/routes/assignments"
	organizationsRouter "backend/routes/organizations"
	paymentRouter "backend/routes/payments"
	quizRouter "backend/routes/quizzes"
	userRouter "backend/routes/users"
	videosRouter "backend/routes/videos"
	"os"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for the application
func InitializeRoutes(e *echo.Echo) {
	// to enable sending requests from the frontend application
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080",
			"http://192.168.1.100:8080",
			"https://remote-tutor.github.io",
			"https://thematrixeg.com",
			"https://remotetutoreg.com",
		},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Logger())
	middleware.ErrJWTMissing.Message = "Please login"
	adminRouter := e.Group("/admin",
		middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))),
		authControllers.CheckAdmin)

	userRouter.InitializeRoutes(e, adminRouter)
	announcementRouter.InitializeRoutes(e, adminRouter)
	quizRouter.InitializeRoutes(e, adminRouter)
	paymentRouter.InitializeRoutes(e, adminRouter)
	assignmentsRouter.InitializeRoutes(e, adminRouter)
	videosRouter.InitializeRoutes(e, adminRouter)
	organizationsRouter.InitializeRoutes(e, adminRouter)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "From APache")
	})

}
