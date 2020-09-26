package payments

import (
	"net/http"

	"github.com/labstack/echo"

	authController "backend/controllers/auth"
	paymentsDBInteractions "backend/database/payments"
	paymentsModel "backend/models/payments"
	"backend/utils"
)

// GetPaymentsByUserAndMonth gets the payments of specific user in a specific month
func GetPaymentsByUserAndMonth(c echo.Context) error {
	userID := utils.ConvertToUInt(c.QueryParam("userID"))
	startDate := utils.ConvertToTime(c.QueryParam("startDate"))
	endDate := utils.ConvertToTime(c.QueryParam("endDate"))
	endDate = endDate.AddDate(0, 0, 1)

	payments := paymentsDBInteractions.GetPaymentsByUserAndMonth(userID, startDate, endDate)
	return c.JSON(http.StatusOK, echo.Map{
		"payments": payments,
	})
}

// GetPaymentsByUserAndWeek verifies if the user has payment for this week or not
func GetPaymentsByUserAndWeek(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	eventDate := utils.ConvertToTime(c.QueryParam("eventDate"))

	payment := paymentsDBInteractions.GetPaymentByUserAndWeek(userID, eventDate)
	if payment.ID == 0 {
		return c.JSON(http.StatusOK, echo.Map{
			"status": false,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
	})
}

// CreatePayment creates a new payment to the user
func CreatePayment(c echo.Context) error {
	payment := new(paymentsModel.Payment)
	payment.UserID = utils.ConvertToUInt(c.FormValue("userID"))
	payment.StartDate = utils.ConvertToTime(c.FormValue("startDate"))
	payment.EndDate = utils.ConvertToTime(c.FormValue("endDate"))

	paymentsDBInteractions.CreatePayment(payment)
	return c.JSON(http.StatusOK, echo.Map{})
}

// DeletePayment deletes the payment from the database
func DeletePayment(c echo.Context) error {
	payment := new(paymentsModel.Payment)
	payment.ID = utils.ConvertToUInt(c.FormValue("id"))

	paymentsDBInteractions.DeletePayment(payment)
	return c.JSON(http.StatusOK, echo.Map{})
}
