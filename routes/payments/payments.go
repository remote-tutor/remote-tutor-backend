package payments

import (
	paymentsController "backend/controllers/payments"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes to deal with payments
func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	payments := e.Group("/payments", middleware.JWT([]byte("secret")))
	payments.GET("", paymentsController.GetPaymentsByUserAndMonth)

	adminPayments := adminRouter.Group("/payments")
	adminPayments.POST("", paymentsController.CreatePayment)
	adminPayments.DELETE("", paymentsController.DeletePayment)
}
