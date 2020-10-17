package payments

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	paymentsModel "backend/models/payments"
	"time"
)

// CreatePayment inserts a new payment to the database
func CreatePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Create(payment).Error
	diagnostics.WriteError(err, "CreatePayment")
	return err
}

// UpdatePayment updates the payment data in the database
func UpdatePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Save(payment).Error
	return err
}

// DeletePayment deletes the payment
func DeletePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(payment).Error
	diagnostics.WriteError(err, "DeletePayment")
	return err
}

// GetPaymentsByUserAndMonthAndClass gets the payment of specific user in a specific month
func GetPaymentsByUserAndMonthAndClass(userID uint, startDate, endDate time.Time, class string) []paymentsModel.Payment {
	payments := make([]paymentsModel.Payment, 0)
	dbInstance.GetDBConnection().Where("user_id = ? AND start_date >= ? AND end_date < ? AND class_hash = ?",
		userID, startDate, endDate, class).Find(&payments)
	return payments
}

// GetPaymentByUserAndWeekAndClass returns the payment of the user in a specific week (if found)
func GetPaymentByUserAndWeekAndClass(userID uint, eventDate time.Time, class string) paymentsModel.Payment {
	var payment paymentsModel.Payment
	dbInstance.GetDBConnection().Where("user_id = ? AND start_date <= ? AND end_date >= ? AND class_hash = ?",
		userID, eventDate, eventDate, class).First(&payment)
	return payment
}
