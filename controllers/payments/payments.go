package payments

import (
	"net/http"
	"time"

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

// GetPaymentsByWeekAndClass gets the payments for specific collection of user by a specific week
func GetPaymentsByWeekAndClass(c echo.Context) error {
	date := utils.ConvertToStartOfDay(c.QueryParam("date"))
	if date.Weekday() != time.Friday {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Please select start of the week (Friday) to proceed",
		})
	}
	endDate := date.AddDate(0, 0, 7)
	class := c.QueryParam("selectedClass")
	queryParams := c.Request().URL.Query()
	usersIDs := utils.ConvertToUIntArray(queryParams["usersIDs[]"])
	payments := paymentsDBInteractions.GetPaymentsByWeekAndClass(usersIDs, date, endDate, class)
	return c.JSON(http.StatusOK, echo.Map{
		"payments": payments,
	})
}

// CreatePayment creates a new payment to the user
func CreatePayment(c echo.Context) error {
	payment := new(paymentsModel.Payment)
	payment.UserID = utils.ConvertToUInt(c.FormValue("userID"))
	payment.StartDate = utils.ConvertToStartOfDay(c.FormValue("startDate"))
	payment.EndDate = utils.ConvertToStartOfDay(c.FormValue("endDate"))
	payment.ClassHash = c.FormValue("selectedClass")

	err := paymentsDBInteractions.CreatePayment(payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (payment not created), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{})
}

// UpdatePayments creates/removes payments for a specific group of students
func UpdatePayments(c echo.Context) error {
	startDate := utils.ConvertToStartOfDay(c.FormValue("startDate"))
	endDate := startDate.AddDate(0, 0, 7)
	classHash := c.FormValue("selectedClass")
	addedTo := utils.ConvertToUIntArray(
		utils.ConvertToFormArray(c.FormValue("addedTo[]")))
	removedFrom := utils.ConvertToUIntArray(
		utils.ConvertToFormArray(c.FormValue("removedFrom[]")))
	paymentsToAdd := make([]paymentsModel.Payment, len(addedTo))
	for index, idToAdd := range addedTo {
		paymentsToAdd[index] = paymentsModel.Payment{
			UserID:    idToAdd,
			StartDate: startDate,
			EndDate:   endDate,
			ClassHash: classHash,
		}
	}
	paymentsToRemove := paymentsDBInteractions.GetPaymentsByWeekAndClass(removedFrom, startDate, endDate, classHash)
	err := paymentsDBInteractions.UpdatePayments(paymentsToAdd, paymentsToRemove)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (payments not fully updated), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Payments updated successfully",
	})
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
