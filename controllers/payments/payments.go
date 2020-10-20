package payments

import (
	"net/http"

	"github.com/labstack/echo"

	authController "backend/controllers/auth"
	paymentsDBInteractions "backend/database/payments"
	paymentsModel "backend/models/payments"
	"backend/utils"
)

// GetPaymentsByUserAndMonthAndClass gets the payments of specific user in a specific month
func GetPaymentsByUserAndMonthAndClass(c echo.Context) error {
	admin := authController.FetchLoggedInUserAdminStatus(c)
	userID := uint(0)
	if admin {
		userID = utils.ConvertToUInt(c.QueryParam("userID"))
	} else {
		userID = authController.FetchLoggedInUserID(c)
	}
	startDate := utils.ConvertToTime(c.QueryParam("startDate"))
	endDate := utils.ConvertToTime(c.QueryParam("endDate"))
	endDate = endDate.AddDate(0, 0, 1)
	class := c.QueryParam("selectedClass")
	payments := paymentsDBInteractions.GetPaymentsByUserAndMonthAndClass(userID, startDate, endDate, class)
	return c.JSON(http.StatusOK, echo.Map{
		"payments": payments,
	})
}

// GetPaymentsByUserAndWeekAndClass verifies if the user has payment for this week or not
func GetPaymentsByUserAndWeekAndClass(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	eventDate := utils.ConvertToTime(c.QueryParam("eventDate"))
	class := c.QueryParam("selectedClass")
	payment := paymentsDBInteractions.GetPaymentByUserAndWeekAndClass(userID, eventDate, class)
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
	payment.ClassHash = c.FormValue("selectedClass")

	err := paymentsDBInteractions.CreatePayment(payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (payment not created), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{})
}

// DeletePayment deletes the payment from the database
func DeletePayment(c echo.Context) error {
	payment := new(paymentsModel.Payment)
	payment.ID = utils.ConvertToUInt(c.FormValue("id"))

	err := paymentsDBInteractions.DeletePayment(payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (payment not deleted), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{})
}

func GiveAccessToAllStudents(c echo.Context) error {
	startDate := utils.ConvertToTime(c.FormValue("startDate"))
	endDate := utils.ConvertToTime(c.FormValue("endDate"))
	class := c.FormValue("selectedClass")
	paymentsDBInteractions.GiveAccessToAllStudents(startDate, endDate, class)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Access has been given to all students",
	})
}