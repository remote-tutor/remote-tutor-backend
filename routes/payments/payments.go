package payments

import (
	paymentsController "backend/controllers/payments"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes to deal with payments
func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	payments := e.Group("/payments", middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
	payments.GET("/month", paymentsController.GetPaymentsByUserAndMonthAndClass)
	payments.GET("/week", paymentsController.GetPaymentsByUserAndWeekAndClass)

	adminPayments := adminRouter.Group("/payments")
	adminPayments.GET("/week", paymentsController.GetPaymentsByWeekAndClass)
	adminPayments.POST("/week", paymentsController.UpdatePayments)
	adminPayments.POST("", paymentsController.CreatePayment)
	adminPayments.POST("/all", paymentsController.GiveAccessToAllStudents)
	adminPayments.DELETE("", paymentsController.DeletePayment)
}
